// Code generated by MockGen. DO NOT EDIT.
// Source: user_config_repo.go

// Package mock_users is a generated GoMock package.
package users

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockUserConfigRepo is a mock of UserConfigRepo interface.
type MockUserConfigRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserConfigRepoMockRecorder
}

// MockUserConfigRepoMockRecorder is the mock recorder for MockUserConfigRepo.
type MockUserConfigRepoMockRecorder struct {
	mock *MockUserConfigRepo
}

// NewMockUserConfigRepo creates a new mock instance.
func NewMockUserConfigRepo(ctrl *gomock.Controller) *MockUserConfigRepo {
	mock := &MockUserConfigRepo{ctrl: ctrl}
	mock.recorder = &MockUserConfigRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserConfigRepo) EXPECT() *MockUserConfigRepoMockRecorder {
	return m.recorder
}

// GetByUserID mocks base method.
func (m *MockUserConfigRepo) GetByUserID(ctx context.Context, userID uuid.UUID) (*UserConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserID", ctx, userID)
	ret0, _ := ret[0].(*UserConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserID indicates an expected call of GetByUserID.
func (mr *MockUserConfigRepoMockRecorder) GetByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserID", reflect.TypeOf((*MockUserConfigRepo)(nil).GetByUserID), ctx, userID)
}
