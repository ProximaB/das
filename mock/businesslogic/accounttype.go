// Code generated by MockGen. DO NOT EDIT.
// Source: ./businesslogic/accounttype.go

// Package mock_businesslogic is a generated GoMock package.
package mock_businesslogic

import (
	businesslogic "github.com/DancesportSoftware/das/businesslogic"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIAccountTypeRepository is a mock of IAccountTypeRepository interface
type MockIAccountTypeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAccountTypeRepositoryMockRecorder
}

// MockIAccountTypeRepositoryMockRecorder is the mock recorder for MockIAccountTypeRepository
type MockIAccountTypeRepositoryMockRecorder struct {
	mock *MockIAccountTypeRepository
}

// NewMockIAccountTypeRepository creates a new mock instance
func NewMockIAccountTypeRepository(ctrl *gomock.Controller) *MockIAccountTypeRepository {
	mock := &MockIAccountTypeRepository{ctrl: ctrl}
	mock.recorder = &MockIAccountTypeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIAccountTypeRepository) EXPECT() *MockIAccountTypeRepositoryMockRecorder {
	return m.recorder
}

// GetAccountTypes mocks base method
func (m *MockIAccountTypeRepository) GetAccountTypes() ([]businesslogic.AccountType, error) {
	ret := m.ctrl.Call(m, "GetAccountTypes")
	ret0, _ := ret[0].([]businesslogic.AccountType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountTypes indicates an expected call of GetAccountTypes
func (mr *MockIAccountTypeRepositoryMockRecorder) GetAccountTypes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountTypes", reflect.TypeOf((*MockIAccountTypeRepository)(nil).GetAccountTypes))
}
