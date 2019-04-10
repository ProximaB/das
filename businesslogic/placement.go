package businesslogic

import "time"

// Placement defines the minimal data for a dance placement: Round + Dance + Adjudicator + Partnership + Placement together
// defines a unique placement
type Placement struct {
	ID                        int
	AdjudicatorRoundEntryID   int
	PartnershipRoundEntryID   int
	RoundDanceID              int
	PreliminaryRoundIndicator bool
	Placement                 int
	CreateUserID              int
	DateTimeCreated           time.Time
	UpdateUserID              int
	DateTimeUpdated           time.Time
}

// SearchPlacementCriteria specifies the parameters that can be used to search Placement in a repository
type SearchPlacementCriteria struct {
	CompetitionID int
	EventID       int
	PartnershipID int
}

// IPlacementRepository specifies the functions that a Placement Repository should implement
type IPlacementRepository interface {
	CreatePlacement(placement *Placement) error
	DeletePlacement(placement Placement) error
	SearchPlacement(criteria SearchPlacementCriteria) ([]Placement, error)
	UpdatePlacement(placement Placement) error
}
