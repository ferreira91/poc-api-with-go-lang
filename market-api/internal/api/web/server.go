package web

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"market-api/internal/core/domain"
	"market-api/utils"
)

type Server struct {
	Service domain.IMarketService
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) InitWebServer(port string) error {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}
	InitRoutes(s, e)

	err := e.Start(port)
	if err != nil {
		return err
	}
	utils.LoggerInfo("Web server is up")
	return nil
}
