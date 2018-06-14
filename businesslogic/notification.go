package businesslogic

import "time"

const (
	NewPartnershipRequestNotification = 1
)

// NotificationPreference stores the preference of how user would like to receive system-generated notification
type NotificationPreference struct {
}

type Notification struct {
	ID              int
	HasRead         bool
	DateTimeCreated time.Time
}
