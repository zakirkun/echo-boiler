package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/zakirkun/echo-boiler/internal/server"
)

type Infrastructure interface {
	WebServer()
}

type infrastructureContext struct {
	Server *echo.Echo
}

func NewInfrastructure(
	server *echo.Echo,
) Infrastructure {
	return infrastructureContext{
		Server: server,
	}
}

func (i infrastructureContext) WebServer() {
	s := server.New(i.Server)
	s.Start()
}
