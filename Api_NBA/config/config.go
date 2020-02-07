package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server   string
	Database string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("/go/src/github.com/johnatansantana/Api_NBA/config/config.toml", &c); err != nil {
		log.Fatal(err)
	}

}
