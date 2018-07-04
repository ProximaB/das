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

const (
	// PartnershipRequestReceived labels the request as "received", once it's viewed by recipient
	PartnershipRequestReceived = 1
	// PartnershipRequestSent labels the request at "sent", once it's sent out by sender
	PartnershipRequestSent = 2
)

const (
	// PartnershipRoleLead is the reference value for the Lead role
	PartnershipRoleLead = "LEAD"
	// PartnershipRoleFollow is the reference value for the Follow role
	PartnershipRoleFollow = "FOLLOW"
)

// Partnership defines the combination of a lead and a follow. A partnership is uniquely identified
// if the lead and follow are confirmed.
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

// IPartnershipRepository defines the interface that a partnership repository should implement
type IPartnershipRepository interface {
	CreatePartnership(partnership *Partnership) error
	SearchPartnership(criteria SearchPartnershipCriteria) ([]Partnership, error)
	UpdatePartnership(partnership Partnership) error
	DeletePartnership(partnership Partnership) error
}

// SearchPartnershipCriteria provides the parameters that an IPartnershipRepository can use to search by
type SearchPartnershipCriteria struct {
	PartnershipID int `schema:"id"`
	LeadID        int `schema:"lead"`
	FollowID      int `schema:"follow"`
}

// GetAllPartnerships returns all the partnerships that caller account is in, including as a lead and as a follow
func (account Account) GetAllPartnerships(repo IPartnershipRepository) ([]Partnership, error) {
	asLeads, err := repo.SearchPartnership(SearchPartnershipCriteria{
		LeadID: account.ID,
	})
	if err != nil {
		return nil, err
	}

	asFollows, err := repo.SearchPartnership(SearchPartnershipCriteria{
		FollowID: account.ID,
	})
	if err != nil {
		return nil, err
	}

	allPartnerships := make([]Partnership, 0)
	for _, each := range asLeads {
		allPartnerships = append(allPartnerships, each)
	}
	for _, each := range asFollows {
		allPartnerships = append(allPartnerships, each)
	}
	return allPartnerships, err
}

// GetPartnershipByID retrieves the Partnership in the provided repository by the specified ID
func GetPartnershipByID(id int, repo IPartnershipRepository) (Partnership, error) {
	searchResults, err := repo.SearchPartnership(SearchPartnershipCriteria{PartnershipID: id})
	if err != nil || searchResults == nil || len(searchResults) != 1 {
		return Partnership{}, err
	}
	return searchResults[0], err
}

// MustGetPartnershipByID uses an known ID and a concrete PartnershipRepository to find the
// partnership by the ID provided. If such partnership is not found, system will panic.
func MustGetPartnershipByID(id int, repo IPartnershipRepository) Partnership {
	searchResults, err := repo.SearchPartnership(SearchPartnershipCriteria{PartnershipID: id})
	if err != nil {
		panic(err.Error())
	}
	if len(searchResults) != 1 {
		panic("cannot find partnership with this ID")
	}
	return searchResults[0]
}

// HasAthlete checks if the provided Athlete ID is in this partnership
func (partnership Partnership) HasAthlete(athleteID int) bool {
	return partnership.LeadID == athleteID || partnership.FollowID == athleteID
}
