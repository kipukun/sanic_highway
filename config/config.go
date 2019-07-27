package config

import (
	"io"
	"log"

	"github.com/BurntSushi/toml"
)

var DefaultConfig = Config{
	DB: database{
		Driver: "postgres",
		DSN:    "postgres://postgres:cock@127.0.0.1/test?sslmode=disable",
	},
	Web: web{
		Addr: ":1337",
	},
}

type Config struct {
	DB  database
	Web web
}

type database struct {
	Driver, DSN string
}

type web struct {
	Addr string
}

func Load(r io.Reader) (Config, error) {
	var c = DefaultConfig
	m, err := toml.DecodeReader(r, &c)
	if err != nil {
		return DefaultConfig, err
	}

	for _, key := range m.Undecoded() {
		log.Printf("[*] unknown config fields: %s", key)
	}

	return c, nil
}
