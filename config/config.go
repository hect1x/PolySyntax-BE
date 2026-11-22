package config

import (
	"os"

	"github.com/joho/godotenv"
)


type Config struct {
	Addr string
	Env  string
}

func Load() *Config {
	godotenv.Load()

	addr := os.Getenv("PORT")
	if addr == "" {
		addr = ":8080"
	} else {
		addr = ":" + addr
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	return &Config{
		Addr: addr,
		Env:  env,
	}
}
