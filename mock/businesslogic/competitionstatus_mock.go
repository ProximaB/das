// Code generated by MockGen. DO NOT EDIT.
// Source: ./businesslogic/competitionstatus.go

// Package mock_businesslogic is a generated GoMock package.
package mock_businesslogic

import (
	businesslogic "github.com/DancesportSoftware/das/businesslogic"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockICompetitionStatusRepository is a mock of ICompetitionStatusRepository interface
type MockICompetitionStatusRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICompetitionStatusRepositoryMockRecorder
}

// MockICompetitionStatusRepositoryMockRecorder is the mock recorder for MockICompetitionStatusRepository
type MockICompetitionStatusRepositoryMockRecorder struct {
	mock *MockICompetitionStatusRepository
}

// NewMockICompetitionStatusRepository creates a new mock instance
func NewMockICompetitionStatusRepository(ctrl *gomock.Controller) *MockICompetitionStatusRepository {
	mock := &MockICompetitionStatusRepository{ctrl: ctrl}
	mock.recorder = &MockICompetitionStatusRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICompetitionStatusRepository) EXPECT() *MockICompetitionStatusRepositoryMockRecorder {
	return m.recorder
}

// GetCompetitionAllStatus mocks base method
func (m *MockICompetitionStatusRepository) GetCompetitionAllStatus() ([]businesslogic.CompetitionStatus, error) {
	ret := m.ctrl.Call(m, "GetCompetitionAllStatus")
	ret0, _ := ret[0].([]businesslogic.CompetitionStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompetitionAllStatus indicates an expected call of GetCompetitionAllStatus
func (mr *MockICompetitionStatusRepositoryMockRecorder) GetCompetitionAllStatus() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompetitionAllStatus", reflect.TypeOf((*MockICompetitionStatusRepository)(nil).GetCompetitionAllStatus))
}
