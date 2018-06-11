package referencebll

import (
	"time"
)

// Age contains data for the age category requirement for events
// Age is associated with Division, which is associated with Federation
type Age struct {
	ID              int
	Name            string
	Description     string
	DivisionID      int
	Enforced        bool // if required, AgeMinimum and AgeMaximum must have non-zero value
	AgeMinimum      int
	AgeMaximum      int
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

// SearchAgeCriteria provides parameters when searching Age in IAgeRepository
type SearchAgeCriteria struct {
	FederationID int `schema:"federation"`
	DivisionID   int `schema:"division"`
	AgeID        int `schema:"id"`
}

// IAgeRepository provides an interface for other businesslogic code to access Age data
type IAgeRepository interface {
	CreateAge(age *Age) error
	SearchAge(criteria SearchAgeCriteria) ([]Age, error)
	UpdateAge(age Age) error
	DeleteAge(age Age) error
}
