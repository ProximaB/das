// Code generated by MockGen. DO NOT EDIT.
// Source: ./businesslogic/competition.go

// Package mock_businesslogic is a generated GoMock package.
package mock_businesslogic

import (
	businesslogic "github.com/DancesportSoftware/das/businesslogic"
	reference "github.com/DancesportSoftware/das/businesslogic/reference"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockICompetitionRepository is a mock of ICompetitionRepository interface
type MockICompetitionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICompetitionRepositoryMockRecorder
}

// MockICompetitionRepositoryMockRecorder is the mock recorder for MockICompetitionRepository
type MockICompetitionRepositoryMockRecorder struct {
	mock *MockICompetitionRepository
}

// NewMockICompetitionRepository creates a new mock instance
func NewMockICompetitionRepository(ctrl *gomock.Controller) *MockICompetitionRepository {
	mock := &MockICompetitionRepository{ctrl: ctrl}
	mock.recorder = &MockICompetitionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICompetitionRepository) EXPECT() *MockICompetitionRepositoryMockRecorder {
	return m.recorder
}

// CreateCompetition mocks base method
func (m *MockICompetitionRepository) CreateCompetition(competition *businesslogic.Competition) error {
	ret := m.ctrl.Call(m, "CreateCompetition", competition)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCompetition indicates an expected call of CreateCompetition
func (mr *MockICompetitionRepositoryMockRecorder) CreateCompetition(competition interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompetition", reflect.TypeOf((*MockICompetitionRepository)(nil).CreateCompetition), competition)
}

// SearchCompetition mocks base method
func (m *MockICompetitionRepository) SearchCompetition(criteria businesslogic.SearchCompetitionCriteria) ([]businesslogic.Competition, error) {
	ret := m.ctrl.Call(m, "SearchCompetition", criteria)
	ret0, _ := ret[0].([]businesslogic.Competition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchCompetition indicates an expected call of SearchCompetition
func (mr *MockICompetitionRepositoryMockRecorder) SearchCompetition(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchCompetition", reflect.TypeOf((*MockICompetitionRepository)(nil).SearchCompetition), criteria)
}

// UpdateCompetition mocks base method
func (m *MockICompetitionRepository) UpdateCompetition(competition businesslogic.Competition) error {
	ret := m.ctrl.Call(m, "UpdateCompetition", competition)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCompetition indicates an expected call of UpdateCompetition
func (mr *MockICompetitionRepositoryMockRecorder) UpdateCompetition(competition interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCompetition", reflect.TypeOf((*MockICompetitionRepository)(nil).UpdateCompetition), competition)
}

// DeleteCompetition mocks base method
func (m *MockICompetitionRepository) DeleteCompetition(competition businesslogic.Competition) error {
	ret := m.ctrl.Call(m, "DeleteCompetition", competition)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCompetition indicates an expected call of DeleteCompetition
func (mr *MockICompetitionRepositoryMockRecorder) DeleteCompetition(competition interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCompetition", reflect.TypeOf((*MockICompetitionRepository)(nil).DeleteCompetition), competition)
}

// MockIEventMetaRepository is a mock of IEventMetaRepository interface
type MockIEventMetaRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIEventMetaRepositoryMockRecorder
}

// MockIEventMetaRepositoryMockRecorder is the mock recorder for MockIEventMetaRepository
type MockIEventMetaRepositoryMockRecorder struct {
	mock *MockIEventMetaRepository
}

// NewMockIEventMetaRepository creates a new mock instance
func NewMockIEventMetaRepository(ctrl *gomock.Controller) *MockIEventMetaRepository {
	mock := &MockIEventMetaRepository{ctrl: ctrl}
	mock.recorder = &MockIEventMetaRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIEventMetaRepository) EXPECT() *MockIEventMetaRepositoryMockRecorder {
	return m.recorder
}

// GetEventUniqueFederations mocks base method
func (m *MockIEventMetaRepository) GetEventUniqueFederations(competition businesslogic.Competition) ([]reference.Federation, error) {
	ret := m.ctrl.Call(m, "GetEventUniqueFederations", competition)
	ret0, _ := ret[0].([]reference.Federation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventUniqueFederations indicates an expected call of GetEventUniqueFederations
func (mr *MockIEventMetaRepositoryMockRecorder) GetEventUniqueFederations(competition interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventUniqueFederations", reflect.TypeOf((*MockIEventMetaRepository)(nil).GetEventUniqueFederations), competition)
}

// GetEventUniqueDivisions mocks base method
func (m *MockIEventMetaRepository) GetEventUniqueDivisions(competition businesslogic.Competition) ([]reference.Division, error) {
	ret := m.ctrl.Call(m, "GetEventUniqueDivisions", competition)
	ret0, _ := ret[0].([]reference.Division)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventUniqueDivisions indicates an expected call of GetEventUniqueDivisions
func (mr *MockIEventMetaRepositoryMockRecorder) GetEventUniqueDivisions(competition interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventUniqueDivisions", reflect.TypeOf((*MockIEventMetaRepository)(nil).GetEventUniqueDivisions), competition)
}

// GetEventUniqueAges mocks base method
func (m *MockIEventMetaRepository) GetEventUniqueAges(competition businesslogic.Competition) ([]reference.Age, error) {
	ret := m.ctrl.Call(m, "GetEventUniqueAges", competition)
	ret0, _ := ret[0].([]reference.Age)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventUniqueAges indicates an expected call of GetEventUniqueAges
func (mr *MockIEventMetaRepositoryMockRecorder) GetEventUniqueAges(competition interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventUniqueAges", reflect.TypeOf((*MockIEventMetaRepository)(nil).GetEventUniqueAges), competition)
}

// GetEventUniqueProficiencies mocks base method
func (m *MockIEventMetaRepository) GetEventUniqueProficiencies(competition businesslogic.Competition) ([]reference.Proficiency, error) {
	ret := m.ctrl.Call(m, "GetEventUniqueProficiencies", competition)
	ret0, _ := ret[0].([]reference.Proficiency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventUniqueProficiencies indicates an expected call of GetEventUniqueProficiencies
func (mr *MockIEventMetaRepositoryMockRecorder) GetEventUniqueProficiencies(competition interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventUniqueProficiencies", reflect.TypeOf((*MockIEventMetaRepository)(nil).GetEventUniqueProficiencies), competition)
}

// GetEventUniqueStyles mocks base method
func (m *MockIEventMetaRepository) GetEventUniqueStyles(competition businesslogic.Competition) ([]reference.Style, error) {
	ret := m.ctrl.Call(m, "GetEventUniqueStyles", competition)
	ret0, _ := ret[0].([]reference.Style)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventUniqueStyles indicates an expected call of GetEventUniqueStyles
func (mr *MockIEventMetaRepositoryMockRecorder) GetEventUniqueStyles(competition interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventUniqueStyles", reflect.TypeOf((*MockIEventMetaRepository)(nil).GetEventUniqueStyles), competition)
}
