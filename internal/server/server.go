package server

import (
	"github.com/labstack/echo"
	"github.com/zakirkun/echo-boiler/internal/config"
)

type IServer interface {
	Start()
}

type Server struct {
	srv *echo.Echo
}

func New(e *echo.Echo) IServer {
	return Server{srv: e}
}

func (s Server) Start() {
	s.srv.Logger.Fatal(s.srv.Start(":" + config.GetString("server.port")))
}
