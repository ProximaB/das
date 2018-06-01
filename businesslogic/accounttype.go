package businesslogic

import "time"

const (
	ACCOUNT_TYPE_NOAUTH        = 0
	ACCOUNT_TYPE_ATHLETE       = 1
	ACCOUNT_TYPE_ADJUDICATOR   = 2
	ACCOUNT_TYPE_SCRUTINEER    = 3
	ACCOUNT_TYPE_ORGANIZER     = 4
	ACCOUNT_TYPE_DECKCAPTAIN   = 5
	ACCOUNT_TYPE_EMCEE         = 6
	ACCOUNT_TYPE_ADMINISTRATOR = 7
)

type AccountType struct {
	ID              int
	Name            string
	Description     string
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}

type IAccountTypeRepository interface {
	GetAccountTypes() ([]AccountType, error)
}
