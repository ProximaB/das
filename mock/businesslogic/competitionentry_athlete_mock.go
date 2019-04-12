// Code generated by MockGen. DO NOT EDIT.
// Source: ./businesslogic/competitionentry_athlete.go

// Package mock_businesslogic is a generated GoMock package.
package mock_businesslogic

import (
	businesslogic "github.com/DancesportSoftware/das/businesslogic"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIAthleteCompetitionEntryRepository is a mock of IAthleteCompetitionEntryRepository interface
type MockIAthleteCompetitionEntryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAthleteCompetitionEntryRepositoryMockRecorder
}

// MockIAthleteCompetitionEntryRepositoryMockRecorder is the mock recorder for MockIAthleteCompetitionEntryRepository
type MockIAthleteCompetitionEntryRepositoryMockRecorder struct {
	mock *MockIAthleteCompetitionEntryRepository
}

// NewMockIAthleteCompetitionEntryRepository creates a new mock instance
func NewMockIAthleteCompetitionEntryRepository(ctrl *gomock.Controller) *MockIAthleteCompetitionEntryRepository {
	mock := &MockIAthleteCompetitionEntryRepository{ctrl: ctrl}
	mock.recorder = &MockIAthleteCompetitionEntryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIAthleteCompetitionEntryRepository) EXPECT() *MockIAthleteCompetitionEntryRepositoryMockRecorder {
	return m.recorder
}

