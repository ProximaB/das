// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

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
func (self Account) GetAllPartnerships(repo IPartnershipRepository) ([]Partnership, error) {
	asLeads, err := repo.SearchPartnership(SearchPartnershipCriteria{
		LeadID: self.ID,
	})
	if err != nil {
		return nil, err
	}

	asFollows, err := repo.SearchPartnership(SearchPartnershipCriteria{
		FollowID: self.ID,
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
