// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	bytes "bytes"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	core "github.com/sgsoul/internal/core"
)

// MockClientXKCD is a mock of ClientXKCD interface.
type MockClientXKCD struct {
	ctrl     *gomock.Controller
	recorder *MockClientXKCDMockRecorder
}

// MockClientXKCDMockRecorder is the mock recorder for MockClientXKCD.
type MockClientXKCDMockRecorder struct {
	mock *MockClientXKCD
}

// NewMockClientXKCD creates a new mock instance.
func NewMockClientXKCD(ctrl *gomock.Controller) *MockClientXKCD {
	mock := &MockClientXKCD{ctrl: ctrl}
	mock.recorder = &MockClientXKCDMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientXKCD) EXPECT() *MockClientXKCDMockRecorder {
	return m.recorder
}

// RunWorkers mocks base method.
func (m *MockClientXKCD) RunWorkers(workers int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RunWorkers", workers)
}

// RunWorkers indicates an expected call of RunWorkers.
func (mr *MockClientXKCDMockRecorder) RunWorkers(workers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunWorkers", reflect.TypeOf((*MockClientXKCD)(nil).RunWorkers), workers)
}

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

// CreateUser mocks base method.
func (m *MockStorage) CreateUser(username, password, role string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", username, password, role)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStorageMockRecorder) CreateUser(username, password, role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStorage)(nil).CreateUser), username, password, role)
}

// GetAllComics mocks base method.
func (m *MockStorage) GetAllComics() ([]core.Comic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllComics")
	ret0, _ := ret[0].([]core.Comic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllComics indicates an expected call of GetAllComics.
func (mr *MockStorageMockRecorder) GetAllComics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllComics", reflect.TypeOf((*MockStorage)(nil).GetAllComics))
}

// GetComicByID mocks base method.
func (m *MockStorage) GetComicByID(id int) (core.Comic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComicByID", id)
	ret0, _ := ret[0].(core.Comic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComicByID indicates an expected call of GetComicByID.
func (mr *MockStorageMockRecorder) GetComicByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComicByID", reflect.TypeOf((*MockStorage)(nil).GetComicByID), id)
}

// GetCount mocks base method.
func (m *MockStorage) GetCount() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount.
func (mr *MockStorageMockRecorder) GetCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockStorage)(nil).GetCount))
}

// GetUserByUsername mocks base method.
func (m *MockStorage) GetUserByUsername(username string) (core.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", username)
	ret0, _ := ret[0].(core.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockStorageMockRecorder) GetUserByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockStorage)(nil).GetUserByUsername), username)
}

// PrettyPrint mocks base method.
func (m *MockStorage) PrettyPrint(v []core.Comic) bytes.Buffer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrettyPrint", v)
	ret0, _ := ret[0].(bytes.Buffer)
	return ret0
}

// PrettyPrint indicates an expected call of PrettyPrint.
func (mr *MockStorageMockRecorder) PrettyPrint(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrettyPrint", reflect.TypeOf((*MockStorage)(nil).PrettyPrint), v)
}

// SaveComicToDatabase mocks base method.
func (m *MockStorage) SaveComicToDatabase(comic core.Comic) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveComicToDatabase", comic)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveComicToDatabase indicates an expected call of SaveComicToDatabase.
func (mr *MockStorageMockRecorder) SaveComicToDatabase(comic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveComicToDatabase", reflect.TypeOf((*MockStorage)(nil).SaveComicToDatabase), comic)
}