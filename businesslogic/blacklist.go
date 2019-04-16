package businesslogic

import (
	"time"
)

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

type SearchPartnershipRequestBlacklistCriteria struct {
	ReporterID    int
	BlockedUserID int
	Whitelisted   bool
	ReasonID      int
}

type PartnershipRequestBlacklistEntry struct {
	ID              int
	Reporter        Account
	BlockedUser     Account
	BlockedReason   PartnershipRequestBlacklistReason
	Detail          string
	Whitelisted     bool
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
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
func (account Account) GetBlacklistedAccounts(accountRepo IAccountRepository, blacklistRepo IPartnershipRequestBlacklistRepository) ([]Account, error) {
	blacklist := make([]Account, 0)
	entries, err := blacklistRepo.SearchPartnershipRequestBlacklist(SearchPartnershipRequestBlacklistCriteria{ReporterID: account.ID, Whitelisted: false})
	if err != nil {
		return blacklist, err
	}
	for _, each := range entries {
		blacklist = append(blacklist, GetAccountByID(each.BlockedUser.ID, accountRepo))
	}
	return blacklist, err
}
