package businesslogic_test

import (
	"errors"
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/mock/businesslogic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testDances = []int{1, 2}

func TestEvent_HasDance(t *testing.T) {
	event := businesslogic.NewEvent()

	dance_1 := 1
	dance_2 := 2
	dance_1_dup := 1

	event.AddDance(dance_1)
	event.AddDance(dance_2)

	assert.True(t, event.HasDance(dance_1_dup), "should has dance as long as value is matches")
}

func TestEvent_GetDances(t *testing.T) {
	event := businesslogic.NewEvent()

	dance_2 := 2
	dance_5 := 5
	dance_1 := 1

	event.AddDance(dance_2)
	event.AddDance(dance_5)
	event.AddDance(dance_1)

	assert.EqualValues(t, 3, len(event.GetDances()), "should only get unique dances")

	assert.EqualValues(t, 1, event.GetDances()[0])
	assert.EqualValues(t, 2, event.GetDances()[1])
	assert.EqualValues(t, 5, event.GetDances()[2])
}

func TestEvent_RemoveDance(t *testing.T) {
	event := businesslogic.NewEvent()

	dance_1 := 1
	dance_2 := 2
	event.AddDance(dance_1)
	event.AddDance(dance_2)

	assert.EqualValues(t, 2, len(event.GetDances()))

	event.RemoveDance(dance_1)

	assert.EqualValues(t, 1, len(event.GetDances()))
}

func TestEvent_Equivalent(t *testing.T) {
	event_1 := businesslogic.NewEvent()
	event_2 := businesslogic.NewEvent()

	assert.EqualValues(t, event_1, event_2)

	dance_1 := 1
	dance_2 := 2
	dance_3 := 3

	event_1.AddDance(dance_1)
	event_2.AddDance(dance_1)

	assert.True(t, event_1.EquivalentTo(*event_2))

	event_1.FederationID = 1
	event_2.FederationID = 1
	assert.True(t, event_1.EquivalentTo(*event_2))

	event_1.DivisionID = 3
	event_2.DivisionID = 4
	assert.False(t, event_1.EquivalentTo(*event_2))

	event_1.AddDance(dance_2)
	assert.False(t, event_1.EquivalentTo(*event_2))

	event_2.AddDance(dance_3)
	assert.False(t, event_1.EquivalentTo(*event_2))

}

func TestCreateEvent(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	event := businesslogic.NewEvent()
	event.SetDances(testDances)

	eventRepository := mock_businesslogic.NewMockIEventRepository(mockCtrl)
	eventRepository.EXPECT().CreateEvent(event).Return(errors.New("should not allow wrong events to be created"))

	err := eventRepository.CreateEvent(event)
	assert.NotNil(t, err, "creating an uninitialized event should result in an error")
}

func TestCreateEvent_BadCompetitionStatus(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	event := businesslogic.NewEvent()
	event.CompetitionID = 22
	event.FederationID = 3
	event.DivisionID = 19
	event.AgeID = 19
	event.ProficiencyID = 10
	event.StyleID = 7
	event.StatusID = businesslogic.EVENT_STATUS_DRAFT

	event.SetDances(testDances)

	compRepository := mock_businesslogic.NewMockICompetitionRepository(mockCtrl)
	eventRepository := mock_businesslogic.NewMockIEventRepository(mockCtrl)
	eventDanceRepo := mock_businesslogic.NewMockIEventDanceRepository(mockCtrl)

	expectedCompetition := businesslogic.Competition{ID: 22}
	expectedCompetition.UpdateStatus(businesslogic.CompetitionStatusClosedRegistration)

	assert.Equal(t, businesslogic.CompetitionStatusClosedRegistration, expectedCompetition.GetStatus(), "competition status should be updated")

	compRepository.EXPECT().SearchCompetition(gomock.Any()).Return([]businesslogic.Competition{
		expectedCompetition,
	}, nil)

	err := businesslogic.CreateEvent(*event, compRepository, eventRepository, eventDanceRepo)
	assert.NotNil(t, err, "creating event for competition that is closed for registration should throw an error")
}
