package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.57

import (
	"context"
	"time"

	"github.com/TrinityKnights/Backend/internal/delivery/graph"
	graphmodel "github.com/TrinityKnights/Backend/internal/delivery/graph/model"
	"github.com/TrinityKnights/Backend/internal/domain/model"
)

// ID is the resolver for the id field.
func (r *eventResponseResolver) ID(ctx context.Context, obj *model.EventResponse) (int, error) {
	return int(obj.ID), nil
}

// Time is the resolver for the time field.
func (r *eventResponseResolver) Time(ctx context.Context, obj *model.EventResponse) (*time.Time, error) {
	t := time.Time(obj.Time)
	return &t, nil
}

// VenueID is the resolver for the venueId field.
func (r *eventResponseResolver) VenueID(ctx context.Context, obj *model.EventResponse) (int, error) {
	return int(obj.VenueID), nil
}

// Venue is the resolver for the venue field.
func (r *eventResponseResolver) Venue(ctx context.Context, obj *model.EventResponse) (*model.VenueResponse, error) {
	venue, err := r.VenueService.GetVenueByID(ctx, &model.GetVenueRequest{
		ID: uint(obj.VenueID),
	})
	if err != nil {
		return nil, err
	}
	return venue, nil
}

// CreateEvent is the resolver for the createEvent field.
func (r *mutationResolver) CreateEvent(ctx context.Context, name string, description string, date string, time string, venueID int) (*model.EventResponse, error) {
	event, err := r.EventService.CreateEvent(ctx, &model.CreateEventRequest{
		Name:        name,
		Description: description,
		Date:        date,
		Time:        time,
		VenueID:     uint(venueID),
	})
	if err != nil {
		return nil, err
	}
	return event, nil
}

// UpdateEvent is the resolver for the updateEvent field.
func (r *mutationResolver) UpdateEvent(ctx context.Context, id int, input graphmodel.UpdateEventInput) (*model.EventResponse, error) {
	event, err := r.EventService.UpdateEvent(ctx, &model.UpdateEventRequest{
		ID:          uint(id),
		Name:        *input.Name,
		Description: *input.Description,
		Date:        *input.Date,
		Time:        *input.Time,
		VenueID:     uint(*input.VenueID),
	})
	if err != nil {
		return nil, err
	}
	return event, nil
}

// CreateVenue is the resolver for the createVenue field.
func (r *mutationResolver) CreateVenue(ctx context.Context, name string, address string, capacity int, city string, state string, zip string) (*model.VenueResponse, error) {
	venue, err := r.VenueService.CreateVenue(ctx, &model.CreateVenueRequest{
		Name:     name,
		Address:  address,
		Capacity: capacity,
		City:     city,
		State:    state,
		Zip:      zip,
	})
	if err != nil {
		return nil, err
	}
	return venue, nil
}

// UpdateVenue is the resolver for the updateVenue field.
func (r *mutationResolver) UpdateVenue(ctx context.Context, id int, input graphmodel.UpdateVenueInput) (*model.VenueResponse, error) {
	venue, err := r.VenueService.UpdateVenue(ctx, &model.UpdateVenueRequest{
		ID:       uint(id),
		Name:     *input.Name,
		Address:  *input.Address,
		Capacity: *input.Capacity,
		City:     *input.City,
		State:    *input.State,
		Zip:      *input.Zip,
	})
	if err != nil {
		return nil, err
	}
	return venue, nil
}

// CreateTicket is the resolver for the createTicket field.
func (r *mutationResolver) CreateTicket(ctx context.Context, input graphmodel.CreateTicketInput) ([]*graphmodel.TicketResponse, error) {
	tickets, err := r.TicketService.CreateTicket(ctx, &model.CreateTicketRequest{
		EventID: uint(input.EventID),
		Price:   float64(input.Price),
		Type:    input.Type,
		Count:   input.Count,
	})
	if err != nil {
		return nil, err
	}

	graphTickets := make([]*graphmodel.TicketResponse, len(tickets))
	for i, ticket := range tickets {
		graphTickets[i] = &graphmodel.TicketResponse{
			ID:         ticket.ID,
			EventID:    int(ticket.EventID),
			Price:      ticket.Price,
			Type:       ticket.Type,
			SeatNumber: ticket.SeatNumber,
		}
	}
	return graphTickets, nil
}

