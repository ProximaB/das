package businesslogic

import "time"

// UserPreference stores the basic preferences of user
type UserPreference struct {
	ID              int
	AccountID       int
	DefaultRole     int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

type SearchUserPreferenceCriteria struct {
	AccountID int
}

type IUserPreferenceRepository interface {
	CreatePreference(preference *UserPreference) error
	SearchPreference(criteria SearchUserPreferenceCriteria) ([]UserPreference, error)
	UpdatePreference(preference UserPreference) error
}
