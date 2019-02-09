package businesslogic

import "time"

// RoundOrder defines the order of the round, from lowest to the highest
type RoundOrder struct {
	ID              int
	Rank            int
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}

// Round defines the round for each event
type Round struct {
	ID              int
	EventID         int
	Order           RoundOrder
	Entries         []EventEntry
	StartTime       time.Time
	EndTime         time.Time
	DateTimeCreated time.Time
	CreateUserID    int
	DateTimeUpdated time.Time
	UpdateUserID    int
}

// SearchRoundCriteria specifies the parameters that can be used to search Rounds in a Repository
type SearchRoundCriteria struct {
	CompetitionID int
	EventID       int
	RoundOrderID  int
}

// IRoundRepository specifies the interface that a Round Repository should implement
type IRoundRepository interface {
	CreateRound(round *Round) error
	DeleteRound(round Round) error
	SearchRound(criteria SearchRoundCriteria) ([]Round, error)
	UpdateRound(round Round) error
}
