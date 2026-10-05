package booking

import (
	"go_tickets/internal/domain/booking/dto"
	"go_tickets/internal/domain/event"

	"github.com/google/uuid"
)

func generateBookingCode() string {
	return "GT" + uuid.New().String()
}

type service struct {
	bookingRepo Repository
	eventRepo   event.Repository
}

func NewService(bookingRepo Repository, eventRepo event.Repository) *service {
	return &service{
		bookingRepo: bookingRepo,
		eventRepo:   eventRepo,
	}
}

func (s *service) CreateBooking(userId uint, req dto.CreateBookingRequest) (*dto.ResponseBooking, error) {
	booking, err := s.bookingRepo.CreateBookingWithTicketUpdate(userId, req.EventID, req.Quantity)
	if err != nil {
		return nil, err
	}
	return booking.BookingCreatedResponse(), nil
}

func (s *service) GetMyBookings(userId uint) ([]*dto.ResponseBooking, error) {
	bookings, err := s.bookingRepo.GetByUserId(userId)
	if err != nil {
		return nil, err
	}
	var responses = make([]*dto.ResponseBooking, len(bookings))
	for i, booking := range bookings {
		responses[i] = booking.BookingCreatedResponse()
	}
	return responses, nil
}
