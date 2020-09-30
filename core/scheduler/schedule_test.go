package scheduler_test

import (
	"github.com/ProximaB/das/core/scheduler"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestEvent_EstimateRounds(t *testing.T) {
	evt := scheduler.EventScheduler{}
	evt.SetRecallRate(60)
	evt.SetTotalEntries(80)
	evt.SetTargetFinalSize(6)

	rounds, err := evt.EstimateRounds()
	assert.Nil(t, err, "should not return an error when rounds ")
	assert.Equal(t, 6, rounds, "80 couples with 60% recall rate targeting 6-couple final should be 6 rounds")

	evt.SetTargetFinalSize(3)
	rounds, _ = evt.EstimateRounds()
	assert.Equal(t, 7, rounds, "80 couples with 60% recall rate targeting 4-couple final should be 7 rounds")

	evt.SetRecallRate(60)
	evt.SetTotalEntries(9)
	evt.SetTargetFinalSize(6)
	rounds, _ = evt.EstimateRounds()
	assert.Equal(t, 2, rounds)

	evt.SetRecallRate(60)
	evt.SetTotalEntries(4)
	evt.SetTargetFinalSize(6)
	rounds, _ = evt.EstimateRounds()
	assert.Equal(t, 1, rounds)
}

func TestEventScheduler_EstimateRoundEntries(t *testing.T) {
	evt := scheduler.EventScheduler{}
	evt.SetRecallRate(60)
	evt.SetTotalEntries(80)
	evt.SetTargetFinalSize(6)
	evt.SetFloorCapacity(20)

	rosters, err := evt.EstimateRoundEntries()
	assert.Equal(t, 6, len(rosters), "should have at least 6 rounds with 80 couples")
	assert.Nil(t, err)

	evt.SetTotalEntries(10)
	rosters, err = evt.EstimateRoundEntries()
	assert.Equal(t, 2, len(rosters), "should have at least 2 rounds with 10 couples")
}

func TestEventScheduler_EstimateRuntime(t *testing.T) {
	evt := scheduler.EventScheduler{}
	evt.SetRecallRate(60)
	evt.SetTotalEntries(80)
	evt.SetTargetFinalSize(6)
	evt.SetFloorCapacity(20)
	evt.SetTotalDances(4)
	evt.SetDanceDuration(90)

	rosters, _ := evt.EstimateRoundEntries()
	log.Println(rosters)
	runtime, err := evt.EstimateRuntime()
	assert.Nil(t, err)
	assert.Equal(t, 65, runtime, "should take 65 minutes to run this event")

	evt.SetTotalEntries(10)
	evt.SetTotalDances(5)
	evt.SetTargetFinalSize(6)
	evt.SetFloorCapacity(20)
	evt.SetRecallRate(60)
	evt.SetDanceDuration(90)
	rounds, _ := evt.EstimateRounds()
	assert.Equal(t, 2, rounds)
	runtime, _ = evt.EstimateRuntime()
	assert.Equal(t, 15, runtime)
}
