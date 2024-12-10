package event

import (
	"context"

	"github.com/TrinityKnights/Backend/internal/domain/model"
)

type EventService interface {
	GetEventWithDetails(ctx context.Context, id uint) (*model.EventResponse, error)
}
