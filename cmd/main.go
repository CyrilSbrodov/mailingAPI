package main

import (
	"mailingAPI/cmd/config"
	"mailingAPI/internal/app"
)

func main() {
	cfg := config.ServerConfigInit()
	srv := app.NewServerApp(cfg)
	srv.Run()
}
