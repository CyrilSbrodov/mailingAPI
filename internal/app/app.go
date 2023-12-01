package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

	"mailingAPI/cmd/config"
	"mailingAPI/cmd/loggers"
	"mailingAPI/internal/handlers"
	"mailingAPI/internal/storage/repositories"
	"mailingAPI/pkg/postgres"
)

type ServerApp struct {
	cfg    *config.ServerConfig
	logger *loggers.Logger
	router *chi.Mux
}

func NewServerApp(cfg *config.ServerConfig) *ServerApp {
	router := chi.NewRouter()
	logger := loggers.NewLogger()

	return &ServerApp{
		router: router,
		cfg:    cfg,
		logger: logger,
	}
}

func (a *ServerApp) Run() {
	var err error
	//определение хендлера
	client, err := postgres.NewClient(context.Background(), 5, a.cfg, a.logger)
	if err != nil {
		a.logger.LogErr(err, "")
		os.Exit(1)
	}
	store, err := repositories.NewPGStore(client, a.cfg, a.logger)
	if err != nil {
		a.logger.LogErr(err, "")
		os.Exit(1)
	}

	handler := handlers.NewHandler(a.cfg, a.logger, store)
	//регистрация хендлера
	handler.Register(a.router)

	srv := http.Server{
		Addr:    a.cfg.Addr,
		Handler: a.router,
	}
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.LogErr(err, "server not started")
		}
	}()
	a.logger.LogInfo("server is listen:", a.cfg.Addr, "start server")

	//gracefullshutdown
	<-done

	a.logger.LogInfo("", "", "server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctx); err != nil {
		a.logger.LogErr(err, "Server Shutdown Failed")
	}
	a.logger.LogInfo("", "", "Server Exited Properly")

}
