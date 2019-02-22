package skating

import (
	"errors"
	"fmt"
	"sort"
)

var DUPLICATE_PLACEMENT_ERROR = func(placement int) error {
	return errors.New(fmt.Sprintf("placement %v is already taken by someone else", placement))
}

var INVALID_PLACEMENT_ERROR = func(coupleID, placement int) error {
	return errors.New(fmt.Sprintf("couple %v of has invalid placement: %v", coupleID, placement))
}

var PLACEMENT_OUTOFRANGE_ERROR = func(coupleID, placement int) error {
	return errors.New(fmt.Sprintf("placement of couple %v is out of range: %v", coupleID, placement))
}

type JudgeMarks struct {
	callbacks  map[int]bool // key: coupleID, value: recall
	placements map[int]int  // key: placement, value: coupleID
	roundSize  int
}

func NewJudgeMarks(roundSize int) JudgeMarks {
	return JudgeMarks{
		callbacks:  make(map[int]bool),
		placements: make(map[int]int),
		roundSize:  roundSize,
	}
}

func (marks *JudgeMarks) AddCallback(coupleID int, callback bool) {
	marks.callbacks[coupleID] = callback
}

// AddPlacement adds judge's placement of couple to the marks. This function ensures that all couples receives
// a valid placement.
//
// This method ensures compliance with Rule 4.
func (marks *JudgeMarks) AddPlacement(coupleID, placement int) error {
	if marks.placements[placement] != 0 {
		return DUPLICATE_PLACEMENT_ERROR(placement)
	}
	if placement < 1 {
		return INVALID_PLACEMENT_ERROR(coupleID, placement)
	}
	if placement > marks.roundSize {
		return PLACEMENT_OUTOFRANGE_ERROR(coupleID, placement)
	}
	marks.placements[placement] = coupleID
	return nil
}

// GetCallbacks retrieves all the couples that are called back by the adjudicator.
//
// This method ensures compliance with Rule 1.
func (marks JudgeMarks) GetCallbacks() []int {
	callbacks := make([]int, 0)
	for couple, yes := range marks.callbacks {
		if yes {
			callbacks = append(callbacks, couple)
		}
	}
	sort.Ints(callbacks)
	return callbacks
}

func (marks JudgeMarks) GetPlacements() [][]int {
	placements := make([][]int, 0)
	for place, couple := range marks.placements {
		placements = append(placements, []int{place, couple})
	}
	return placements
}

func (marks JudgeMarks) GetCouples() []int {
	couples := make([]int, 0)
	for each := range marks.callbacks {
		couples = append(couples, each)
	}
	sort.Ints(couples)
	return couples
}
