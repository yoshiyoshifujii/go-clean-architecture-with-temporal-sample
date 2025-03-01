package api

import "github.com/labstack/echo/v4"

type (
	Server struct {
		echo *echo.Echo
	}
)

func NewServer() *Server {
	e := echo.New()
	RegisterRoutes(e)
	return &Server{echo: e}
}

func (s *Server) Start(address string) error {
	return s.echo.Start(address)
}
