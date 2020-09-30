package viewmodel

import (
	"github.com/ProximaB/das/businesslogic"
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

type PartnershipTinyViewModel struct {
	ID     int    `json:"id"`
	Lead   string `json:"lead"`
	Follow string `json:"follow"`
}

func PartnershipToTinyViewModel(partnership businesslogic.Partnership) PartnershipTinyViewModel {
	return PartnershipTinyViewModel{
		ID:     partnership.ID,
		Lead:   partnership.Lead.FullName(),
		Follow: partnership.Follow.FullName(),
	}
}

func PartnershipDataModelToViewModel(currentUser businesslogic.Account, partnership businesslogic.Partnership) Partnership {
	dto := Partnership{
		ID:         partnership.ID,
		LeadName:   partnership.Lead.FullName(),
		FollowName: partnership.Follow.FullName(),
		Since:      partnership.DateTimeCreated,
		SameSexIND: partnership.SameSex,
	}

	if currentUser.ID == partnership.Lead.ID {
		dto.Favorite = partnership.FavoriteByLead
	}
	if currentUser.ID == partnership.Follow.ID {
		dto.Favorite = partnership.FavoriteByFollow
	}

	return dto
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

// CreatePartnershipRequest defines the JSON payload structure for creating a partnership request.
type CreatePartnershipRequest struct {
	SenderID       int    `json:"sender" validate:"min=1"`
	RecipientEmail string `json:"recipient" validate:"min=5"`
	RecipientRole  int    `json:"roleId" validate:"min=1,max=2"`
	Message        string `json:"message"`
}

// PartnershipRequest view model
type PartnershipRequest struct {
	ID              int       `json:"id"`
	Sender          string    `json:"sender"`
	Recipient       string    `json:"recipient"`
	Message         string    `json:"message"`
	Status          int       `json:"statusId"`
	Role            string    `json:"role"`
	DateTimeCreated time.Time `json:"sent"`
}

// PartnershipRequestStatus defines the JSON structure of reference data PartnershipRequestStatus.
type PartnershipRequestStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// PartnershipRole defines the JSON structure reference data PartnershipRole
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
