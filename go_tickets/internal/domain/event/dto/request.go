package dto

import "time"

type CreateRequest struct {
	TITLE           string    `json:"title" validate:"required,min=2,max=150"`
	DESCRIPTION     string    `json:"description" validate:"omitempty,max=1000"`
	LOCATION        string    `json:"location" validate:"required"`
	STARTSAT        time.Time `json:"starts_at" validate:"required"`
	TOTALTICKETS    int       `json:"total_tickets" validate:"required,gt=0"`
	PRICE           int       `json:"price" validate:"gte=0"`
}
type UpdateRequest struct {
	TITLE       string `json:"title" validate:"min=2,max=150"`
	DESCRIPTION string `json:"description" validate:"omitempty,max=1000"`
	LOCATION    string `json:"location" validate:"omitempty,max=200" `
	STARTSAT    time.Time `json:"starts_at" validate:"omitempty"`
	PRICE       int    `json:"price" validate:"gte=0"`
}
