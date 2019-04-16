// Code generated by MockGen. DO NOT EDIT.
// Source: ./businesslogic/event.go

// Package mock_businesslogic is a generated GoMock package.
package mock_businesslogic

import (
	businesslogic "github.com/DancesportSoftware/das/businesslogic"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIEventStatusRepository is a mock of IEventStatusRepository interface
type MockIEventStatusRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIEventStatusRepositoryMockRecorder
}

// MockIEventStatusRepositoryMockRecorder is the mock recorder for MockIEventStatusRepository
type MockIEventStatusRepositoryMockRecorder struct {
	mock *MockIEventStatusRepository
}

// NewMockIEventStatusRepository creates a new mock instance
func NewMockIEventStatusRepository(ctrl *gomock.Controller) *MockIEventStatusRepository {
	mock := &MockIEventStatusRepository{ctrl: ctrl}
	mock.recorder = &MockIEventStatusRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIEventStatusRepository) EXPECT() *MockIEventStatusRepositoryMockRecorder {
	return m.recorder
}

// GetEventStatus mocks base method
func (m *MockIEventStatusRepository) GetEventStatus() ([]businesslogic.EventStatus, error) {
	ret := m.ctrl.Call(m, "GetEventStatus")
	ret0, _ := ret[0].([]businesslogic.EventStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventStatus indicates an expected call of GetEventStatus
func (mr *MockIEventStatusRepositoryMockRecorder) GetEventStatus() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventStatus", reflect.TypeOf((*MockIEventStatusRepository)(nil).GetEventStatus))
}

// MockIEventRepository is a mock of IEventRepository interface
type MockIEventRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIEventRepositoryMockRecorder
}

// MockIEventRepositoryMockRecorder is the mock recorder for MockIEventRepository
type MockIEventRepositoryMockRecorder struct {
	mock *MockIEventRepository
}

// NewMockIEventRepository creates a new mock instance
func NewMockIEventRepository(ctrl *gomock.Controller) *MockIEventRepository {
	mock := &MockIEventRepository{ctrl: ctrl}
	mock.recorder = &MockIEventRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIEventRepository) EXPECT() *MockIEventRepositoryMockRecorder {
	return m.recorder
}

// SearchEvent mocks base method
func (m *MockIEventRepository) SearchEvent(criteria businesslogic.SearchEventCriteria) ([]businesslogic.Event, error) {
	ret := m.ctrl.Call(m, "SearchEvent", criteria)
	ret0, _ := ret[0].([]businesslogic.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchEvent indicates an expected call of SearchEvent
func (mr *MockIEventRepositoryMockRecorder) SearchEvent(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchEvent", reflect.TypeOf((*MockIEventRepository)(nil).SearchEvent), criteria)
}

// CreateEvent mocks base method
func (m *MockIEventRepository) CreateEvent(event *businesslogic.Event) error {
	ret := m.ctrl.Call(m, "CreateEvent", event)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEvent indicates an expected call of CreateEvent
func (mr *MockIEventRepositoryMockRecorder) CreateEvent(event interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockIEventRepository)(nil).CreateEvent), event)
}

