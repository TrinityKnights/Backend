package model

import "github.com/TrinityKnights/Backend/internal/domain/entity"

type EventResponse struct {
	ID          uint         `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Date        string       `json:"date"`
	Time        string       `json:"time"`
	VenueID     uint         `json:"venue_id"`
	Venue       entity.Venue `json:"venue"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
}
