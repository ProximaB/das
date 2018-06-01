package businesslogic

import (
	"time"
)

type SearchPartnershipRequestBlacklistCriteria struct {
	ReporterID    int
	BlockedUserID int
	Whitelisted   bool
	ReasonID      int
}

type PartnershipRequestBlacklistEntry struct {
	ID                int
	ReporterID        int
	BlockedUserID     int
	BlackListReasonID int
	Detail            string
	Whitelisted       bool
	CreateUserID      int
	DateTimeCreated   time.Time
	UpdateUserID      int
	DateTimeUpdated   time.Time
}

type IPartnershipRequestBlacklistRepository interface {
	SearchPartnershipRequestBlacklist(criteria SearchPartnershipRequestBlacklistCriteria) ([]PartnershipRequestBlacklistEntry, error)
	CreatePartnershipRequestBlacklist(blacklist *PartnershipRequestBlacklistEntry) error
	DeletePartnershipRequestBlacklist(blacklist PartnershipRequestBlacklistEntry) error
	UpdatePartnershipRequestBlacklist(blacklist PartnershipRequestBlacklistEntry) error
}

func GetBlacklistedAccountsForUser(userID int, accountRepo IAccountRepository,
	blacklistRepo IPartnershipRequestBlacklistRepository) ([]Account, error) {
	blacklist := make([]Account, 0)
	entries, err := blacklistRepo.SearchPartnershipRequestBlacklist(SearchPartnershipRequestBlacklistCriteria{ReporterID: userID, Whitelisted: false})
	if err != nil {
		return blacklist, err
	}
	for _, each := range entries {
		blacklist = append(blacklist, GetAccountByID(each.BlockedUserID, accountRepo))
	}
	return blacklist, err
}
