package main

import (
	"mailingAPI/cmd/config"
	"mailingAPI/internal/app"
)

func main() {
	cfg := config.NewServerConfig()
	srv := app.NewServerApp(cfg)
	srv.Run()
}
