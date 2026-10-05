package dto

import "time"

type Response struct {
	ID              uint      `json:"id"`
	TITLE           string    `json:"title"`
	DESCRIPTION     string    `json:"description"`
	LOCATION        string    `json:"location"`
	STARTSAT        time.Time `json:"starts_at"`
	TOTALTICKETS    int       `json:"total_tickets"`
	AVAILABLETICKET int       `json:"available_tickets"`
	PRICE           int       `json:"price"`
	CREATEDAT       string    `json:"created_at"`
}
type DeleteResponse struct {
	ID              uint      `json:"id"`
	TITLE           string    `json:"title"`
	DESCRIPTION     string    `json:"description"`
	LOCATION        string    `json:"location"`
	STARTSAT        time.Time `json:"starts_at"`
	TOTALTICKETS    int       `json:"total_tickets"`
	AVAILABLETICKET int       `json:"available_tickets"`
	PRICE           int       `json:"price"`
	DELETEDAT       string    `json:"deleted_at"`
}
