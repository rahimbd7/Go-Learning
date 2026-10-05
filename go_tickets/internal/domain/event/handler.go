package event

import (
	"errors"
	"go_tickets/internal/domain/event/dto"
	"go_tickets/internal/httpresponse"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(s *service) *handler {
	return &handler{service: s}
}

func eventErrorResponse(c *echo.Context, err error) error {
	if errors.Is(err, ErrEventNotFount) {
		return c.JSON(http.StatusNotFound, httpresponse.Error{
			Code:    http.StatusNotFound,
			Message: "Event not found",
		})
	}

	return c.JSON(http.StatusInternalServerError, httpresponse.Error{
		Code:    http.StatusInternalServerError,
		Message: "Something went wrong",
		Details: err.Error(),
	})
}

func (h *handler) CreateEvent(c *echo.Context) error {
	var req dto.CreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadGateway,
			Message: "Invalid Request Payload",
			Details: err.Error(),
		})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}
	response, err := h.service.CreateEvent(req)
	if err != nil {
		return eventErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, response)
}
func (h *handler) GetAllEvent(c *echo.Context) error {
	events, err := h.service.GetAllEvent()
	if err != nil {
		return eventErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, events)
}
func (h *handler) GetEventById(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid event id",
			Details: err.Error(),
		})
	}
	response, err := h.service.GetById(uint(id))
	if err != nil {
		return eventErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, response)
}

func (h *handler) Update(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid event id",
			Details: err.Error(),
		})
	}
	var req dto.UpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	response, err := h.service.Update(uint(id), req)
	if err != nil {
		return eventErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, response)
}

func (h *handler) HardDelete(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid event id",
			Details: err.Error(),
		})
	}
	if err := h.service.HardDelete(uint(id)); err != nil {
		return eventErrorResponse(c, err)
	}
	return nil

}
func (h *handler) SoftDelete(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid event id",
			Details: err.Error(),
		})
	}
	response, err := h.service.SoftDelete(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Failed to delete",
			Details: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response)

}
