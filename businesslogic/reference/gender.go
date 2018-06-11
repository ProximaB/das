package referencebll

import (
	"time"
)

const (
	GENDER_MALE   = 2
	GENDER_FEMALE = 1
)

type IGenderRepository interface {
	GetAllGenders() ([]Gender, error)
}

type Gender struct {
	ID              int
	Name            string
	Abbreviation    string
	Description     string
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}
