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

// IPartnershipRequestBlacklistRepository defines the interface that a partnership request blacklist repository should implement
type IPartnershipRequestBlacklistRepository interface {
	SearchPartnershipRequestBlacklist(criteria SearchPartnershipRequestBlacklistCriteria) ([]PartnershipRequestBlacklistEntry, error)
	CreatePartnershipRequestBlacklist(blacklist *PartnershipRequestBlacklistEntry) error
	DeletePartnershipRequestBlacklist(blacklist PartnershipRequestBlacklistEntry) error
	UpdatePartnershipRequestBlacklist(blacklist PartnershipRequestBlacklistEntry) error
}

// GetBlacklistedAccounts searches all records of Blacklist reports, finds blacklisted accounts, and returns those accounts
// that were blacklisted
func (self Account) GetBlacklistedAccounts(accountRepo IAccountRepository, blacklistRepo IPartnershipRequestBlacklistRepository) ([]Account, error) {
	blacklist := make([]Account, 0)
	entries, err := blacklistRepo.SearchPartnershipRequestBlacklist(SearchPartnershipRequestBlacklistCriteria{ReporterID: self.ID, Whitelisted: false})
	if err != nil {
		return blacklist, err
	}
	for _, each := range entries {
		blacklist = append(blacklist, GetAccountByID(each.BlockedUserID, accountRepo))
	}
	return blacklist, err
}