// UpdateEvent mocks base method
func (m *MockIEventRepository) UpdateEvent(event businesslogic.Event) error {
	ret := m.ctrl.Call(m, "UpdateEvent", event)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEvent indicates an expected call of UpdateEvent
func (mr *MockIEventRepositoryMockRecorder) UpdateEvent(event interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEvent", reflect.TypeOf((*MockIEventRepository)(nil).UpdateEvent), event)
}

// DeleteEvent mocks base method
func (m *MockIEventRepository) DeleteEvent(event businesslogic.Event) error {
	ret := m.ctrl.Call(m, "DeleteEvent", event)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEvent indicates an expected call of DeleteEvent
func (mr *MockIEventRepositoryMockRecorder) DeleteEvent(event interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEvent", reflect.TypeOf((*MockIEventRepository)(nil).DeleteEvent), event)
}

// MockIEventDanceRepository is a mock of IEventDanceRepository interface
type MockIEventDanceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIEventDanceRepositoryMockRecorder
}

// MockIEventDanceRepositoryMockRecorder is the mock recorder for MockIEventDanceRepository
type MockIEventDanceRepositoryMockRecorder struct {
	mock *MockIEventDanceRepository
}

// NewMockIEventDanceRepository creates a new mock instance
func NewMockIEventDanceRepository(ctrl *gomock.Controller) *MockIEventDanceRepository {
	mock := &MockIEventDanceRepository{ctrl: ctrl}
	mock.recorder = &MockIEventDanceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIEventDanceRepository) EXPECT() *MockIEventDanceRepositoryMockRecorder {
	return m.recorder
}

// SearchEventDance mocks base method
func (m *MockIEventDanceRepository) SearchEventDance(criteria businesslogic.SearchEventDanceCriteria) ([]businesslogic.EventDance, error) {
	ret := m.ctrl.Call(m, "SearchEventDance", criteria)
	ret0, _ := ret[0].([]businesslogic.EventDance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchEventDance indicates an expected call of SearchEventDance
func (mr *MockIEventDanceRepositoryMockRecorder) SearchEventDance(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchEventDance", reflect.TypeOf((*MockIEventDanceRepository)(nil).SearchEventDance), criteria)
}

// CreateEventDance mocks base method
func (m *MockIEventDanceRepository) CreateEventDance(eventDance *businesslogic.EventDance) error {
	ret := m.ctrl.Call(m, "CreateEventDance", eventDance)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEventDance indicates an expected call of CreateEventDance
func (mr *MockIEventDanceRepositoryMockRecorder) CreateEventDance(eventDance interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEventDance", reflect.TypeOf((*MockIEventDanceRepository)(nil).CreateEventDance), eventDance)
}

// DeleteEventDance mocks base method
func (m *MockIEventDanceRepository) DeleteEventDance(eventDance businesslogic.EventDance) error {
	ret := m.ctrl.Call(m, "DeleteEventDance", eventDance)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEventDance indicates an expected call of DeleteEventDance
func (mr *MockIEventDanceRepositoryMockRecorder) DeleteEventDance(eventDance interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEventDance", reflect.TypeOf((*MockIEventDanceRepository)(nil).DeleteEventDance), eventDance)
}

// UpdateEventDance mocks base method
func (m *MockIEventDanceRepository) UpdateEventDance(eventDance businesslogic.EventDance) error {
	ret := m.ctrl.Call(m, "UpdateEventDance", eventDance)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEventDance indicates an expected call of UpdateEventDance
func (mr *MockIEventDanceRepositoryMockRecorder) UpdateEventDance(eventDance interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEventDance", reflect.TypeOf((*MockIEventDanceRepository)(nil).UpdateEventDance), eventDance)
}

// MockICompetitionEventTemplateRepository is a mock of ICompetitionEventTemplateRepository interface
type MockICompetitionEventTemplateRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICompetitionEventTemplateRepositoryMockRecorder
}

// MockICompetitionEventTemplateRepositoryMockRecorder is the mock recorder for MockICompetitionEventTemplateRepository
type MockICompetitionEventTemplateRepositoryMockRecorder struct {
	mock *MockICompetitionEventTemplateRepository
}

// NewMockICompetitionEventTemplateRepository creates a new mock instance
func NewMockICompetitionEventTemplateRepository(ctrl *gomock.Controller) *MockICompetitionEventTemplateRepository {
	mock := &MockICompetitionEventTemplateRepository{ctrl: ctrl}
	mock.recorder = &MockICompetitionEventTemplateRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICompetitionEventTemplateRepository) EXPECT() *MockICompetitionEventTemplateRepositoryMockRecorder {
	return m.recorder
}

// SearchCompetitionEventTemplates mocks base method
func (m *MockICompetitionEventTemplateRepository) SearchCompetitionEventTemplates(criteria businesslogic.SearchCompetitionEventTemplateCriteria) ([]businesslogic.CompetitionEventTemplate, error) {
	ret := m.ctrl.Call(m, "SearchCompetitionEventTemplates", criteria)
	ret0, _ := ret[0].([]businesslogic.CompetitionEventTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchCompetitionEventTemplates indicates an expected call of SearchCompetitionEventTemplates
func (mr *MockICompetitionEventTemplateRepositoryMockRecorder) SearchCompetitionEventTemplates(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchCompetitionEventTemplates", reflect.TypeOf((*MockICompetitionEventTemplateRepository)(nil).SearchCompetitionEventTemplates), criteria)
}
