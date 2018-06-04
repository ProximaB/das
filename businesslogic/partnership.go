package businesslogic

import (
	"time"
)

const (
	PARTNERSHIP_REQUEST_RECEIVED = 1
	PARTNERSHIP_REQUEST_SENT     = 2
)

const (
	PARTNERSHIP_ROLE_LEAD   = "LEAD"
	PARTNERSHIP_ROLE_FOLLOW = "FOLLOW"
)

type Partnership struct {
	PartnershipID   int
	LeadID          int
	FollowID        int
	Lead            Account
	Follow          Account
	SameSex         bool
	FavoriteLead    bool
	FavoriteFollow  bool
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}

type IPartnershipRepository interface {
	CreatePartnership(partnership *Partnership) error
	SearchPartnership(criteria SearchPartnershipCriteria) ([]Partnership, error)
	UpdatePartnership(partnership Partnership) error
	DeletePartnership(partnership Partnership) error
}

type SearchPartnershipCriteria struct {
	PartnershipID int `schema:"id"`
	LeadID        int `schema:"lead"`
	FollowID      int `schema:"follow"`
}
