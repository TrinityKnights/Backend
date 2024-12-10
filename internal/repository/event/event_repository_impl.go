package event

import (
	"github.com/TrinityKnights/Backend/internal/domain/entity"
	"github.com/TrinityKnights/Backend/internal/repository"
	"gorm.io/gorm"
)

type EventRepositoryImpl struct {
	repository.RepositoryImpl[entity.Event]
}

func NewEventRepository(db *gorm.DB) *EventRepositoryImpl {
	return &EventRepositoryImpl{
		RepositoryImpl: repository.RepositoryImpl[entity.Event]{DB: db},
	}
}

func (r *EventRepositoryImpl) FindAll(db *gorm.DB, events []*entity.Event) error {
	return db.Preload("Venue").Find(&events).Error
}

func (r *EventRepositoryImpl) FindById(db *gorm.DB,id uint ,event *entity.Event) error {
	return db.Where("id = ?", id).Preload("Venue").First(&event).Error
}
