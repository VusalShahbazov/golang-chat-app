package api

import (
	"back/internal/helpers"
	"os"
)

type Config struct {
	BindAddr string
}


func NewConfig() *Config {
	bindAddr := helpers.GetOrDefault(os.Getenv("BIND_ADDR"), ":3000")
	return &Config{
		BindAddr: bindAddr,
	}
}






