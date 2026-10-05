package event

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	handler := NewHandler(svc)

	api := e.Group("/api/v1/events")
	api.POST("", handler.CreateEvent)
	api.GET("", handler.GetAllEvent)
	api.GET("/:id",handler.GetEventById)
	api.PATCH("/:id",handler.Update)
	api.DELETE("/:id",handler.HardDelete)
	api.DELETE("/softdelete/:id",handler.SoftDelete)
}
