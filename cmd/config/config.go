package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/caarlos0/env/v6"
)

// ServerConfig структура конфига для сервера.
type ServerConfig struct {
	Addr          string        `json:"address" env:"ADDRESS"`
	DatabaseDSN   string        `json:"database_dsn" env:"DATABASE_DSN"`
	ExternalAddr  string        `json:"external_addr" env:"EXTERNAL_ADDRESS"`
	Token         string        `json:"token" env:"TOKEN"`
	CheckInterval time.Duration `json:"check_interval" env:"CHECK_INTERVAL"`
}

// NewServerConfig инициализация конфига.
func NewServerConfig() *ServerConfig {
	cfgSrv := &ServerConfig{}
	flag.StringVar(&cfgSrv.Addr, "a", "localhost:8080", "ADDRESS")
	flag.StringVar(&cfgSrv.ExternalAddr, "e", "https://probe.fbrq.cloud/v1/send/", "external address gor send mailing")
	flag.StringVar(&cfgSrv.DatabaseDSN, "d", "postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable", "DATABASE_DSN")
	flag.StringVar(&cfgSrv.Token, "token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzI3OTk0MzksImlzcyI6ImZhYnJpcXVlIiwibmFtZSI6Imh0dHBzOi8vdC5tZS9jeXJpbHNicm9kb3YifQ.UEIeNXd517YYWYMWjKoRXzYI0VuspTR8irWd82Bb1qM", "JWT Token for external addr")
	flag.DurationVar(&cfgSrv.CheckInterval, "c", 5*time.Second, "time interval for checking a new mailing list")
	flag.Parse()
	if err := env.Parse(cfgSrv); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return cfgSrv
}
