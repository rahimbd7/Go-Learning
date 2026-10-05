package booking

import (
	"errors"
	"fmt"
	"go_tickets/internal/domain/booking/dto"
	"go_tickets/internal/domain/event"
	"go_tickets/internal/httpresponse"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	bookingService *service
}

func NewHandler(bookingService *service) *handler {
	return &handler{
		bookingService: bookingService,
	}
}

func bookingErrorResponse(c *echo.Context, err error) error {
	if errors.Is(err, ErrBookingNotFound) {
		return c.JSON(http.StatusNotFound, httpresponse.Error{
			Code:    http.StatusNotFound,
			Message: "Booking not found",
		})
	}

	if errors.Is(err, event.ErrEventNotFount) {
		return c.JSON(http.StatusNotFound, httpresponse.Error{
			Code:    http.StatusNotFound,
			Message: "Event not found",
		})
	}

	if errors.Is(err, ErrNotEnoughTickets) {
		return c.JSON(http.StatusConflict, httpresponse.Error{
			Code:    http.StatusConflict,
			Message: "Not enough tickets available",
		})
	}

	if errors.Is(err, ErrBookingAlreadyCancelled) {
		return c.JSON(http.StatusConflict, httpresponse.Error{
			Code:    http.StatusConflict,
			Message: "Booking is already cancelled",
		})
	}

	if errors.Is(err, ErrForbiddenBookingAccess) {
		return c.JSON(http.StatusForbidden, httpresponse.Error{
			Code:    http.StatusForbidden,
			Message: "You do not own this booking",
		})
	}

	return c.JSON(http.StatusInternalServerError, httpresponse.Error{
		Code:    http.StatusInternalServerError,
		Message: "Something went wrong",
		Details: err.Error(),
	})
}

func (h *handler) CreateBooking(c *echo.Context) error {
	userId, ok := c.Get("user_id").(uint)
	fmt.Printf("user_id value=%v type=%T\n", userId, userId)
	if !ok {
		return c.JSON(http.StatusUnauthorized, httpresponse.Error{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}
	var req dto.CreateBookingRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request",
			Details: err.Error(),
		})
	}
	booking, err := h.bookingService.CreateBooking(userId, req)
	if err != nil {
		return bookingErrorResponse(c, err)
	}
	return c.JSON(http.StatusCreated, booking)
}

func (h *handler) GetMyBookings(c *echo.Context) error {
	userId, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, httpresponse.Error{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}
	bookings, err := h.bookingService.GetMyBookings(userId)
	if err != nil {
		return bookingErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, bookings)
}
