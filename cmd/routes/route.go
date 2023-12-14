package routes

import (
	"github.com/fazaalexander/montirku-be/common"
	"github.com/fazaalexander/montirku-be/middleware/log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartRoute(handler common.Handler) *echo.Echo {
	e := echo.New()
	log.LogMiddleware(e)
	e.Use(middleware.CORS())

	handler.AuthHandler.RegisterRoutes(e)
	handler.BengkelHandler.RegisterRoutes(e)
	handler.TransactionHandler.RegisterRoutes(e)

	return e
}
