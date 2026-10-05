package booking

import (
	"go_tickets/internal/domain/booking/dto"

	"gorm.io/gorm"
)

const (
	StatusPending   = "pending"
	StatusConfirmed = "confirmed"
	StatusCancelled = "cancelled"
)

type Booking struct {
	gorm.Model
	UserID      uint   `json:"user_id" gorm:"not null"`
	EventID     uint   `json:"event_id" gorm:"not null"`
	Quantity    int    `json:"quantity" gorm:"not null"`
	TotalPrice  int    `json:"total_price" gorm:"not null"`
	Status      string `json:"status" gorm:"type:varchar(50);not null"`
	BookingCode string `json:"booking_code" gorm:"uniqueIndex;not null"`
}

func (b *Booking) BookingCreatedResponse() *dto.ResponseBooking {
	return &dto.ResponseBooking{
		UserID:      b.UserID,
		Quantity:    b.Quantity,
		TotalPrice:  b.TotalPrice,
		Status:      b.Status,
		BookingCode: b.BookingCode,
		CreatedAt:   b.CreatedAt.String(),
	}
}
