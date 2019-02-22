package skating

const (
	ALGORITHM_SKATING = "skating"
)

type DanceScoreSheet struct {
}

func NewDanceScoreSheet() DanceScoreSheet {
	return DanceScoreSheet{}
}

func (sheet *DanceScoreSheet) AddJudgeMarks(marks JudgeMarks) {}

func (sheet DanceScoreSheet) CalculateDancePlacements(algorithm string) ([][]int, error) {
	return make([][]int, 0), nil
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
