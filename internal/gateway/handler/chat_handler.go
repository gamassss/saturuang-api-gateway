package handler

import (
	"github.com/labstack/echo/v4"
	"saturuang-api-gateway/internal/gateway/service"
)

func ChatServiceHandler(c echo.Context) error {
	return service.ForwardRequest(c, "http://localhost:3002")
}
