package businesslogic

import "time"

type CompetitionOfficial struct {
	ID              int
	Competition     Competition
	Official        Account   // the ID for AccountRole
	OfficialRoleID  int       // the ID for AccountType
	EffectiveFrom   time.Time // have privileged access to competition data
	EffectiveUntil  time.Time
	AssignedBy      int // ID of an AccountRole object, must be an organizer. TODO: may use invitation instead of assignment
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

// Active checks if the status of this position is still active.
func (official CompetitionOfficial) Active() bool {
	return time.Now().Before(official.EffectiveUntil) && time.Now().After(official.EffectiveFrom)
}

type SearchCompetitionOfficialCriteria struct {
	ID             int
	CompetitionID  int
	OfficialID     string
	OfficialRoleID int
}

type ICompetitionOfficialRepository interface {
	CreateCompetitionOfficial(official *CompetitionOfficial) error
	DeleteCompetitionOfficial(official CompetitionOfficial) error
	SearchCompetitionOfficial(criteria SearchCompetitionOfficialCriteria) ([]CompetitionOfficial, error)
	UpdateCompetitionOfficial(official CompetitionOfficial) error
}