// CreateEntry mocks base method
func (m *MockIAthleteCompetitionEntryRepository) CreateEntry(entry *businesslogic.AthleteCompetitionEntry) error {
	ret := m.ctrl.Call(m, "CreateEntry", entry)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEntry indicates an expected call of CreateEntry
func (mr *MockIAthleteCompetitionEntryRepositoryMockRecorder) CreateEntry(entry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntry", reflect.TypeOf((*MockIAthleteCompetitionEntryRepository)(nil).CreateEntry), entry)
}

// DeleteEntry mocks base method
func (m *MockIAthleteCompetitionEntryRepository) DeleteEntry(entry businesslogic.AthleteCompetitionEntry) error {
	ret := m.ctrl.Call(m, "DeleteEntry", entry)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEntry indicates an expected call of DeleteEntry
func (mr *MockIAthleteCompetitionEntryRepositoryMockRecorder) DeleteEntry(entry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEntry", reflect.TypeOf((*MockIAthleteCompetitionEntryRepository)(nil).DeleteEntry), entry)
}

// SearchEntry mocks base method
func (m *MockIAthleteCompetitionEntryRepository) SearchEntry(criteria businesslogic.SearchAthleteCompetitionEntryCriteria) ([]businesslogic.AthleteCompetitionEntry, error) {
	ret := m.ctrl.Call(m, "SearchEntry", criteria)
	ret0, _ := ret[0].([]businesslogic.AthleteCompetitionEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchEntry indicates an expected call of SearchEntry
func (mr *MockIAthleteCompetitionEntryRepositoryMockRecorder) SearchEntry(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchEntry", reflect.TypeOf((*MockIAthleteCompetitionEntryRepository)(nil).SearchEntry), criteria)
}

// UpdateEntry mocks base method
func (m *MockIAthleteCompetitionEntryRepository) UpdateEntry(entry businesslogic.AthleteCompetitionEntry) error {
	ret := m.ctrl.Call(m, "UpdateEntry", entry)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEntry indicates an expected call of UpdateEntry
func (mr *MockIAthleteCompetitionEntryRepositoryMockRecorder) UpdateEntry(entry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEntry", reflect.TypeOf((*MockIAthleteCompetitionEntryRepository)(nil).UpdateEntry), entry)
}

// NextAvailableLeadTag mocks base method
func (m *MockIAthleteCompetitionEntryRepository) NextAvailableLeadTag(competition businesslogic.Competition) (int, error) {
	ret := m.ctrl.Call(m, "NextAvailableLeadTag", competition)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextAvailableLeadTag indicates an expected call of NextAvailableLeadTag
func (mr *MockIAthleteCompetitionEntryRepositoryMockRecorder) NextAvailableLeadTag(competition interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextAvailableLeadTag", reflect.TypeOf((*MockIAthleteCompetitionEntryRepository)(nil).NextAvailableLeadTag), competition)
}

// GetEntriesByCompetition mocks base method
func (m *MockIAthleteCompetitionEntryRepository) GetEntriesByCompetition(competitionId int) ([]businesslogic.AthleteCompetitionEntry, error) {
	ret := m.ctrl.Call(m, "GetEntriesByCompetition", competitionId)
	ret0, _ := ret[0].([]businesslogic.AthleteCompetitionEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntriesByCompetition indicates an expected call of GetEntriesByCompetition
func (mr *MockIAthleteCompetitionEntryRepositoryMockRecorder) GetEntriesByCompetition(competitionId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntriesByCompetition", reflect.TypeOf((*MockIAthleteCompetitionEntryRepository)(nil).GetEntriesByCompetition), competitionId)
}

// MockICompetitionLeadTagRepository is a mock of ICompetitionLeadTagRepository interface
type MockICompetitionLeadTagRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICompetitionLeadTagRepositoryMockRecorder
}

// MockICompetitionLeadTagRepositoryMockRecorder is the mock recorder for MockICompetitionLeadTagRepository
type MockICompetitionLeadTagRepositoryMockRecorder struct {
	mock *MockICompetitionLeadTagRepository
}

// NewMockICompetitionLeadTagRepository creates a new mock instance
func NewMockICompetitionLeadTagRepository(ctrl *gomock.Controller) *MockICompetitionLeadTagRepository {
	mock := &MockICompetitionLeadTagRepository{ctrl: ctrl}
	mock.recorder = &MockICompetitionLeadTagRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICompetitionLeadTagRepository) EXPECT() *MockICompetitionLeadTagRepositoryMockRecorder {
	return m.recorder
}

// CreateCompetitionLeadTag mocks base method
func (m *MockICompetitionLeadTagRepository) CreateCompetitionLeadTag(tag *businesslogic.CompetitionLeadTag) error {
	ret := m.ctrl.Call(m, "CreateCompetitionLeadTag", tag)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCompetitionLeadTag indicates an expected call of CreateCompetitionLeadTag
func (mr *MockICompetitionLeadTagRepositoryMockRecorder) CreateCompetitionLeadTag(tag interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompetitionLeadTag", reflect.TypeOf((*MockICompetitionLeadTagRepository)(nil).CreateCompetitionLeadTag), tag)
}

// DeleteCompetitionLeadTag mocks base method
func (m *MockICompetitionLeadTagRepository) DeleteCompetitionLeadTag(tag businesslogic.CompetitionLeadTag) error {
	ret := m.ctrl.Call(m, "DeleteCompetitionLeadTag", tag)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCompetitionLeadTag indicates an expected call of DeleteCompetitionLeadTag
func (mr *MockICompetitionLeadTagRepositoryMockRecorder) DeleteCompetitionLeadTag(tag interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCompetitionLeadTag", reflect.TypeOf((*MockICompetitionLeadTagRepository)(nil).DeleteCompetitionLeadTag), tag)
}

// SearchCompetitionLeadTag mocks base method
func (m *MockICompetitionLeadTagRepository) SearchCompetitionLeadTag(criteria businesslogic.SearchCompetitionLeadTagCriteria) ([]businesslogic.CompetitionLeadTag, error) {
	ret := m.ctrl.Call(m, "SearchCompetitionLeadTag", criteria)
	ret0, _ := ret[0].([]businesslogic.CompetitionLeadTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchCompetitionLeadTag indicates an expected call of SearchCompetitionLeadTag
func (mr *MockICompetitionLeadTagRepositoryMockRecorder) SearchCompetitionLeadTag(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchCompetitionLeadTag", reflect.TypeOf((*MockICompetitionLeadTagRepository)(nil).SearchCompetitionLeadTag), criteria)
}

// UpdateCompetitionLeadTag mocks base method
func (m *MockICompetitionLeadTagRepository) UpdateCompetitionLeadTag(tag businesslogic.CompetitionLeadTag) error {
	ret := m.ctrl.Call(m, "UpdateCompetitionLeadTag", tag)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCompetitionLeadTag indicates an expected call of UpdateCompetitionLeadTag
func (mr *MockICompetitionLeadTagRepositoryMockRecorder) UpdateCompetitionLeadTag(tag interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCompetitionLeadTag", reflect.TypeOf((*MockICompetitionLeadTagRepository)(nil).UpdateCompetitionLeadTag), tag)
}
