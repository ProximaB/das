// Code generated by MockGen. DO NOT EDIT.
// Source: ./businesslogic/competition_official.go

// Package mock_businesslogic is a generated GoMock package.
package mock_businesslogic

import (
	businesslogic "github.com/DancesportSoftware/das/businesslogic"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockICompetitionOfficialRepository is a mock of ICompetitionOfficialRepository interface
type MockICompetitionOfficialRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICompetitionOfficialRepositoryMockRecorder
}

// MockICompetitionOfficialRepositoryMockRecorder is the mock recorder for MockICompetitionOfficialRepository
type MockICompetitionOfficialRepositoryMockRecorder struct {
	mock *MockICompetitionOfficialRepository
}

// NewMockICompetitionOfficialRepository creates a new mock instance
func NewMockICompetitionOfficialRepository(ctrl *gomock.Controller) *MockICompetitionOfficialRepository {
	mock := &MockICompetitionOfficialRepository{ctrl: ctrl}
	mock.recorder = &MockICompetitionOfficialRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICompetitionOfficialRepository) EXPECT() *MockICompetitionOfficialRepositoryMockRecorder {
	return m.recorder
}

// CreateCompetitionOfficial mocks base method
func (m *MockICompetitionOfficialRepository) CreateCompetitionOfficial(official *businesslogic.CompetitionOfficial) error {
	ret := m.ctrl.Call(m, "CreateCompetitionOfficial", official)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCompetitionOfficial indicates an expected call of CreateCompetitionOfficial
func (mr *MockICompetitionOfficialRepositoryMockRecorder) CreateCompetitionOfficial(official interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompetitionOfficial", reflect.TypeOf((*MockICompetitionOfficialRepository)(nil).CreateCompetitionOfficial), official)
}

// DeleteCompetitionOfficial mocks base method
func (m *MockICompetitionOfficialRepository) DeleteCompetitionOfficial(official businesslogic.CompetitionOfficial) error {
	ret := m.ctrl.Call(m, "DeleteCompetitionOfficial", official)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCompetitionOfficial indicates an expected call of DeleteCompetitionOfficial
func (mr *MockICompetitionOfficialRepositoryMockRecorder) DeleteCompetitionOfficial(official interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCompetitionOfficial", reflect.TypeOf((*MockICompetitionOfficialRepository)(nil).DeleteCompetitionOfficial), official)
}

// SearchCompetitionOfficial mocks base method
func (m *MockICompetitionOfficialRepository) SearchCompetitionOfficial(criteria businesslogic.SearchCompetitionOfficialCriteria) ([]businesslogic.CompetitionOfficial, error) {
	ret := m.ctrl.Call(m, "SearchCompetitionOfficial", criteria)
	ret0, _ := ret[0].([]businesslogic.CompetitionOfficial)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchCompetitionOfficial indicates an expected call of SearchCompetitionOfficial
func (mr *MockICompetitionOfficialRepositoryMockRecorder) SearchCompetitionOfficial(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchCompetitionOfficial", reflect.TypeOf((*MockICompetitionOfficialRepository)(nil).SearchCompetitionOfficial), criteria)
}

// UpdateCompetitionOfficial mocks base method
func (m *MockICompetitionOfficialRepository) UpdateCompetitionOfficial(official businesslogic.CompetitionOfficial) error {
	ret := m.ctrl.Call(m, "UpdateCompetitionOfficial", official)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCompetitionOfficial indicates an expected call of UpdateCompetitionOfficial
func (mr *MockICompetitionOfficialRepositoryMockRecorder) UpdateCompetitionOfficial(official interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCompetitionOfficial", reflect.TypeOf((*MockICompetitionOfficialRepository)(nil).UpdateCompetitionOfficial), official)
}