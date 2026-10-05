package user

import (
	"go_tickets/internal/auth"
	"go_tickets/internal/config"
	"go_tickets/internal/middleware"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	jwtService := auth.NewJWTService(cfg.JWTSecret)
	userRepository := NewRepository(db)
	userService := NewService(userRepository, jwtService)
	userHandler := NewHandler(userService)

	api := e.Group("/api/v1/auth")

	api.POST("/register", userHandler.CreateUser)                         // api/v1/auth/register
	api.POST("/login", userHandler.LoginUser)                             // api/v1/auth/login
	api.GET("/me", userHandler.Me, middleware.AuthMiddleWare(jwtService)) // api/v1/auth/login
}
