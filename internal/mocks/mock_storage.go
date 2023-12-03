// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	reflect "reflect"

	models "mailingAPI/internal/storage/models"

	gomock "github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// ActiveProcessMailing mocks base method.
func (m *MockStorage) ActiveProcessMailing() (map[string][]models.ActiveMailing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveProcessMailing")
	ret0, _ := ret[0].(map[string][]models.ActiveMailing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActiveProcessMailing indicates an expected call of ActiveProcessMailing.
func (mr *MockStorageMockRecorder) ActiveProcessMailing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveProcessMailing", reflect.TypeOf((*MockStorage)(nil).ActiveProcessMailing))
}

// AddClient mocks base method.
func (m *MockStorage) AddClient(c *models.Client) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddClient", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddClient indicates an expected call of AddClient.
func (mr *MockStorageMockRecorder) AddClient(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddClient", reflect.TypeOf((*MockStorage)(nil).AddClient), c)
}

// AddMailing mocks base method.
func (m_2 *MockStorage) AddMailing(m *models.Mailing) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "AddMailing", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMailing indicates an expected call of AddMailing.
func (mr *MockStorageMockRecorder) AddMailing(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMailing", reflect.TypeOf((*MockStorage)(nil).AddMailing), m)
}

// DeleteClient mocks base method.
func (m *MockStorage) DeleteClient(c *models.Client) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteClient", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteClient indicates an expected call of DeleteClient.
func (mr *MockStorageMockRecorder) DeleteClient(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClient", reflect.TypeOf((*MockStorage)(nil).DeleteClient), c)
}

// DeleteMailing mocks base method.
func (m_2 *MockStorage) DeleteMailing(m *models.Mailing) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "DeleteMailing", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMailing indicates an expected call of DeleteMailing.
func (mr *MockStorageMockRecorder) DeleteMailing(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMailing", reflect.TypeOf((*MockStorage)(nil).DeleteMailing), m)
}

// GetAllMailingStat mocks base method.
func (m *MockStorage) GetAllMailingStat() (models.AllStatistics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllMailingStat")
	ret0, _ := ret[0].(models.AllStatistics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllMailingStat indicates an expected call of GetAllMailingStat.
func (mr *MockStorageMockRecorder) GetAllMailingStat() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllMailingStat", reflect.TypeOf((*MockStorage)(nil).GetAllMailingStat))
}

// GetOneMailingStatByID mocks base method.
func (m_2 *MockStorage) GetOneMailingStatByID(m *models.Statistics) ([]models.Statistics, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "GetOneMailingStatByID", m)
	ret0, _ := ret[0].([]models.Statistics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneMailingStatByID indicates an expected call of GetOneMailingStatByID.
func (mr *MockStorageMockRecorder) GetOneMailingStatByID(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneMailingStatByID", reflect.TypeOf((*MockStorage)(nil).GetOneMailingStatByID), m)
}

// UpdateClient mocks base method.
func (m *MockStorage) UpdateClient(c *models.Client) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClient", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClient indicates an expected call of UpdateClient.
func (mr *MockStorageMockRecorder) UpdateClient(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClient", reflect.TypeOf((*MockStorage)(nil).UpdateClient), c)
}

// UpdateMailing mocks base method.
func (m_2 *MockStorage) UpdateMailing(m *models.Mailing) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "UpdateMailing", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMailing indicates an expected call of UpdateMailing.
func (mr *MockStorageMockRecorder) UpdateMailing(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMailing", reflect.TypeOf((*MockStorage)(nil).UpdateMailing), m)
}

// UpdateStatusMessage mocks base method.
func (m *MockStorage) UpdateStatusMessage(arg0 []models.ActiveMailing) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatusMessage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatusMessage indicates an expected call of UpdateStatusMessage.
func (mr *MockStorageMockRecorder) UpdateStatusMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatusMessage", reflect.TypeOf((*MockStorage)(nil).UpdateStatusMessage), arg0)
}
