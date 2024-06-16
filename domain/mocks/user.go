// Code generated by MockGen. DO NOT EDIT.
// Source: user.go
//
// Generated by this command:
//
//	mockgen -source=user.go -destination=mocks/user.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	domain "business-logic/domain"
	context "context"
	reflect "reflect"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockIUserRepository is a mock of IUserRepository interface.
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository.
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance.
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIUserRepository) Create(ctx context.Context, user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIUserRepositoryMockRecorder) Create(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUserRepository)(nil).Create), ctx, user)
}

// DeleteById mocks base method.
func (m *MockIUserRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockIUserRepositoryMockRecorder) DeleteById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockIUserRepository)(nil).DeleteById), ctx, id)
}

// GetAll mocks base method.
func (m *MockIUserRepository) GetAll(ctx context.Context, page int) ([]*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, page)
	ret0, _ := ret[0].([]*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIUserRepositoryMockRecorder) GetAll(ctx, page any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIUserRepository)(nil).GetAll), ctx, page)
}

// GetById mocks base method.
func (m *MockIUserRepository) GetById(ctx context.Context, userId uuid.UUID) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, userId)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockIUserRepositoryMockRecorder) GetById(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockIUserRepository)(nil).GetById), ctx, userId)
}

// GetByUsername mocks base method.
func (m *MockIUserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsername", ctx, username)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsername indicates an expected call of GetByUsername.
func (mr *MockIUserRepositoryMockRecorder) GetByUsername(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsername", reflect.TypeOf((*MockIUserRepository)(nil).GetByUsername), ctx, username)
}

// Update mocks base method.
func (m *MockIUserRepository) Update(ctx context.Context, user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIUserRepositoryMockRecorder) Update(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIUserRepository)(nil).Update), ctx, user)
}

// MockIUserService is a mock of IUserService interface.
type MockIUserService struct {
	ctrl     *gomock.Controller
	recorder *MockIUserServiceMockRecorder
}

// MockIUserServiceMockRecorder is the mock recorder for MockIUserService.
type MockIUserServiceMockRecorder struct {
	mock *MockIUserService
}

// NewMockIUserService creates a new mock instance.
func NewMockIUserService(ctrl *gomock.Controller) *MockIUserService {
	mock := &MockIUserService{ctrl: ctrl}
	mock.recorder = &MockIUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserService) EXPECT() *MockIUserServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIUserService) Create(user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIUserServiceMockRecorder) Create(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUserService)(nil).Create), user)
}

// DeleteById mocks base method.
func (m *MockIUserService) DeleteById(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockIUserServiceMockRecorder) DeleteById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockIUserService)(nil).DeleteById), id)
}

// GetAll mocks base method.
func (m *MockIUserService) GetAll(page int) ([]*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", page)
	ret0, _ := ret[0].([]*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIUserServiceMockRecorder) GetAll(page any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIUserService)(nil).GetAll), page)
}

// GetById mocks base method.
func (m *MockIUserService) GetById(userId uuid.UUID) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", userId)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockIUserServiceMockRecorder) GetById(userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockIUserService)(nil).GetById), userId)
}

// GetByUsername mocks base method.
func (m *MockIUserService) GetByUsername(username string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsername", username)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsername indicates an expected call of GetByUsername.
func (mr *MockIUserServiceMockRecorder) GetByUsername(username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsername", reflect.TypeOf((*MockIUserService)(nil).GetByUsername), username)
}

// Update mocks base method.
func (m *MockIUserService) Update(user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIUserServiceMockRecorder) Update(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIUserService)(nil).Update), user)
}
