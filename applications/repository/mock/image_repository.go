// Code generated by MockGen. DO NOT EDIT.
// Source: ./image_repository.go

// Package repositorymock is a generated GoMock package.
package repositorymock

import (
	context "context"
	domain "homeapi/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockImageRepository is a mock of ImageRepository interface.
type MockImageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockImageRepositoryMockRecorder
}

// MockImageRepositoryMockRecorder is the mock recorder for MockImageRepository.
type MockImageRepositoryMockRecorder struct {
	mock *MockImageRepository
}

// NewMockImageRepository creates a new mock instance.
func NewMockImageRepository(ctrl *gomock.Controller) *MockImageRepository {
	mock := &MockImageRepository{ctrl: ctrl}
	mock.recorder = &MockImageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageRepository) EXPECT() *MockImageRepositoryMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockImageRepository) Insert(ctx context.Context, image *domain.Image) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, image)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockImageRepositoryMockRecorder) Insert(ctx, image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockImageRepository)(nil).Insert), ctx, image)
}
