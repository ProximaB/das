package viewmodel

import (
	"github.com/yubing24/das/businesslogic"
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
		ID:         partnership.PartnershipID,
		LeadName:   partnership.Lead.GetName(),
		FollowName: partnership.Follow.GetName(),
		Since:      partnership.DateTimeCreated,
		SameSexIND: partnership.SameSex,
		Favorite:   partnership.FavoriteLead,
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
