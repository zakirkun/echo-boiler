package infrastructure

import (
	"log"

	"github.com/zakirkun/echo-boiler/internal/config"
)

type Configuration interface {
	TomlConfiguration()
}

type configurationContext struct {
	Configfile string
}

func NewConfiguration(filename string) Configuration {
	return configurationContext{Configfile: filename}
}

func (c configurationContext) TomlConfiguration() {
	if err := config.Initialize(c.Configfile); err != nil {
		log.Fatalf("Error reading configuration: %s\n", err.Error())
	}
}
