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

package viewmodel

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"time"
)

type Partnership struct {
	ID         int       `json:"id"`
	LeadName   string    `json:"lead"`
	FollowName string    `json:"follow"`
	Since      time.Time `json:"since"`
	SameSexIND bool      `json:"samesex"`
	Favorite   bool      `json:"favorite"`
}

func PartnershipDataModelToViewModel(partnership businesslogic.Partnership) Partnership {
	return Partnership{
		ID:         partnership.ID,
		LeadName:   partnership.Lead.FullName(),
		FollowName: partnership.Follow.FullName(),
		Since:      partnership.DateTimeCreated,
		SameSexIND: partnership.SameSex,
		Favorite:   partnership.FavoriteByLead,
	}
}

type SearchPartnershipRequestViewModel struct {
	RequestID       int    `schema:"id"`
	Type            int    `schema:"type"`
	Sender          string `schema:"sender"`
	Recipient       string `schema:"recipient"`
	RequestStatusID int    `schema:"status"`
}

type PartnershipRequestResponse struct {
	RequestID int `json:"request"`
	Response  int `json:"response"`
}

type CreatePartnershipRequest struct {
	SenderID       int    `json:"sender"`
	RecipientEmail string `json:"recipient"`
	RecipientRole  string `json:"role"`
	Message        string `json:"message"`
}

type PartnershipRequest struct {
	ID              int       `json:"id"`
	Sender          string    `json:"sender"`
	Recipient       string    `json:"recipient"`
	Message         string    `json:"message"`
	Status          int       `json:"status"`
	Role            string    `json:"role"`
	DateTimeCreated time.Time `json:"sent"`
}

type PartnershipRequestStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PartnershipRole struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func PartnershipRoleDataModelToViewModel(dataModel businesslogic.PartnershipRole) PartnershipRole {
	return PartnershipRole{
		ID:   dataModel.ID,
		Name: dataModel.Name,
	}
}
