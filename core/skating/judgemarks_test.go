package skating_test

import (
	"github.com/DancesportSoftware/das/core/skating"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJudgeMarks_NoCallbacks(t *testing.T) {
	marks := skating.NewJudgeMarks(0)
	assert.Equal(t, 0, len(marks.GetCallbacks()), "marks of 0 couple should have 0 callback")
}

func TestJudgeMarks_ValidCallbacks(t *testing.T) {
	marks := skating.NewJudgeMarks(2)

	marks.AddCallback(101, true)
	marks.AddCallback(102, true)

	assert.Equal(t, 2, len(marks.GetCallbacks()), "total callback should be the sum of all callbacks")
	assert.Equal(t,[]int{101, 102}, marks.GetCallbacks(), "callbacks should return the ID of couples that are recalled")
}

func TestJudgeMarks_HasIgnoredCouples(t *testing.T) {
	marks := skating.NewJudgeMarks(3)

	marks.AddCallback(101, true)
	marks.AddCallback(102, true)
	marks.AddCallback(103, false)

	assert.Equal(t, 2, len(marks.GetCallbacks()), "total callback should be the sum of all callbacks")
	assert.Equal(t,[]int{101, 102}, marks.GetCallbacks(), "callbacks should return the ID of couples that are recalled")
}

func TestJudgeMarks_GetCouples(t *testing.T) {
	marks := skating.NewJudgeMarks(3)

	marks.AddCallback(103, true)
	marks.AddCallback(105, true)
	marks.AddCallback(101, false)

	assert.Equal(t, 3,len(marks.GetCouples()), "should return all couples that are placed regardless of callback")
	assert.Equal(t, []int{101, 103, 105}, marks.GetCouples(), "should return all couples in a sorted order")
}

func TestJudgeMarks_AddUniquePlacement (t *testing.T) {
	marks := skating.NewJudgeMarks(6)

	marks.AddPlacement(101, 2)
	marks.AddPlacement(102, 1)
	marks.AddPlacement(103, 5)
	marks.AddPlacement(104, 3)
	marks.AddPlacement(105, 4)
	marks.AddPlacement(106, 6)

	assert.Equal(t, 6, len(marks.GetPlacements()), "should get all the placements of couples")
}

func TestJudgeMarks_AddDuplicatePlacements (t *testing.T) {
	marks := skating.NewJudgeMarks(3)

	err := marks.AddPlacement(101, 2)
	err = marks.AddPlacement(102, 1)
	err = marks.AddPlacement(103, 2)
	assert.Equal(t, skating.DUPLICATE_PLACEMENT_ERROR(2), err)
}

func TestJudgeMarks_InvalidPlacements (t *testing.T) {
	marks := skating.NewJudgeMarks(2)

	err := marks.AddPlacement(101, -2)
	assert.Equal(t, skating.INVALID_PLACEMENT_ERROR(101, -2), err)
	err = marks.AddPlacement(102, 15)
	assert.Equal(t, skating.PLACEMENT_OUTOFRANGE_ERROR(102, 15), err)
}