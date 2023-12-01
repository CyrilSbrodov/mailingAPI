package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"mailingAPI/internal/storage/models"
)

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

func (h *Handler) GetAllStatistic() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO проверка полей юзера
		m, err := h.storage.GetAllMailingStat()
		if err != nil {
			h.logger.LogErr(err, "")
			w.WriteHeader(http.StatusInternalServerError)
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

func (h *Handler) GetDetailStatistic() http.HandlerFunc {
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
		mail, err := h.storage.GetOneMailingStat(&m)
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

func (h *Handler) ActivateProcessMailing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO
		//content, err := io.ReadAll(r.Body)
		//if err != nil {
		//	h.logger.LogErr(err, "")
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		//var m models.Mailing
		//if err = json.Unmarshal(content, &m); err != nil {
		//	h.logger.LogErr(err, "")
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		//// TODO проверка полей юзера
		//mail, err := h.storage.GetOneMailingStat(&m)
		//if err != nil {
		//	h.logger.LogErr(err, "")
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		//jsonMail, err := json.Marshal(mail)
		//if err != nil {
		//	h.logger.LogErr(err, "")
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		//w.Header().Set("Content-Type", "application/json")
		//w.WriteHeader(http.StatusOK)
		//w.Write(jsonMail)
		return
	}
}
