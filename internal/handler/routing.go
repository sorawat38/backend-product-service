package handler

import (
	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/handler/menuhdl"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, menuHandler menuhdl.HTTPHandler) {

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.RequestID(),
	)

	menu := e.Group("/menu")
	menu.GET("/:id", nil)
	menu.GET("/flavours", menuHandler.GetAll)
}