// UpdateTicket is the resolver for the updateTicket field.
func (r *mutationResolver) UpdateTicket(ctx context.Context, id string, input graphmodel.UpdateTicketInput) (*graphmodel.TicketResponse, error) {
	var orderID uint
	if input.OrderID != nil {
		orderID = uint(*input.OrderID)
	}

	ticket, err := r.TicketService.UpdateTicket(ctx, &model.UpdateTicketRequest{
		ID:         id,
		Price:      *input.Price,
		OrderID:    &orderID,
		Type:       *input.Type,
		SeatNumber: *input.SeatNumber,
	})
	if err != nil {
		return nil, err
	}
	orderIDInt := int(ticket.OrderID)
	return &graphmodel.TicketResponse{
		ID:         ticket.ID,
		OrderID:    &orderIDInt,
		EventID:    int(ticket.EventID),
		Price:      ticket.Price,
		Type:       ticket.Type,
		SeatNumber: ticket.SeatNumber,
	}, nil
}

// Event is the resolver for the event field.
func (r *queryResolver) Event(ctx context.Context, id int) (*model.EventResponse, error) {
	event, err := r.EventService.GetEventByID(ctx, &model.GetEventRequest{
		ID: uint(id),
	})
	if err != nil {
		return nil, err
	}
	return event, nil
}

// Events is the resolver for the events field.
func (r *queryResolver) Events(ctx context.Context, page *int, size *int, sort *string, order *string) (*graphmodel.EventsResponse, error) {
	// Set default values
	defaultPage := 1
	defaultSize := 10
	defaultSort := "created_at"
	defaultOrder := "desc"

	// Use provided values or defaults
	requestPage := defaultPage
	if page != nil {
		requestPage = *page
	}

	requestSize := defaultSize
	if size != nil {
		requestSize = *size
	}

	requestSort := defaultSort
	if sort != nil {
		requestSort = *sort
	}

	requestOrder := defaultOrder
	if order != nil {
		requestOrder = *order
	}

	paginated, err := r.EventService.GetEvents(ctx, &model.EventsRequest{
		Page:  requestPage,
		Size:  requestSize,
		Sort:  requestSort,
		Order: requestOrder,
	})
	if err != nil {
		return nil, err
	}

	return &graphmodel.EventsResponse{
		Data:   *paginated.Data,
		Paging: (*graphmodel.PageMetadata)(paginated.Paging),
	}, nil
}

// SearchEvents is the resolver for the searchEvents field.
func (r *queryResolver) SearchEvents(ctx context.Context, name *string, description *string, date *string, time *string, venueID *int, page *int, size *int, sort *string, order *string) (*graphmodel.EventsResponse, error) {
	defaultPage := 1
	defaultSize := 10
	defaultSort := "created_at"
	defaultOrder := "desc"

	var nameStr string
	if name != nil {
		nameStr = *name
	}

	var descriptionStr string
	if description != nil {
		descriptionStr = *description
	}

	var dateStr string
	if date != nil {
		dateStr = *date
	}

	var timeStr string
	if time != nil {
		timeStr = *time
	}

	var venueIDUint uint
	if venueID != nil {
		venueIDUint = uint(*venueID)
	}

	requestPage := defaultPage
	if page != nil {
		requestPage = *page
	}

	requestSize := defaultSize
	if size != nil {
		requestSize = *size
	}

	requestSort := defaultSort
	if sort != nil {
		requestSort = *sort
	}

	requestOrder := defaultOrder
	if order != nil {
		requestOrder = *order
	}

	request := &model.EventSearchRequest{
		Name:        nameStr,
		Description: descriptionStr,
		Date:        dateStr,
		Time:        timeStr,
		VenueID:     venueIDUint,
		Page:        requestPage,
		Size:        requestSize,
		Sort:        requestSort,
		Order:       requestOrder,
	}

	events, err := r.EventService.SearchEvents(ctx, request)
	if err != nil {
		return nil, err
	}

	return &graphmodel.EventsResponse{
		Data:   *events.Data,
		Paging: (*graphmodel.PageMetadata)(events.Paging),
	}, nil
}

// Ticket is the resolver for the ticket field.
func (r *queryResolver) Ticket(ctx context.Context, id string) (*graphmodel.TicketResponse, error) {
	ticket, err := r.TicketService.GetTicketByID(ctx, &model.GetTicketRequest{
		ID: id,
	})
	if err != nil {
		return nil, err
	}

	orderIDInt := int(ticket.OrderID)
	return &graphmodel.TicketResponse{
		ID:         ticket.ID,
		EventID:    int(ticket.EventID),
		OrderID:    &orderIDInt,
		Price:      ticket.Price,
		Type:       ticket.Type,
		SeatNumber: ticket.SeatNumber,
	}, nil
}

