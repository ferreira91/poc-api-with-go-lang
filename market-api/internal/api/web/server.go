package web

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"market-api/internal/core/domain"
)

type Server struct {
	Service domain.IMarketService
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) InitWebServer(port string) {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}
	InitRoutes(s, e)

	e.Logger.Fatal(e.Start(port))
}
