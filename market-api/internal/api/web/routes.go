package web

import (
	"github.com/labstack/echo/v4"
)

func InitRoutes(s *Server, e *echo.Echo) {
	e.POST("/v1/markets", func(c echo.Context) error {
		return CreateMarket(s, c)
	})
	e.GET("/v1/markets/:id", func(c echo.Context) error {
		return GetMarketByID(s, c)
	})
	e.GET("/v1/markets", func(c echo.Context) error {
		return GetMarkets(s, c)
	})
	e.PUT("/v1/markets/:id", func(c echo.Context) error {
		return UpdateMarket(s, c)
	})
	e.DELETE("/v1/markets", func(c echo.Context) error {
		return DeleteMarket(s, c)
	})
}
