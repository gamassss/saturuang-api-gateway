package config

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"saturuang-api-gateway/pkg/jwtutils"
)

func JWTConfig() echojwt.Config {
	return echojwt.Config{
		SigningKey:  jwtutils.SecretKey,
		TokenLookup: "header:Authorization:Bearer",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &jwtutils.CustomClaims{}
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Invalid or missing token",
			})
		},
		ContinueOnIgnoredError: false,
	}
}
