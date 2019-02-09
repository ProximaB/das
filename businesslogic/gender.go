package businesslogic

import (
	"time"
)

const (
	GENDER_MALE    = 2
	GENDER_FEMALE  = 1
	GENDER_UNKNOWN = 3 // registering a new account no longer requires specifying gender
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
