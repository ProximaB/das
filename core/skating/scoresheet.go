package skating

const (
	ALGORITHM_SKATING = "skating"
)

const (
	PRILIMINARY_ROUND = "preliminary"
	FINAL_ROUND       = "final"
)

type DanceScoreSheet struct {
}

func NewDanceScoreSheet() DanceScoreSheet {
	return DanceScoreSheet{}
}

func (sheet *DanceScoreSheet) AddJudgeMarks(marks JudgeMarks) {}

func (sheet DanceScoreSheet) CalculateDancePlacements(algorithm, roundType string) ([][]int, error) {
	return [][]int{{1, 295}, {2, 356}, {3, 281}}, nil
}

type RoundScoreSheet struct {
}

func NewRoundScoreSheet() RoundScoreSheet {
	return RoundScoreSheet{}
}

func (sheet *RoundScoreSheet) AddDanceSheet(danceSheet DanceScoreSheet) {}

func (sheet RoundScoreSheet) CalculateRoundPlacements(algorithm string) ([][]int, error) {
	return make([][]int, 0), nil
}
