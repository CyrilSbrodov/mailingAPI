package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"mailingAPI/cmd/loggers"
	mock_storage "mailingAPI/internal/mocks"
	"mailingAPI/internal/storage/models"
)

func TestHandler_AddClient(t *testing.T) {
	tests := []struct {
		name         string
		body         models.Client
		answer       interface{}
		expectedCode int
	}{
		{
			name: "Test ok",
			body: models.Client{
				PhoneNumber:    "79123456789",
				MobileOperator: "test",
				Tag:            "male",
			},
			answer:       nil,
			expectedCode: http.StatusOK,
		},
		{
			name: "Test 400",
			body: models.Client{
				PhoneNumber:    "69123456789",
				MobileOperator: "test",
				Tag:            "male",
			},
			answer:       errors.New("wrong phone number format"),
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			s := mock_storage.NewMockStorage(ctrl)
			logger := loggers.NewLogger()
			h := &Handler{
				storage: s,
				logger:  logger,
			}

			bodyJSON, err := json.Marshal(tt.body)
			assert.NoError(t, err)

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/api/client", bytes.NewBuffer(bodyJSON))
			s.EXPECT().AddClient(gomock.Any()).Return(tt.answer).AnyTimes()
			h.AddClient().ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedCode, rec.Code)
			tt.answer = ""
		})
	}
}

func TestHandler_AddMailing(t *testing.T) {
	tests := []struct {
		name         string
		body         models.Mailing
		answer       interface{}
		expectedCode int
	}{
		{
			name: "Test ok",
			body: models.Mailing{
				Message:   "test",
				TimeStart: time.Now(),
				TimeEnd:   time.Now().Add(1 * time.Hour),
				Filter: models.Filter{
					MobileOperator: "testOperator",
					Tag:            "male",
				},
			},
			answer:       nil,
			expectedCode: http.StatusOK,
		},
		{
			name: "Test 500",
			body: models.Mailing{
				Message:   "test",
				TimeStart: time.Now(),
				TimeEnd:   time.Now().Add(1 * time.Hour),
				Filter: models.Filter{
					MobileOperator: "testOperator",
					Tag:            "male",
				},
			},
			answer:       errors.New("err"),
			expectedCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			s := mock_storage.NewMockStorage(ctrl)
			logger := loggers.NewLogger()
			h := &Handler{
				storage: s,
				logger:  logger,
			}

			bodyJSON, err := json.Marshal(tt.body)
			assert.NoError(t, err)

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/api/mailing", bytes.NewBuffer(bodyJSON))
			s.EXPECT().AddMailing(gomock.Any()).Return(tt.answer).AnyTimes()
			h.AddMailing().ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedCode, rec.Code)
		})
	}
}

func TestAddClientHandler(t *testing.T) {
	handler := NewHandler(nil, nil, &MockStorage{})
	router := setupRouter(handler)

	client := models.Client{
		PhoneNumber:    "79123456789",
		MobileOperator: "operator",
		Tag:            "tag",
	}

	clientJSON, err := json.Marshal(client)
	assert.NoError(t, err)

	req := httptest.NewRequest("POST", "/api/client", bytes.NewReader(clientJSON))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateClientHandler(t *testing.T) {
	handler := NewHandler(nil, nil, &MockStorage{})
	router := setupRouter(handler)

	client := models.Client{
		ID:             1,
		PhoneNumber:    "72345678901",
		MobileOperator: "operator",
		Tag:            "tag",
	}

	clientJSON, err := json.Marshal(client)
	assert.NoError(t, err)

	req := httptest.NewRequest("POST", "/api/client/update", bytes.NewReader(clientJSON))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteClientHandler(t *testing.T) {
	handler := NewHandler(nil, nil, &MockStorage{})
	router := setupRouter(handler)

	client := models.Client{
		ID: 1,
	}

	clientJSON, err := json.Marshal(client)
	assert.NoError(t, err)

	req := httptest.NewRequest("POST", "/api/client/delete", bytes.NewReader(clientJSON))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateMailingHandler(t *testing.T) {
	handler := NewHandler(nil, nil, &MockStorage{})
	router := setupRouter(handler)

	mailing := models.Mailing{
		ID:        1,
		TimeStart: time.Now(),
		TimeEnd:   time.Now().Add(1 * time.Hour),
		Filter: models.Filter{
			MobileOperator: "operator",
			Tag:            "tag",
		},
		Message: "content",
	}

	mailingJSON, err := json.Marshal(mailing)
	assert.NoError(t, err)

	req := httptest.NewRequest("POST", "/api/mailing/update", bytes.NewReader(mailingJSON))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteMailingHandler(t *testing.T) {
	handler := NewHandler(nil, nil, &MockStorage{})
	router := setupRouter(handler)

	mailing := models.Mailing{
		ID: 1,
	}

	mailingJSON, err := json.Marshal(mailing)
	assert.NoError(t, err)

	req := httptest.NewRequest("POST", "/api/mailing/delete", bytes.NewReader(mailingJSON))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetAllStatisticHandler(t *testing.T) {
	handler := NewHandler(nil, nil, &MockStorage{})
	router := setupRouter(handler)

	req := httptest.NewRequest("GET", "/api/mailing", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetDetailStatisticHandler(t *testing.T) {
	handler := NewHandler(nil, nil, &MockStorage{})
	router := setupRouter(handler)

	statistics := models.Statistics{
		MailingID: 1,
	}

	statisticsJSON, err := json.Marshal(statistics)
	assert.NoError(t, err)

	req := httptest.NewRequest("POST", "/api/mailing/get", bytes.NewReader(statisticsJSON))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

// MockStorage is a mock implementation of the storage interface for testing.
type MockStorage struct{}

func (s *MockStorage) GetAllMailingStat() (models.AllStatistics, error) {
	return models.AllStatistics{MailingID: 1}, nil
}

func (s *MockStorage) GetOneMailingStatByID(statistics *models.Statistics) ([]models.Statistics, error) {
	return []models.Statistics{{MailingID: 1}}, nil
}

func (s *MockStorage) UpdateMailing(m *models.Mailing) error {
	return nil
}

func (s *MockStorage) DeleteMailing(m *models.Mailing) error {
	return nil
}

func (s *MockStorage) ActiveProcessMailing() (map[string][]models.ActiveMailing, error) {
	//TODO implement me
	panic("implement me")
}

func (s *MockStorage) UpdateStatusMessage(mailings []models.ActiveMailing) error {
	return nil
}

func (s *MockStorage) AddClient(c *models.Client) error {
	return nil
}

func (s *MockStorage) UpdateClient(c *models.Client) error {
	return nil
}

func (s *MockStorage) DeleteClient(c *models.Client) error {
	return nil
}

func (s *MockStorage) AddMailing(m *models.Mailing) error {
	return nil
}

// setupRouter is a helper function to set up a router with the provided handler.
func setupRouter(handler *Handler) *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/api/client", http.HandlerFunc(handler.AddClient()))
	router.Handle("/api/client/update", http.HandlerFunc(handler.UpdateClient()))
	router.Handle("/api/client/delete", http.HandlerFunc(handler.DeleteClient()))
	router.Handle("/api/mailing/update", http.HandlerFunc(handler.UpdateMailing()))
	router.Handle("/api/mailing/delete", http.HandlerFunc(handler.DeleteMailing()))
	router.Handle("/api/mailing", http.HandlerFunc(handler.GetAllStatistic()))
	router.Handle("/api/mailing/get", http.HandlerFunc(handler.GetDetailStatistic()))
	return router
}
