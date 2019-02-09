package businesslogic

import (
	"time"
)

type Division struct {
	ID              int
	Name            string
	Abbreviation    string
	Description     string
	FederationID    int
	Note            string
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

type SearchDivisionCriteria struct {
	ID           int    `schema:"id"`
	Name         string `schema:"name"`
	FederationID int    `schema:"federation"`
}

type IDivisionRepository interface {
	CreateDivision(division *Division) error
	SearchDivision(criteria SearchDivisionCriteria) ([]Division, error)
	UpdateDivision(division Division) error
	DeleteDivision(division Division) error
}
