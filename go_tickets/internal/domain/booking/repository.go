package booking

import (
	"errors"
	"go_tickets/internal/domain/event"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrBookingNotFound         = errors.New("booking not found")
	ErrNotEnoughTickets        = errors.New("not enough tickets available")
	ErrBookingAlreadyCancelled = errors.New("booking is already cancelled")
	ErrForbiddenBookingAccess  = errors.New("forbidden to update booking")
)

type Repository interface {
	CreateBooking(booking *Booking) error
	GetById(id uint) (*Booking, error)
	GetByUserId(userId uint) ([]*Booking, error)
	UpdateBooking(booking *Booking) error
	DeleteBooking(booking *Booking) error
	CreateBookingWithTicketUpdate(userId uint, eventId uint, quantity int) (*Booking, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateBooking(booking *Booking) error {
	return r.db.Create(booking).Error
}

func (r *repository) GetById(id uint) (*Booking, error) {
	var booking Booking
	err := r.db.First(&booking, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBookingNotFound
		}
		return nil, err
	}
	return &booking, nil
}
func (r *repository) GetByUserId(userId uint) ([]*Booking, error) {
	var bookings []*Booking
	err := r.db.Where("user_id = ?", userId).Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	return bookings, nil
}
func (r *repository) UpdateBooking(booking *Booking) error {
	return r.db.Save(booking).Error
}

func (r *repository) DeleteBooking(booking *Booking) error {
	return r.db.Delete(booking).Error
}

func (r *repository) CreateBookingWithTicketUpdate(userId uint, eventId uint, quantity int) (*Booking, error) {
	var booking Booking
	err := r.db.Transaction(func(tx *gorm.DB) error {
		var eventData event.Event
		err := tx.Clauses(clause.Locking{Strength: "Update"}).First(&eventData, eventId).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrBookingNotFound
			}
			return err
		}
		if eventData.AVAILABLETICKET < int(quantity) {
			return ErrNotEnoughTickets
		}
		booking = Booking{
			UserID:      userId,
			EventID:     eventData.ID,
			Quantity:    quantity,
			Status:      StatusConfirmed,
			TotalPrice:  eventData.PRICE * quantity,
			BookingCode: generateBookingCode(),
		}
		if err := tx.Create(&booking).Error; err != nil {
			return err
		}
		eventData.AVAILABLETICKET -= int(quantity)
		if err := tx.Save(&eventData).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &booking, nil
}