// Tickets is the resolver for the tickets field.
func (r *queryResolver) Tickets(ctx context.Context, page *int, size *int, sort *string, order *string) (*graphmodel.TicketsResponse, error) {
	defaultPage := 1
	defaultSize := 10
	defaultSort := "created_at"
	defaultOrder := "desc"

	requestPage := defaultPage
	if page != nil {
		requestPage = *page
	}

	requestSize := defaultSize
	if size != nil {
		requestSize = *size
	}

	requestSort := defaultSort
	if sort != nil {
		requestSort = *sort
	}

	requestOrder := defaultOrder
	if order != nil {
		requestOrder = *order
	}

	paginated, err := r.TicketService.GetTickets(ctx, &model.TicketsRequest{
		Page:  requestPage,
		Size:  requestSize,
		Sort:  requestSort,
		Order: requestOrder,
	})
	if err != nil {
		return nil, err
	}

	graphTickets := make([]*graphmodel.TicketResponse, len(*paginated.Data))
	if paginated.Data != nil {
		for i, ticket := range *paginated.Data {
			orderIDInt := int(ticket.OrderID)
			graphTickets[i] = &graphmodel.TicketResponse{
				ID:         ticket.ID,
				EventID:    int(ticket.EventID),
				OrderID:    &orderIDInt,
				Price:      ticket.Price,
				Type:       ticket.Type,
				SeatNumber: ticket.SeatNumber,
			}
		}
	}
	return &graphmodel.TicketsResponse{
		Data:   graphTickets,
		Paging: (*graphmodel.PageMetadata)(paginated.Paging),
	}, nil
}

// SearchTickets is the resolver for the searchTickets field.
func (r *queryResolver) SearchTickets(ctx context.Context, id *string, eventID *int, orderID *int, price *float64, typeArg *string, seatNumber *string, page *int, size *int, sort *string, order *string) (*graphmodel.TicketsResponse, error) {
	defaultPage := 1
	defaultSize := 10
	defaultSort := "created_at"
	defaultOrder := "desc"

	var idStr string
	if id != nil {
		idStr = *id
	}

	var priceFloat float64
	if price != nil {
		priceFloat = *price
	}

	var typeStr string
	if typeArg != nil {
		typeStr = *typeArg
	}

	var seatNumberStr string
	if seatNumber != nil {
		seatNumberStr = *seatNumber
	}

	var eventIDUint, orderIDUint uint
	if eventID != nil {
		eventIDUint = uint(*eventID)
	}
	if orderID != nil {
		orderIDUint = uint(*orderID)
	}

	requestPage := defaultPage
	if page != nil {
		requestPage = *page
	}

	requestSize := defaultSize
	if size != nil {
		requestSize = *size
	}

	requestSort := defaultSort
	if sort != nil {
		requestSort = *sort
	}

	requestOrder := defaultOrder
	if order != nil {
		requestOrder = *order
	}

	paginated, err := r.TicketService.SearchTickets(ctx, &model.TicketSearchRequest{
		ID:         idStr,
		EventID:    eventIDUint,
		OrderID:    orderIDUint,
		Price:      priceFloat,
		Type:       typeStr,
		SeatNumber: seatNumberStr,
		Page:       requestPage,
		Size:       requestSize,
		Sort:       requestSort,
		Order:      requestOrder,
	})
	if err != nil {
		return nil, err
	}

	graphTickets := make([]*graphmodel.TicketResponse, len(*paginated.Data))
	if paginated.Data != nil {
		for i, ticket := range *paginated.Data {
			orderIDInt := int(ticket.OrderID)
			graphTickets[i] = &graphmodel.TicketResponse{
				ID:         ticket.ID,
				EventID:    int(ticket.EventID),
				OrderID:    &orderIDInt,
				Price:      ticket.Price,
				Type:       ticket.Type,
				SeatNumber: ticket.SeatNumber,
			}
		}
	}
	return &graphmodel.TicketsResponse{
		Data:   graphTickets,
		Paging: (*graphmodel.PageMetadata)(paginated.Paging),
	}, nil
}

