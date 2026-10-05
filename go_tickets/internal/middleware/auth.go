package middleware

import (
	"fmt"
	"go_tickets/internal/auth"
	"net/http"
	"strings"
	"github.com/labstack/echo/v5"
)

func AuthMiddleWare(jwtService auth.JWTService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			//extract token from header
			authToken := c.Request().Header.Get("Authorization")
			if authToken == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"err": "Missing unauthorized header",
				})
			}
			//check bearer exists
			parts := strings.Split(authToken, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"err": "invalid unauthorized header",
				})
			}
			token := parts[1]
			// varify token
			claims, err := jwtService.ValidateToken(token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"err": "you are not authorized",
				})
			}
			fmt.Println(claims)
			c.Set("user_id", claims.UserID)
			c.Set("user_email", claims.Email)
			c.Set("user_name", claims.Name)
			return next(c)

		}
	}
}
