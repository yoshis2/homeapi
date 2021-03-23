// Code generated by MockGen. DO NOT EDIT.
// Source: ./temperature_repository.go

// Package repositorymock is a generated GoMock package.
package repositorymock

import (
	context "context"
	domain "homeapi/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTemperatureRepository is a mock of TemperatureRepository interface.
type MockTemperatureRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTemperatureRepositoryMockRecorder
}

// MockTemperatureRepositoryMockRecorder is the mock recorder for MockTemperatureRepository.
type MockTemperatureRepositoryMockRecorder struct {
	mock *MockTemperatureRepository
}

// NewMockTemperatureRepository creates a new mock instance.
func NewMockTemperatureRepository(ctrl *gomock.Controller) *MockTemperatureRepository {
	mock := &MockTemperatureRepository{ctrl: ctrl}
	mock.recorder = &MockTemperatureRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTemperatureRepository) EXPECT() *MockTemperatureRepositoryMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockTemperatureRepository) Insert(arg0 *domain.Temperature) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockTemperatureRepositoryMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTemperatureRepository)(nil).Insert), arg0)
}

// List mocks base method.
func (m *MockTemperatureRepository) List(ctx context.Context) ([]domain.Temperature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]domain.Temperature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockTemperatureRepositoryMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTemperatureRepository)(nil).List), ctx)
}
