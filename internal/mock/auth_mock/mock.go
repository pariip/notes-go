// Code generated by MockGen. DO NOT EDIT.
// Source: ./auth.go

// Package auth_mock is a generated GoMock package.
package auth_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/pariip/notes-go/internal/models"
	params "github.com/pariip/notes-go/internal/params"
)

// MockAuthService is a mock of AuthService interface.
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService.
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance.
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// GenerateAccessToken mocks base method.
func (m *MockAuthService) GenerateAccessToken(user *models.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockAuthServiceMockRecorder) GenerateAccessToken(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockAuthService)(nil).GenerateAccessToken), user)
}

// GenerateRefreshToken mocks base method.
func (m *MockAuthService) GenerateRefreshToken(user *models.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefreshToken", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRefreshToken indicates an expected call of GenerateRefreshToken.
func (mr *MockAuthServiceMockRecorder) GenerateRefreshToken(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefreshToken", reflect.TypeOf((*MockAuthService)(nil).GenerateRefreshToken), user)
}

// Login mocks base method.
func (m *MockAuthService) Login(req *params.LoginRequest) (*params.UserTokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", req)
	ret0, _ := ret[0].(*params.UserTokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthServiceMockRecorder) Login(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthService)(nil).Login), req)
}

// RefreshToken mocks base method.
func (m *MockAuthService) RefreshToken(refreshToken string, userID uint) (*params.UserTokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshToken", refreshToken, userID)
	ret0, _ := ret[0].(*params.UserTokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshToken indicates an expected call of RefreshToken.
func (mr *MockAuthServiceMockRecorder) RefreshToken(refreshToken, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshToken", reflect.TypeOf((*MockAuthService)(nil).RefreshToken), refreshToken, userID)
}

// Signup mocks base method.
func (m *MockAuthService) Signup(req *params.SignupRequest) (*params.UserTokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signup", req)
	ret0, _ := ret[0].(*params.UserTokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Signup indicates an expected call of Signup.
func (mr *MockAuthServiceMockRecorder) Signup(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signup", reflect.TypeOf((*MockAuthService)(nil).Signup), req)
}