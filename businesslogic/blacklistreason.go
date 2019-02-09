package businesslogic

import "time"

type PartnershipRequestBlacklistReason struct {
	ID              int
	Name            string
	Description     string
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}

type IPartnershipRequestBlacklistReasonRepository interface {
	GetPartnershipRequestBlacklistReasons() ([]PartnershipRequestBlacklistReason, error)
}