// Profile is the resolver for the profile field.
func (r *queryResolver) Profile(ctx context.Context) (*model.UserResponse, error) {
	user, err := r.UserService.Profile(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Venue is the resolver for the venue field.
func (r *queryResolver) Venue(ctx context.Context, id int) (*model.VenueResponse, error) {
	venue, err := r.VenueService.GetVenueByID(ctx, &model.GetVenueRequest{
		ID: uint(id),
	})
	if err != nil {
		return nil, err
	}
	return venue, nil
}

// Venues is the resolver for the venues field.
func (r *queryResolver) Venues(ctx context.Context, page *int, size *int, sort *string, order *string) (*graphmodel.VenuesResponse, error) {
	defaultPage := 1
	defaultSize := 10
	defaultSort := "created_at"
	defaultOrder := "desc"

	requestPage := defaultPage
	if page != nil {
		requestPage = *page
	}

	requestSize := defaultSize
	if size != nil {
		requestSize = *size
	}

	requestSort := defaultSort
	if sort != nil {
		requestSort = *sort
	}

	requestOrder := defaultOrder
	if order != nil {
		requestOrder = *order
	}

	paginated, err := r.VenueService.GetVenues(ctx, &model.VenuesRequest{
		Page:  requestPage,
		Size:  requestSize,
		Sort:  requestSort,
		Order: requestOrder,
	})
	if err != nil {
		return nil, err
	}

	return &graphmodel.VenuesResponse{
		Data:   *paginated.Data,
		Paging: (*graphmodel.PageMetadata)(paginated.Paging),
	}, nil
}

// SearchVenues is the resolver for the searchVenues field.
func (r *queryResolver) SearchVenues(ctx context.Context, name *string, address *string, capacity *int, city *string, state *string, zip *string, page *int, size *int, sort *string, order *string) (*graphmodel.VenuesResponse, error) {
	defaultPage := 1
	defaultSize := 10
	defaultSort := "created_at"
	defaultOrder := "desc"

	// Handle optional string fields
	var nameStr, addressStr, cityStr, stateStr, zipStr string
	if name != nil {
		nameStr = *name
	}
	if address != nil {
		addressStr = *address
	}
	if city != nil {
		cityStr = *city
	}
	if state != nil {
		stateStr = *state
	}
	if zip != nil {
		zipStr = *zip
	}

	// Handle optional capacity
	var capacityVal int
	if capacity != nil {
		capacityVal = *capacity
	}

	requestPage := defaultPage
	if page != nil {
		requestPage = *page
	}

	requestSize := defaultSize
	if size != nil {
		requestSize = *size
	}

	requestSort := defaultSort
	if sort != nil {
		requestSort = *sort
	}

	requestOrder := defaultOrder
	if order != nil {
		requestOrder = *order
	}

	paginated, err := r.VenueService.SearchVenues(ctx, &model.VenueSearchRequest{
		Name:     nameStr,
		Address:  addressStr,
		Capacity: capacityVal,
		City:     cityStr,
		State:    stateStr,
		Zip:      zipStr,
		Page:     requestPage,
		Size:     requestSize,
		Sort:     requestSort,
		Order:    requestOrder,
	})
	if err != nil {
		return nil, err
	}

	return &graphmodel.VenuesResponse{
		Data:   *paginated.Data,
		Paging: (*graphmodel.PageMetadata)(paginated.Paging),
	}, nil
}

// CreatedAt is the resolver for the createdAt field.
func (r *userResponseResolver) CreatedAt(ctx context.Context, obj *model.UserResponse) (*time.Time, error) {
	if obj.CreatedAt == "" {
		return nil, nil
	}
	t, err := time.Parse("2006-01-02 15:04:05.999999 -0700 MST", obj.CreatedAt)
	if err != nil {
		t, err = time.Parse("2006-01-02 15:04:05 -0700 MST", obj.CreatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &t, nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *userResponseResolver) UpdatedAt(ctx context.Context, obj *model.UserResponse) (*time.Time, error) {
	if obj.UpdatedAt == "" {
		return nil, nil
	}
	t, err := time.Parse("2006-01-02 15:04:05.999999 -0700 MST", obj.UpdatedAt)
	if err != nil {
		t, err = time.Parse("2006-01-02 15:04:05 -0700 MST", obj.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &t, nil
}

// ID is the resolver for the id field.
func (r *venueResponseResolver) ID(ctx context.Context, obj *model.VenueResponse) (int, error) {
	return int(obj.ID), nil
}

// EventResponse returns graph.EventResponseResolver implementation.
func (r *Resolver) EventResponse() graph.EventResponseResolver { return &eventResponseResolver{r} }

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

// UserResponse returns graph.UserResponseResolver implementation.
func (r *Resolver) UserResponse() graph.UserResponseResolver { return &userResponseResolver{r} }

// VenueResponse returns graph.VenueResponseResolver implementation.
func (r *Resolver) VenueResponse() graph.VenueResponseResolver { return &venueResponseResolver{r} }

type eventResponseResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResponseResolver struct{ *Resolver }
type venueResponseResolver struct{ *Resolver }