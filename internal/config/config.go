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
	AppAddr    string `env:"SERVER_ADDRESS" envDefault:":80"   json:"server_address"`
	PgURI      string `env:"PG_URI"                            json:"pg_uri"`
	RedisAddr  string `env:"REDIS_ADDR"     envDefault:":6379" json:"redis_addr"`
	MinioAddr  string `env:"MINIO_ADDR"     envDefault:":9000" json:"minio_addr"`
	MinioLogin string `env:"MINIO_LOGIN"                       json:"minio_login"`
	MinioPass  string `env:"MINIO_PASS"                        json:"minio_pass"`
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
