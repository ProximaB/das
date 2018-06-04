package reference

import (
	"time"
)

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

type SearchAgeCriteria struct {
	FederationID int `schema:"federation"`
	DivisionID   int `schema:"division"`
	AgeID        int `schema:"id"`
}

type IAgeRepository interface {
	CreateAge(age *Age) error
	SearchAge(criteria SearchAgeCriteria) ([]Age, error)
	UpdateAge(age Age) error
	DeleteAge(age Age) error
}
