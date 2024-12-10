package event

import (
	"github.com/TrinityKnights/Backend/internal/domain/entity"
	"gorm.io/gorm"
)

type EventRepository interface {
	FindAll(db *gorm.DB, events []*entity.Event) error
	FindById(db *gorm.DB,id uint,event *entity.Event) error
	GetAllEventsWithDetails(db *gorm.DB,events []*entity.Event) error
}