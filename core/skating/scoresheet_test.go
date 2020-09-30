package skating_test

import (
	"github.com/ProximaB/das/core/skating"
	"github.com/stretchr/testify/assert"
	"testing"
)

//
//OSU DanceSport Classic 2014 2014
//Collegiate Novice Latin  Final
//Cha Cha
// 	10	11	12	13	14	 	1	1-2	1-3	P
//281	1	2	2	2	1	 	2	5		2
//356	2	1	1	1	3	 	3			1
//295	3	3	3	3	2	 		1	5	3
//
//Rumba
// 	10	11	12	13	14	 	1	1-2	1-3	P
//281	2	3	2	3	2	 		3(6)		3
//356	3	1	1	2	3	 	2	3(4)		2
//295	1	2	3	1	1	 	3			1
//
//Samba
// 	10	11	12	13	14	 	1	1-2	1-3	P
//281	1	3	1	3	1	 	3			1
//356	3	2	2	1	3	 	1	3		3
//295	2	1	3	2	2	 	1	4		2
//
//Summary
// 	C	R	S	 	Tot.	R10	R11(1)	R11(2)	R10	R11(2)	Res
//281	2	3	1	 	6	1(1)	5	11			1
//356	1	2	3	 	6	1(1)	6	10	2(3)	10	2
//295	3	1	2	 	6	1(1)	4	9	2(3)	9	3
//
func TestDanceScoreSheet_CalculateDancePlacements(t *testing.T) {
	sheet := skating.NewDanceScoreSheet()

	// use rumba dance
	marks1 := skating.NewJudgeMarks(3)
	marks1.AddPlacement(281, 2)
	marks1.AddPlacement(356, 3)
	marks1.AddPlacement(295, 1)

	marks2 := skating.NewJudgeMarks(3)
	marks2.AddPlacement(281, 3)
	marks2.AddPlacement(356, 1)
	marks2.AddPlacement(295, 2)

	marks3 := skating.NewJudgeMarks(3)
	marks3.AddPlacement(281, 2)
	marks3.AddPlacement(356, 1)
	marks3.AddPlacement(295, 3)

	marks4 := skating.NewJudgeMarks(3)
	marks4.AddPlacement(281, 3)
	marks4.AddPlacement(256, 2)
	marks4.AddPlacement(295, 1)

	marks5 := skating.NewJudgeMarks(3)
	marks5.AddPlacement(281, 2)
	marks5.AddPlacement(356, 3)
	marks5.AddPlacement(295, 1)

	sheet.AddJudgeMarks(marks1)
	sheet.AddJudgeMarks(marks2)
	sheet.AddJudgeMarks(marks3)

	ranks, err := sheet.CalculateDancePlacements(skating.ALGORITHM_SKATING, skating.FINAL_ROUND)
	assert.Equal(t, [][]int{{1, 295}, {2, 356}, {3, 281}}, ranks)
	assert.Nil(t, err)
}

/*
func TestRoundScoreSheet_NewRoundScoreSheet (t *testing.T) {
	roundSheet := skating.NewRoundScoreSheet()

	danceSheetA := skating.NewDanceScoreSheet()
	danceSheetB := skating.NewDanceScoreSheet()
	danceSheetC := skating.NewDanceScoreSheet()


	judgeMarksA1 := skating.NewJudgeMarks()
	judgeMarksA2 := skating.NewJudgeMarks()
	judgeMarksA3 := skating.NewJudgeMarks()
	judgeMarksB1 := skating.NewJudgeMarks()
	judgeMarksB2 := skating.NewJudgeMarks()
	judgeMarksB3 := skating.NewJudgeMarks()
	judgeMarksC1 := skating.NewJudgeMarks()
	judgeMarksC2 := skating.NewJudgeMarks()
	judgeMarksC3 := skating.NewJudgeMarks()

	danceSheetA.AddJudgeMarks(judgeMarksA1)
	danceSheetA.AddJudgeMarks(judgeMarksA2)
	danceSheetA.AddJudgeMarks(judgeMarksA3)
	danceSheetB.AddJudgeMarks(judgeMarksB1)
	danceSheetB.AddJudgeMarks(judgeMarksB2)
	danceSheetB.AddJudgeMarks(judgeMarksB3)
	danceSheetC.AddJudgeMarks(judgeMarksC1)
	danceSheetC.AddJudgeMarks(judgeMarksC2)
	danceSheetC.AddJudgeMarks(judgeMarksC3)

	roundSheet.AddDanceSheet(danceSheetA)
	roundSheet.AddDanceSheet(danceSheetB)
	roundSheet.AddDanceSheet(danceSheetC)

	ranks, err := roundSheet.CalculateRoundPlacements(skating.ALGORITHM_SKATING)
}
*/
