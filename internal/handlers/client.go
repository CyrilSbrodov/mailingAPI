package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"mailingAPI/internal/storage/models"
)

// AddClient
// @Summary AddClient
// @Tags Client
// @Description add a new client
// @ID AddClient
// @Accept json
// @Produce json
// @Param input body models.Client true "client info"
// @Success 200 {integer} integer 1
// @Failure 400
// @Failure 500
// @Router /api/client [post]
func (h *Handler) AddClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var client models.Client
		if err = json.Unmarshal(content, &client); err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !client.CheckNumber() {
			h.logger.LogErr(fmt.Errorf("wrong format phone number"), "")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("wrong format phone number"))
			return
		}
		if err = h.storage.AddClient(&client); err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
}

// UpdateClient
// @Summary UpdateClient
// @Tags Client
// @Description update client
// @ID UpdateClient
// @Accept json
// @Produce json
// @Param input body models.Client true "update client info"
// @Success 200 {integer} integer 1
// @Failure 400
// @Failure 500
// @Router /api/client/update [post]
func (h *Handler) UpdateClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var client models.Client
		if err = json.Unmarshal(content, &client); err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !client.CheckNumber() {
			h.logger.LogErr(fmt.Errorf("wrong format phone number"), "")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("wrong format phone number"))
			return
		}
		if err = h.storage.UpdateClient(&client); err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
}

// DeleteClient
// @Summary DeleteClient
// @Tags Client
// @Description delete client
// @ID DeleteClient
// @Accept json
// @Produce json
// @Param input body models.Client true "delete client by id"
// @Success 200 {integer} integer 1
// @Failure 404
// @Router /api/client/delete [post]
func (h *Handler) DeleteClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var client models.Client
		if err = json.Unmarshal(content, &client); err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// TODO проверка полей юзера
		if err = h.storage.DeleteClient(&client); err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
}
