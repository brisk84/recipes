package config

import (
	"embed"
	"encoding/json"
	"flag"

	"github.com/caarlos0/env/v6"
)

//go:embed settings.json
var fsett embed.FS

type Config struct {
	AppAddr string `env:"SERVER_ADDRESS" envDefault:":80" json:"server_address"`
	PgURI   string `env:"PG_URI"                          json:"pg_uri"`
}

func New() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return Config{}, nil
	}

	cnt, err := fsett.ReadFile("settings.json")
	if err != nil {
		return Config{}, err
	}
	err = json.Unmarshal(cnt, &cfg)
	if err != nil {
		return Config{}, err
	}

	addr := flag.String("a", cfg.AppAddr, "-a localhost:80")
	flag.Parse()
	cfg.AppAddr = *addr

	return cfg, nil
}
