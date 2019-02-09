package businesslogic

import "time"

type PartnershipRequestStatus struct {
	ID              int
	Code            string
	Description     string
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}

type IPartnershipRequestStatusRepository interface {
	GetPartnershipRequestStatus() ([]PartnershipRequestStatus, error)
}
