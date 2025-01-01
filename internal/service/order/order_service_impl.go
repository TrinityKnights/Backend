package order

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/TrinityKnights/Backend/internal/domain/entity"
	"github.com/TrinityKnights/Backend/internal/domain/model"
	"github.com/TrinityKnights/Backend/internal/domain/model/converter"
	"github.com/TrinityKnights/Backend/internal/repository/order"
	"github.com/TrinityKnights/Backend/internal/repository/ticket"
	"github.com/TrinityKnights/Backend/internal/service/payment"
	"github.com/TrinityKnights/Backend/pkg/cache"
	domainErrors "github.com/TrinityKnights/Backend/pkg/errors"
	"github.com/TrinityKnights/Backend/pkg/helper"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderServiceImpl struct {
	DB               *gorm.DB
	Cache            *cache.ImplCache
	Log              *logrus.Logger
	Validate         *validator.Validate
	OrderRepository  order.OrderRepository
	TicketRepository ticket.TicketRepository
	PaymentService   payment.PaymentService
	helper           *helper.ContextHelper
}

func NewOrderServiceImpl(db *gorm.DB, cacheImpl *cache.ImplCache, log *logrus.Logger, validate *validator.Validate, orderRepository order.OrderRepository, ticketRepository ticket.TicketRepository, paymentService payment.PaymentService) *OrderServiceImpl {
	return &OrderServiceImpl{
		DB:               db,
		Cache:            cacheImpl,
		Log:              log,
		Validate:         validate,
		OrderRepository:  orderRepository,
		TicketRepository: ticketRepository,
		PaymentService:   paymentService,
		helper:           helper.NewContextHelper(),
	}
}

func (s *OrderServiceImpl) CreateOrder(ctx context.Context, request *model.OrderTicketRequest) (*model.OrderResponse, error) {
	if err := s.Validate.Struct(request); err != nil {
		return nil, domainErrors.ErrValidation
	}

	// Get user ID from JWT claims
	claims, err := s.helper.GetJWTClaims(ctx)
	if err != nil {
		return nil, domainErrors.ErrUnauthorized
	}

	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// Check if event exists
	var event entity.Event
	if err := tx.First(&event, request.EventID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domainErrors.ErrNotFound
		}
		s.Log.Errorf("failed to get event: %v", err)
		return nil, domainErrors.ErrInternalServer
	}

	// First check if tickets exist and are available (without locking)
	tickets, err := s.TicketRepository.Find(tx, &model.TicketQueryOptions{
		EventID:     &event.ID,
		SeatNumbers: &request.SeatNumbers,
	})
	if err != nil {
		s.Log.Errorf("failed to get tickets: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrNotFound
		}
		return nil, domainErrors.ErrInternalServer
	}

	// Check if any seats are already taken
	for _, ticket := range tickets {
		if ticket.OrderID != nil {
			return nil, domainErrors.ErrSeatAlreadyTaken
		}
	}

	// Now lock and update the specific tickets one by one
	targetTickets := make([]*entity.Ticket, 0, len(request.TicketIDs))
	totalPrice := 0.0

	s.Log.Infof("Verifying tickets - IDs: %v, Seats: %v", request.TicketIDs, request.SeatNumbers)

	for _, ticketID := range request.TicketIDs {
		var t entity.Ticket
		// Lock individual ticket for update
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ? AND order_id IS NULL", ticketID).
			First(&t).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, domainErrors.ErrSeatAlreadyTaken
			}
			s.Log.Errorf("failed to lock ticket: %v", err)
			return nil, domainErrors.ErrInternalServer
		}

		// Add ticket to our target tickets slice
		targetTickets = append(targetTickets, &t)
		totalPrice += t.Price
	}

	// Convert pointer slice to value slice
	orderTickets := make([]entity.Ticket, len(targetTickets))
	for i, t := range targetTickets {
		orderTickets[i] = *t
	}

	dataOrder := entity.Order{
		UserID:     claims.UserID,
		Date:       time.Now(),
		TotalPrice: totalPrice,
		Tickets:    orderTickets,
	}

	if err := s.OrderRepository.Create(tx, &dataOrder); err != nil {
		s.Log.Errorf("failed to create order: %v", err)
		return nil, domainErrors.ErrInternalServer
	}

	// Update tickets one by one
	for _, ticket := range targetTickets {
		ticketType := helper.TicketUpper(ticket.Type)
		updateTicket := entity.Ticket{
			ID:         ticket.ID,
			EventID:    ticket.EventID,
			OrderID:    &dataOrder.ID,
			Price:      ticket.Price,
			Type:       ticketType.Long,
			SeatNumber: ticket.SeatNumber,
		}

		if err := s.TicketRepository.Update(tx, &updateTicket); err != nil {
			s.Log.Errorf("failed to update ticket: %v", err)
			return nil, domainErrors.ErrInternalServer
		}
	}

	// Reload order with tickets
	if err := tx.Preload("Tickets").First(&dataOrder, dataOrder.ID).Error; err != nil {
		s.Log.Errorf("failed to reload order: %v", err)
		return nil, domainErrors.ErrInternalServer
	}

	// After creating the order and updating the tickets, create payment
	paymentRequest := &model.CreatePaymentRequest{
		OrderID: dataOrder.ID,
		Amount:  dataOrder.TotalPrice,
	}

	p, err := s.PaymentService.CreateInvoice(ctx, tx, paymentRequest)
	if err != nil {
		s.Log.Errorf("failed to create payment: %v", err)
		return nil, domainErrors.ErrInternalServer
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.Errorf("failed to commit transaction: %v", err)
		return nil, domainErrors.ErrInternalServer
	}

	response := converter.OrderEntityToResponse(&dataOrder)
	response.Payment = p

	return response, nil
}

