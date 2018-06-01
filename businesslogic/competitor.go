package businesslogic

type ICompetitorBehavior interface {
}

type SoloAthlete struct {
	ICompetitorBehavior
}

type Couple struct {
	ICompetitorBehavior
}
