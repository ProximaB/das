package scheduler

import (
	"errors"
	"math"
)

// EventSchedule takes parameters of an event and estimates the run time of an event
type EventScheduler struct {
	totalEntries    int
	totalDances     int
	floorCapacity   int
	recallRate      int
	danceDuration   int
	targetFinalSize int
}

func (event *EventScheduler) SetTotalEntries(entries int) error {
	if entries < 1 {
		return errors.New("entries must be at least 1")
	}
	event.totalEntries = entries
	return nil
}

func (event *EventScheduler) SetTotalDances(dances int) error {
	if dances < 1 {
		return errors.New("dances must be at least 1")
	}
	event.totalDances = dances
	return nil
}

func (event *EventScheduler) SetDanceDuration(seconds int) error {
	if seconds < 1 {
		return errors.New("dances must be at least 1 second")
	}
	event.danceDuration = seconds
	return nil
}

func (event *EventScheduler) SetFloorCapacity(capacity int) error {
	if capacity < 1 {
		return errors.New("floor capacity must be larger than 0")
	}
	event.floorCapacity = capacity
	return nil
}

func (event *EventScheduler) SetRecallRate(rate int) error {
	if rate < 1 || rate > 99 {
		return errors.New("recall rate must be between 1 and 99")
	}
	event.recallRate = rate
	return nil
}

func (event *EventScheduler) SetTargetFinalSize(target int) error {
	if target < 1 {
		return errors.New("must have at least one competitor in the final round")
	}
	event.targetFinalSize = target
	return nil
}

func (event EventScheduler) EstimateRounds() (int, error) {
	if event.targetFinalSize < 1 {
		return 0, errors.New("must have at least one competitor in the final round")
	}
	if event.totalEntries < 1 {
		return 0, errors.New("not enough entries")
	}
	if event.recallRate < 1 || event.recallRate > 99 {
		return 0, errors.New("recall rate must be between 1 and 99")
	}

	// if the total entry is smaller than the final round size, then there should be only one round
	if event.totalEntries <= event.targetFinalSize {
		return 1, nil
	}

	finalRate := float64(event.targetFinalSize) / float64(event.totalEntries)
	recallRate := float64(event.recallRate) / 100.0
	rawRounds := math.Log2(finalRate) / math.Log2(recallRate)

	reducedRounds := int(rawRounds)
	if reducedRounds == 0 {
		reducedRounds = 1 // minimum should be 1 round
	}

	// special situation when it takes only one round to get to the final round.
	if finalRate >= recallRate {
		reducedRounds += 1
	}

	if (rawRounds - float64(reducedRounds)) > 0 {
		reducedRounds += 1
	}
	return reducedRounds, nil
}

func (event EventScheduler) EstimateRoundEntries() ([]int, error) {
	rosters := make([]int, 0)

	rounds, roundsErr := event.EstimateRounds()
	if roundsErr != nil {
		return rosters, roundsErr
	}

	rosters = append(rosters, event.totalEntries)
	for i := 0; i < rounds; i++ {
		entry := int(float64(event.totalEntries) * math.Pow(float64(event.recallRate)/100.0, float64(i+1)))
		if entry >= event.targetFinalSize {
			rosters = append(rosters, entry)
		}
	}
	return rosters, nil
}

// EstimateRuntime estimates the total time to run this event and return the number of minutes and an error (if happens)
func (event EventScheduler) EstimateRuntime() (int, error) {
	rounds, err := event.EstimateRounds()
	if err != nil {
		return 0, err
	}

	rosters, err := event.EstimateRoundEntries()
	if err != nil {
		return 0, err
	}

	totalRuntime := 0.0 // unit is seconds
	for i := 0; i < rounds; i++ {
		roundHeats := float64(rosters[i]) / float64(event.floorCapacity)
		if roundHeats < 1 {
			roundHeats = 1
		}

		roundRuntime := roundHeats * float64(event.totalDances*event.danceDuration)
		totalRuntime += roundRuntime
	}

	preciseFinalRuntime := totalRuntime / 60.0
	rawFinalRuntime := int(preciseFinalRuntime)
	if preciseFinalRuntime > float64(rawFinalRuntime) {
		rawFinalRuntime += 1
	}

	return rawFinalRuntime, nil
}
