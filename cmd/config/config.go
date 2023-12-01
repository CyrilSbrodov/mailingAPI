package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/caarlos0/env/v6"
)

// ServerConfig структура конфига для сервера.
type ServerConfig struct {
	Addr          string `json:"address" env:"ADDRESS"`
	DatabaseDSN   string `json:"database_dsn" env:"DATABASE_DSN"`
	TrustedSubnet string `json:"trusted_subnet" env:"TRUSTED_SUBNET"`
}

// ServerConfigInit инициализация конфига.
func ServerConfigInit() *ServerConfig {
	cfgSrv := &ServerConfig{}
	flag.StringVar(&cfgSrv.Addr, "a", "localhost:8080", "ADDRESS")
	flag.StringVar(&cfgSrv.DatabaseDSN, "d", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable", "DATABASE_DSN")
	flag.StringVar(&cfgSrv.TrustedSubnet, "t", "", "CIDR")
	flag.Parse()
	if err := env.Parse(cfgSrv); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return cfgSrv
}
