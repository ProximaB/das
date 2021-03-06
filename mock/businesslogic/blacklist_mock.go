// Code generated by MockGen. DO NOT EDIT.
// Source: ./businesslogic/blacklist.go

// Package mock_businesslogic is a generated GoMock package.
package mock_businesslogic

import (
	businesslogic "github.com/ProximaB/das/businesslogic"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIPartnershipRequestBlacklistReasonRepository is a mock of IPartnershipRequestBlacklistReasonRepository interface
type MockIPartnershipRequestBlacklistReasonRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIPartnershipRequestBlacklistReasonRepositoryMockRecorder
}

// MockIPartnershipRequestBlacklistReasonRepositoryMockRecorder is the mock recorder for MockIPartnershipRequestBlacklistReasonRepository
type MockIPartnershipRequestBlacklistReasonRepositoryMockRecorder struct {
	mock *MockIPartnershipRequestBlacklistReasonRepository
}

// NewMockIPartnershipRequestBlacklistReasonRepository creates a new mock instance
func NewMockIPartnershipRequestBlacklistReasonRepository(ctrl *gomock.Controller) *MockIPartnershipRequestBlacklistReasonRepository {
	mock := &MockIPartnershipRequestBlacklistReasonRepository{ctrl: ctrl}
	mock.recorder = &MockIPartnershipRequestBlacklistReasonRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIPartnershipRequestBlacklistReasonRepository) EXPECT() *MockIPartnershipRequestBlacklistReasonRepositoryMockRecorder {
	return m.recorder
}

// GetPartnershipRequestBlacklistReasons mocks base method
func (m *MockIPartnershipRequestBlacklistReasonRepository) GetPartnershipRequestBlacklistReasons() ([]businesslogic.PartnershipRequestBlacklistReason, error) {
	ret := m.ctrl.Call(m, "GetPartnershipRequestBlacklistReasons")
	ret0, _ := ret[0].([]businesslogic.PartnershipRequestBlacklistReason)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartnershipRequestBlacklistReasons indicates an expected call of GetPartnershipRequestBlacklistReasons
func (mr *MockIPartnershipRequestBlacklistReasonRepositoryMockRecorder) GetPartnershipRequestBlacklistReasons() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartnershipRequestBlacklistReasons", reflect.TypeOf((*MockIPartnershipRequestBlacklistReasonRepository)(nil).GetPartnershipRequestBlacklistReasons))
}

// MockIPartnershipRequestBlacklistRepository is a mock of IPartnershipRequestBlacklistRepository interface
type MockIPartnershipRequestBlacklistRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIPartnershipRequestBlacklistRepositoryMockRecorder
}

// MockIPartnershipRequestBlacklistRepositoryMockRecorder is the mock recorder for MockIPartnershipRequestBlacklistRepository
type MockIPartnershipRequestBlacklistRepositoryMockRecorder struct {
	mock *MockIPartnershipRequestBlacklistRepository
}

// NewMockIPartnershipRequestBlacklistRepository creates a new mock instance
func NewMockIPartnershipRequestBlacklistRepository(ctrl *gomock.Controller) *MockIPartnershipRequestBlacklistRepository {
	mock := &MockIPartnershipRequestBlacklistRepository{ctrl: ctrl}
	mock.recorder = &MockIPartnershipRequestBlacklistRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIPartnershipRequestBlacklistRepository) EXPECT() *MockIPartnershipRequestBlacklistRepositoryMockRecorder {
	return m.recorder
}

// SearchPartnershipRequestBlacklist mocks base method
func (m *MockIPartnershipRequestBlacklistRepository) SearchPartnershipRequestBlacklist(criteria businesslogic.SearchPartnershipRequestBlacklistCriteria) ([]businesslogic.PartnershipRequestBlacklistEntry, error) {
	ret := m.ctrl.Call(m, "SearchPartnershipRequestBlacklist", criteria)
	ret0, _ := ret[0].([]businesslogic.PartnershipRequestBlacklistEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPartnershipRequestBlacklist indicates an expected call of SearchPartnershipRequestBlacklist
func (mr *MockIPartnershipRequestBlacklistRepositoryMockRecorder) SearchPartnershipRequestBlacklist(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPartnershipRequestBlacklist", reflect.TypeOf((*MockIPartnershipRequestBlacklistRepository)(nil).SearchPartnershipRequestBlacklist), criteria)
}

// CreatePartnershipRequestBlacklist mocks base method
func (m *MockIPartnershipRequestBlacklistRepository) CreatePartnershipRequestBlacklist(blacklist *businesslogic.PartnershipRequestBlacklistEntry) error {
	ret := m.ctrl.Call(m, "CreatePartnershipRequestBlacklist", blacklist)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePartnershipRequestBlacklist indicates an expected call of CreatePartnershipRequestBlacklist
func (mr *MockIPartnershipRequestBlacklistRepositoryMockRecorder) CreatePartnershipRequestBlacklist(blacklist interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePartnershipRequestBlacklist", reflect.TypeOf((*MockIPartnershipRequestBlacklistRepository)(nil).CreatePartnershipRequestBlacklist), blacklist)
}

// DeletePartnershipRequestBlacklist mocks base method
func (m *MockIPartnershipRequestBlacklistRepository) DeletePartnershipRequestBlacklist(blacklist businesslogic.PartnershipRequestBlacklistEntry) error {
	ret := m.ctrl.Call(m, "DeletePartnershipRequestBlacklist", blacklist)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePartnershipRequestBlacklist indicates an expected call of DeletePartnershipRequestBlacklist
func (mr *MockIPartnershipRequestBlacklistRepositoryMockRecorder) DeletePartnershipRequestBlacklist(blacklist interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePartnershipRequestBlacklist", reflect.TypeOf((*MockIPartnershipRequestBlacklistRepository)(nil).DeletePartnershipRequestBlacklist), blacklist)
}

// UpdatePartnershipRequestBlacklist mocks base method
func (m *MockIPartnershipRequestBlacklistRepository) UpdatePartnershipRequestBlacklist(blacklist businesslogic.PartnershipRequestBlacklistEntry) error {
	ret := m.ctrl.Call(m, "UpdatePartnershipRequestBlacklist", blacklist)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePartnershipRequestBlacklist indicates an expected call of UpdatePartnershipRequestBlacklist
func (mr *MockIPartnershipRequestBlacklistRepositoryMockRecorder) UpdatePartnershipRequestBlacklist(blacklist interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePartnershipRequestBlacklist", reflect.TypeOf((*MockIPartnershipRequestBlacklistRepository)(nil).UpdatePartnershipRequestBlacklist), blacklist)
}
