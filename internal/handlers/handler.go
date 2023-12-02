package handlers

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "mailingAPI/docs"

	"mailingAPI/cmd/config"
	"mailingAPI/cmd/loggers"
	"mailingAPI/internal/storage"
)

type Handlers interface {
	Register(router *chi.Mux)
}

type Handler struct {
	cfg     *config.ServerConfig
	logger  *loggers.Logger
	storage storage.Storage
}

func NewHandler(cfg *config.ServerConfig, logger *loggers.Logger, storage storage.Storage) *Handler {
	return &Handler{
		cfg:     cfg,
		logger:  logger,
		storage: storage,
	}
}

func (h *Handler) Register(r *chi.Mux) {
	//r.Group(func(r chi.Router) {
	//	r.Post("/api/register", h.Registration())
	//	r.Post("/api/login", h.Login())
	//})
	r.Group(func(r chi.Router) {
		r.Post("/api/client", h.AddClient())
		r.Post("/api/client/update", h.UpdateClient())
		r.Post("/api/client/delete", h.DeleteClient())
		r.Post("/api/mailing", h.AddMailing())
		r.Post("/api/mailing/update", h.UpdateMailing())
		r.Post("/api/mailing/delete", h.DeleteMailing())
		r.Get("/api/mailing", h.GetAllStatistic())
		r.Post("/api/mailing/get", h.GetDetailStatistic())
		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		))
	})
}
