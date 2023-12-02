package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"mailingAPI/internal/storage/models"
)

// AddMailing
// @Summary AddMailing
// @Tags Mailing
// @Description add a new mailing
// @ID AddMailing
// @Accept json
// @Produce json
// @Param input body models.Mailing true "add new mailing"
// @Success 200 {integer} integer 1
// @Failure 400
// @Failure 500
// @Router /api/mailing [post]
func (h *Handler) AddMailing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var m models.Mailing
		if err = json.Unmarshal(content, &m); err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// TODO проверка полей юзера
		if err = h.storage.AddMailing(&m); err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
}

// UpdateMailing
// @Summary UpdateMailing
// @Tags Mailing
// @Description update mailing params
// @ID UpdateMailing
// @Accept json
// @Produce json
// @Param input body models.Mailing true "new mailing params"
// @Success 200 {integer} integer 1
// @Failure 400
// @Failure 500
// @Router /api/mailing/update [post]
func (h *Handler) UpdateMailing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var m models.Mailing
		if err = json.Unmarshal(content, &m); err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// TODO проверка полей юзера
		if err = h.storage.UpdateMailing(&m); err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
}

// DeleteMailing
// @Summary DeleteMailing
// @Tags Mailing
// @Description delete mailing by id
// @ID DeleteMailing
// @Accept json
// @Produce json
// @Param input body models.Mailing true "delete mailing"
// @Success 200 {integer} integer 1
// @Failure 400
// @Failure 500
// @Router /api/mailing/delete [post]
func (h *Handler) DeleteMailing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var m models.Mailing
		if err = json.Unmarshal(content, &m); err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// TODO проверка полей юзера
		if err = h.storage.DeleteMailing(&m); err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
}

// GetAllStatistic
// @Summary GetAllStatistic
// @Tags Mailing
// @Description get all statistics
// @ID GetAllStatistic
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400
// @Failure 500
// @Router /api/mailing [get]
func (h *Handler) GetAllStatistic() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO проверка полей юзера
		m, err := h.storage.GetAllMailingStat()
		if err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if m.MailingID == 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("there no one statistic success"))
			return
		}
		jsonMail, err := json.Marshal(m)
		if err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonMail)
		return
	}
}

// GetDetailStatistic
// @Summary GetDetailStatistic
// @Tags Mailing
// @Description get detail statistics by mailing id
// @ID GetDetailStatistic
// @Accept json
// @Produce json
// @Param input body models.Statistics true "detail statistics"
// @Success 200 {integer} integer 1
// @Failure 400
// @Failure 500
// @Router /api/mailing/get [post]
func (h *Handler) GetDetailStatistic() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var m models.Statistics
		if err = json.Unmarshal(content, &m); err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		mail, err := h.storage.GetOneMailingStatByID(&m)
		if err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		jsonMail, err := json.Marshal(mail)
		if err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonMail)
		return
	}
}
