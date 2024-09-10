package middleware

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"saturuang-api-gateway/pkg/config"
)

func ErrorHandler(c echo.Context, err error) error {
	return c.JSON(http.StatusUnauthorized, map[string]string{
		"error": "Invalid or missing token",
	})
}

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(config.JWTConfig())
}
