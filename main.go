package main

import (
	"flag"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/zakirkun/echo-boiler/infrastructure"
	"github.com/zakirkun/echo-boiler/router"
)

var cfgFile *string

func init() {
	cfgFile = flag.String("c", "config.toml", "Configuration")
}

func main() {
	// load config
	cfgBundle := infrastructure.NewConfiguration(*cfgFile)
	cfgBundle.TomlConfiguration()

	application := infrastructure.NewInfrastructure(setServer())
	application.WebServer()
}

func setServer() *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	router.Routing(e)

	return e
}
