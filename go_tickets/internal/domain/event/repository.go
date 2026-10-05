package event

import (
	"errors"

	"gorm.io/gorm"
)

var ErrEventNotFount = errors.New("Event not found")

type Repository interface {
	Create(event *Event) error
	GetAll() ([]*Event, error)
	GetById(eventId uint) (*Event, error)
	Update(event *Event) error
	HardDelete(eventId uint) error
	SoftDelete(eventId uint) (*Event, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(event *Event) error {
	return r.db.Create(event).Error
}

func (r *repository) GetAll() ([]*Event, error) {
	var event []*Event
	if err := r.db.Find(&event).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrEventNotFount
		}
		return nil, err
	}
	return event, nil
}

func (r *repository) GetById(eventId uint) (*Event, error) {
	var event Event
	if err := r.db.First(&event, eventId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrEventNotFount
		}
		return nil, err
	}
	return &event, nil
}

func (r *repository) Update(event *Event) error {

	return r.db.Save(event).Error
}

func (r *repository) HardDelete(eventId uint) error {
	var event Event
	return r.db.Unscoped().Delete(&event, uint(eventId)).Error
}

func (r *repository) SoftDelete(eventId uint) (*Event, error) {
	var event Event
	if err := r.db.First(&event, eventId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrEventNotFount
		}
		return nil, err
	}
	if err := r.db.Delete(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}
