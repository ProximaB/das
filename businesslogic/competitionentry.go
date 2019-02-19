package businesslogic

import "time"

// BaseCompetitionEntry is the entry for a competition (not including events) and is a
// base entry for more specific entry such as AthleteCompetitionEntry, PartnershipCompetitionEntry,
// and AdjudicatorCompetitionEntry
type BaseCompetitionEntry struct {
	CompetitionID    int
	CheckInIndicator bool
	DateTimeCheckIn  time.Time
	CreateUserID     int
	DateTimeCreated  time.Time
	UpdateUserID     int
	DateTimeUpdated  time.Time
}
