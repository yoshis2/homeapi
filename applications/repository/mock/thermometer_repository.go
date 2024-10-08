// Code generated by MockGen. DO NOT EDIT.
// Source: thermometer_repository.go
//
// Generated by this command:
//
//	mockgen -package mock -source thermometer_repository.go -destination mock/thermometer_repository.go
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	domain "homeapi/domain"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockThermometerRepository is a mock of ThermometerRepository interface.
type MockThermometerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockThermometerRepositoryMockRecorder
}

// MockThermometerRepositoryMockRecorder is the mock recorder for MockThermometerRepository.
type MockThermometerRepositoryMockRecorder struct {
	mock *MockThermometerRepository
}

// NewMockThermometerRepository creates a new mock instance.
func NewMockThermometerRepository(ctrl *gomock.Controller) *MockThermometerRepository {
	mock := &MockThermometerRepository{ctrl: ctrl}
	mock.recorder = &MockThermometerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockThermometerRepository) EXPECT() *MockThermometerRepositoryMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockThermometerRepository) Insert(ctx context.Context, thermometer *domain.Thermometer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, thermometer)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockThermometerRepositoryMockRecorder) Insert(ctx, thermometer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockThermometerRepository)(nil).Insert), ctx, thermometer)
}

// List mocks base method.
func (m *MockThermometerRepository) List(ctx context.Context) ([]domain.Thermometer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]domain.Thermometer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockThermometerRepositoryMockRecorder) List(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockThermometerRepository)(nil).List), ctx)
}
