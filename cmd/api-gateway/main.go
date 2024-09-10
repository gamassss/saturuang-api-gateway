package main

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"log"
	internalMiddleware "saturuang-api-gateway/internal/gateway/middleware"
)

func main() {
	e := echo.New()

	// Logger and recovery
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())

	// Internal middleware
	e.Use(internalMiddleware.JWTMiddleware())

	log.Println("Starting API Gateway on Port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
