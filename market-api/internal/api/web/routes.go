package web

import (
	"github.com/labstack/echo/v4"
)

func InitRoutes(s *Server, e *echo.Echo) {
	e.POST("/markets", func(c echo.Context) error {
		return CreateMarket(s, c)
	})
	e.GET("/markets/:id", func(c echo.Context) error {
		return GetMarketByID(s, c)
	})
	e.GET("/markets", func(c echo.Context) error {
		return GetMarkets(s, c)
	})
	e.PUT("/markets/:id", func(c echo.Context) error {
		return UpdateMarket(s, c)
	})
	e.DELETE("/markets", func(c echo.Context) error {
		return DeleteMarket(s, c)
	})
}
