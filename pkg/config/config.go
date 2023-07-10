package config

import (
	"github.com/caarlos0/env/v6"
)

type Configuration struct {
	RPCURL   string `env:"RPC_URL" envDefault:"https://cosmos-rpc.quickapi.com:443"`
	DBURL    string `env:"DB_URL" envDefault:"postgres://postgres@localhost:5432/dumptxs?sslmode=disable"`
	FILEPATH string `env:"FILE_DUMP_PATH" envDefault:"./"`
}

// MustGetConfig returns a configuration
func MustGetConfig() *Configuration {
	cfg := &Configuration{}
	err := env.Parse(cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
