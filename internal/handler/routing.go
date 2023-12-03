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
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"}, // Replace with your frontend origin(s)
			AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		}),
	)

	menu := e.Group("/menu")
	menu.GET("/:id", menuHandler.GetById)
	menu.GET("/flavors", menuHandler.GetAll)
}
