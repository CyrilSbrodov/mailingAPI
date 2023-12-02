package main

import (
	"mailingAPI/cmd/config"
	"mailingAPI/internal/app"
)

// @title mailing API
// @version 1.0
// @description API for mailing service

// @host localhost:8080
// @BasePath /

func main() {
	cfg := config.NewServerConfig()
	srv := app.NewServerApp(cfg)
	srv.Run()
}
