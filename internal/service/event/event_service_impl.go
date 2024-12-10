package event

import (
	"context"
	"errors"
	"net/http"

	"github.com/TrinityKnights/Backend/internal/domain/entity"
	"github.com/TrinityKnights/Backend/internal/domain/model"
	"github.com/TrinityKnights/Backend/internal/domain/model/converter"
	"github.com/TrinityKnights/Backend/internal/repository/event"
	"github.com/TrinityKnights/Backend/pkg/helper"
	"github.com/TrinityKnights/Backend/pkg/jwt"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type EventServiceImpl struct {
	DB         *gorm.DB
	Log        *logrus.Logger
	Validate   *validator.Validate
	EventRepository *event.EventRepositoryImpl
	JWTService jwt.JWTService
	helper     *helper.ContextHelper
}

func NewEventServiceImpl(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, eventRepository *event.EventRepositoryImpl, jwtService jwt.JWTService) *EventServiceImpl {
	return &EventServiceImpl{
		DB:         db,
		Log:        log,
		Validate:   validate,
		EventRepository: eventRepository,
		helper:     helper.NewContextHelper(),
	}
}

func (s *EventServiceImpl) GetEventWithDetails(ctx context.Context, id uint) (*model.EventResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	data := &entity.Event{}

	if err := s.EventRepository.FindById(tx, id, data); err != nil {
		s.Log.Errorf("failed to get event by id: %v", err)
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.Errorf("failed to commit transaction: %v", err)
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	return converter.EventToResponse(data), nil
}
