package event

import (
	"time"

	"go_tickets/internal/domain/event/dto"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	TITLE           string    `json:"title" gorm:"type:varchar(150);not null"`
	DESCRIPTION     string    `json:"description" gorm:"type:text"`
	LOCATION        string    `json:"location" gorm:"type:varchar(150);not null"`
	STARTSAT        time.Time `json:"starts_at" gorm:"not null"`
	TOTALTICKETS    int       `json:"total_tickets" gorm:"not null"`
	AVAILABLETICKET int       `json:"available_tickets" gorm:"not null"`
	PRICE           int       `json:"price" gorm:"not null"`
}

func (e *Event) ToResponse() *dto.Response {
	return &dto.Response{
		ID:              e.ID,
		TITLE:           e.TITLE,
		DESCRIPTION:     e.DESCRIPTION,
		LOCATION:        e.LOCATION,
		STARTSAT:        e.STARTSAT,
		TOTALTICKETS:    e.TOTALTICKETS,
		AVAILABLETICKET: e.AVAILABLETICKET,
		PRICE:           e.PRICE,
		CREATEDAT:       e.CreatedAt.String(),
	}
}
func (e *Event) ToDeletedResponse() *dto.DeleteResponse {
	return &dto.DeleteResponse{
		ID:              e.ID,
		TITLE:           e.TITLE,
		DESCRIPTION:     e.DESCRIPTION,
		LOCATION:        e.LOCATION,
		STARTSAT:        e.STARTSAT,
		TOTALTICKETS:    e.TOTALTICKETS,
		AVAILABLETICKET: e.AVAILABLETICKET,
		PRICE:           e.PRICE,
		DELETEDAT:       e.DeletedAt.Time.String(),
	}
}
