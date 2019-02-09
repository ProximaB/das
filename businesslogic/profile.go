package businesslogic

import "time"

// AthleteProfile specifies the data that is contained in an AthleteProfile
type AthleteProfile struct {
	ID              int
	AccountID       int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

type SearchAthleteProfileCriteria struct{}

type IAthleteProfileRepository interface {
	CreateProfile(profile *AthleteProfile) error
	UpdateProfile(profile AthleteProfile) error
	SearchProfile(criteria SearchAthleteProfileCriteria) ([]AthleteProfile, error)
}

type AdjudicatorProfile struct {
	ID              int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

type SearchAdjudicatorProfileCriteria struct{}

type IAdjudicatorProfileRepository interface {
	CreateProfile(profile *AdjudicatorProfile) error
	UpdateProfile(profile AdjudicatorProfile) error
	SearchProfile(criteria SearchAdjudicatorProfileCriteria) ([]AdjudicatorProfile, error)
}

type OrganizerProfile struct {
	ID              int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

type SearchOrganizerProfileCriteria struct{}

type IOrganizerProfileRepository interface {
	CreateProfile(profile *OrganizerProfile) error
	UpdateProfile(profile OrganizerProfile) error
	SearchProfile(criteria SearchOrganizerProfileCriteria) ([]OrganizerProfile, error)
}

type ScrutineerProfile struct {
	ID              int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

type SearchScrutineerProfileCriteria struct{}

type IScrutineerProfileRepository interface {
	CreateProfile(profile *ScrutineerProfile) error
	UpdateProfile(profile ScrutineerProfile) error
	SearchProfile(criteria SearchScrutineerProfileCriteria) ([]ScrutineerProfile, error)
}

type DeckCaptainProfile struct {
	ID              int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

type SearchDeckCaptainProfileCriteria struct{}

type IDeckCaptainProfileRepository interface {
	CreateProfile(profile *DeckCaptainProfile) error
	UpdateProfile(profile DeckCaptainProfile) error
	SearchProfile(criteria SearchDeckCaptainProfileCriteria) ([]DeckCaptainProfile, error)
}

type EmceeProfile struct {
	ID              int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

type SearchEmceeProfileCriteria struct{}

type IEmceeProfileRepository interface {
	CreateProfile(profile *EmceeProfile) error
	UpdateProfile(profile EmceeProfile) error
	SearchProfile(criteria SearchEmceeProfileCriteria) ([]EmceeProfile, error)
}

// AdministratorProfile stores the user profile and preference as an Administrator role. Administrator profile should
// not be created through web services but through direct database insertion.
type AdministratorProfile struct {
	ID              int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}
