package converter

import (
	"github.com/TrinityKnights/Backend/internal/domain/entity"
	"github.com/TrinityKnights/Backend/internal/domain/model"
)

func EventToResponse(event *entity.Event) *model.EventResponse {
	return &model.EventResponse{
		ID:        event.ID,
		Name:      event.Name,
		Date:      event.Date.String(),
		Time:      event.Time,
		VenueID:   event.VenueID,
		Venue:     entity.Venue{
			ID:    event.Venue.ID,
			Name:  event.Venue.Name,
			Address:  event.Venue.Address,
			Capacity: event.Venue.Capacity,
			City:  event.Venue.City,
			State: event.Venue.State,
			Zip:   event.Venue.Zip,
		},	
		CreatedAt: event.CreatedAt.String(),
		UpdatedAt: event.UpdatedAt.String(),
	}
}