func (s *OrderServiceImpl) GetOrderByID(ctx context.Context, request *model.GetOrderRequest) (*model.OrderResponse, error) {
	if err := s.Validate.Struct(request); err != nil {
		return nil, domainErrors.ErrValidation
	}

	// Try to get from cache first
	key := fmt.Sprintf("order:get:id:%d", request.ID)
	var data *model.OrderResponse
	err := s.Cache.Get(key, &data)
	if err != nil && !errors.Is(err, cache.ErrCacheMiss) {
		s.Log.Errorf("failed to get cache: %v", err)
	}

	if data == nil {
		tx := s.DB.WithContext(ctx)

		var dataOrder entity.Order
		if err := s.OrderRepository.GetByIDWithDetails(tx, &dataOrder, request.ID); err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, domainErrors.ErrNotFound
			}
			s.Log.Errorf("failed to get order: %v", err)
			return nil, domainErrors.ErrInternalServer
		}

		response := converter.OrderEntityToResponse(&dataOrder)

		// Cache the response
		if err := s.Cache.Set(key, response, 5*time.Minute); err != nil {
			s.Log.Errorf("failed to set cache: %v", err)
		}

		return response, nil
	}

	return data, nil
}

func (s *OrderServiceImpl) GetOrders(ctx context.Context, request *model.OrdersRequest) (*model.Response[[]*model.OrderResponse], error) {
	if err := s.Validate.Struct(request); err != nil {
		return nil, domainErrors.ErrValidation
	}

	if request.Size <= 0 {
		request.Size = 10
	}
	if request.Page <= 0 {
		request.Page = 1
	}

	// Try to get from cache first
	cacheKey := fmt.Sprintf("order:get:page:%d:size:%d:sort:%s:order:%s",
		request.Page, request.Size, request.Sort, request.Order)
	var cacheResponse model.Response[[]*model.OrderResponse]
	if err := s.Cache.Get(cacheKey, &cacheResponse); err == nil {
		return &cacheResponse, nil
	}

	var orders []entity.Order
	totalItems, err := s.OrderRepository.GetPaginatedOrders(
		s.DB.WithContext(ctx),
		&orders,
		request.Page,
		request.Size,
		request.Sort,
		request.Order,
	)
	if err != nil {
		s.Log.Errorf("failed to get orders: %v", err)
		return nil, domainErrors.ErrInternalServer
	}

	if len(orders) == 0 {
		return nil, domainErrors.ErrNotFound
	}

	response := converter.OrdersToPaginatedResponse(orders, totalItems, request.Page, request.Size)

	// Cache the response
	if err := s.Cache.Set(cacheKey, response, 5*time.Minute); err != nil {
		s.Log.Errorf("failed to cache response: %v", err)
	}

	return response, nil
}
