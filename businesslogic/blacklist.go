// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
func (self Account) GetBlacklistedAccounts(accountRepo IAccountRepository, blacklistRepo IPartnershipRequestBlacklistRepository) ([]Account, error) {
	blacklist := make([]Account, 0)
	entries, err := blacklistRepo.SearchPartnershipRequestBlacklist(SearchPartnershipRequestBlacklistCriteria{ReporterID: self.ID, Whitelisted: false})
	if err != nil {
		return blacklist, err
	}
	for _, each := range entries {
		blacklist = append(blacklist, GetAccountByID(each.BlockedUser.ID, accountRepo))
	}
	return blacklist, err
}
