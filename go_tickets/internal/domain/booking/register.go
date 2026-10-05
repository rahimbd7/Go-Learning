package booking

import (
	"go_tickets/internal/auth"
	"go_tickets/internal/config"
	"go_tickets/internal/domain/event"
	"go_tickets/internal/middleware"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	jwtService := auth.NewJWTService(cfg.JWTSecret)
	bookingRepo := NewRepository(db)
	eventRepo := event.NewRepository(db)
	bookingService := NewService(bookingRepo, eventRepo)
	bookingHandler := NewHandler(bookingService)

	api := e.Group("/api/v1/bookings", middleware.AuthMiddleWare(jwtService))

	api.POST("", bookingHandler.CreateBooking)
	api.GET("", bookingHandler.GetMyBookings)

}
