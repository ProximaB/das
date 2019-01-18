// Dancesport Application System (DAS)
// Copyright (C) 2019 Yubing Hou
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
	"errors"
	"time"
)

const (
	COMPETITION_INVITATION_STATUS_ACCEPTED = "Accepted"
	COMPETITION_INVITATION_STATUS_REJECTED = "Rejected"
	COMPETITION_INVITATION_STATUS_PENDING  = "Pending"
	COMPETITION_INVITATION_STATUS_REVOKED  = "Revoked"
	COMPETITION_INVITATION_STATUS_EXPIRED  = "Expired"
)

type CompetitionOfficialInvitation struct {
	ID                 int
	Sender             Account
	Recipient          Account
	ServiceCompetition Competition // the competition that the recipient will serve at if accepted
	AssignedRoleID     int         // only allow Adjudicator, Scrutineer, Deck Captain, Emcee
	InvitationStatus   string
	ExpirationDate     time.Time
	CreateUserId       int
	DateTimeCreated    time.Time
	UpdateUserId       int
	DateTimeUpdated    time.Time
}

type SearchCompetitionOfficialInvitationCriteria struct {
	SenderID             int
	RecipientID          int
	ServiceCompetitionID int
	AssignedRoleID       int
	Status               string
	CreateUserID         int
	UpdateUserID         int
}

type ICompetitionOfficialInvitationRepository interface {
	CreateCompetitionOfficialInvitationRepository(invitation *CompetitionOfficialInvitation) error
	DeleteCompetitionOfficialInvitationRepository(invitation CompetitionOfficialInvitation) error
	SearchCompetitionOfficialInvitationRepository(criteria SearchCompetitionOfficialInvitationCriteria) ([]CompetitionOfficialInvitation, error)
	UpdateCompetitionOfficialInvitationRepository(invitation CompetitionOfficialInvitation) error
}

type CompetitionOfficialInvitationService struct {
	accountRepo     IAccountRepository
	competitionRepo ICompetitionRepository
	officialRepo    ICompetitionOfficialRepository
	invitationRepo  ICompetitionOfficialInvitationRepository
}

func NewCompetitionOfficialInvitationService(
	accountRepo IAccountRepository,
	competitionRepo ICompetitionRepository,
	officialRep ICompetitionOfficialRepository,
	invitationRepo ICompetitionOfficialInvitationRepository) CompetitionOfficialInvitationService {
	return CompetitionOfficialInvitationService{
		accountRepo:     accountRepo,
		competitionRepo: competitionRepo,
		officialRepo:    officialRep,
		invitationRepo:  invitationRepo,
	}
}

func (service CompetitionOfficialInvitationService) SearchCompetitionOfficialInvitation(criteria SearchCompetitionOfficialInvitationCriteria) ([]CompetitionOfficialInvitation, error) {
	return service.invitationRepo.SearchCompetitionOfficialInvitationRepository(criteria)
}

func (service CompetitionOfficialInvitationService) CreateCompetitionOfficialInvitation(sender, recipient Account, serviceRole int, comp Competition) (CompetitionOfficialInvitation, error) {
	invitation := CompetitionOfficialInvitation{}

	// sender must be the creator of the competition
	if sender.ID != comp.CreateUserID {
		return invitation, errors.New("Not authorized to send competition official invitation.")
	}

	invitation.Sender = sender

	// competition must be prior to running
	if comp.GetStatus() >= CompetitionStatusInProgress {
		return invitation, errors.New("Competition is already running and no more officials can be assigned.")
	}

	invitation.ServiceCompetition = comp

	// recipient must already have the request service role
	if !recipient.HasRole(serviceRole) {
		return invitation, errors.New("Recipient does not have this role provisioned by Administrator.")
	}

	invitation.Recipient = recipient
	invitation.AssignedRoleID = serviceRole
	invitation.DateTimeCreated = time.Now()
	invitation.CreateUserId = sender.ID
	invitation.DateTimeUpdated = time.Now()
	invitation.UpdateUserId = sender.ID

	// invitation will expire either:
	// - 30 days after the invitation, or
	// - after the competition
	thirtyDayLimit := time.Now().AddDate(0, 0, 30)
	if thirtyDayLimit.Before(comp.EndDateTime) {
		invitation.ExpirationDate = thirtyDayLimit
	} else {
		invitation.ExpirationDate = comp.EndDateTime
	}

	// initialize invitation status to pending
	invitation.InvitationStatus = COMPETITION_INVITATION_STATUS_PENDING

	// create the role invitation
	createErr := service.invitationRepo.CreateCompetitionOfficialInvitationRepository(&invitation)

	// TODO: send notification to recipient (requires notification)

	return invitation, createErr
}

func (service CompetitionOfficialInvitationService) UpdateCompetitionOfficialInvitation(currentUser Account, invitation CompetitionOfficialInvitation, response string) error {
	// only the sender and the recipient can make changes to the invitation
	if currentUser.ID != invitation.Sender.ID && currentUser.ID != invitation.Recipient.ID {
		return errors.New("Not authorized to make changes to this invitation")
	}

	// check terminal status
	if invitation.InvitationStatus == COMPETITION_INVITATION_STATUS_EXPIRED {
		return errors.New("Invitation is expired and can no longer be updated")
	}
	if invitation.InvitationStatus == COMPETITION_INVITATION_STATUS_REVOKED {
		return errors.New("Invitation is revoked and can no longer be updated")
	}
	if invitation.InvitationStatus == COMPETITION_INVITATION_STATUS_REJECTED {
		return errors.New("Invitation is rejected and can no longer be updated")
	}

	// for no-terminal status: pending and accepted
	// pending request can be updated by:
	// - recipient to accept/reject
	// - sender to revoke
	// accepted request can be updated by:
	// - recipient to reject
	// - sender to revoke
	canUpdate := false
	if invitation.InvitationStatus == COMPETITION_INVITATION_STATUS_PENDING {
		if currentUser.ID == invitation.Recipient.ID {
			// can accept or reject
			if response != COMPETITION_INVITATION_STATUS_ACCEPTED && response != COMPETITION_INVITATION_STATUS_REJECTED {
				return errors.New("The invitation can only be accepted or rejected.")
			} else {
				canUpdate = true
			}
		} else if currentUser.ID == invitation.Sender.ID {
			if response != COMPETITION_INVITATION_STATUS_REVOKED {
				return errors.New("The invitation can only be revoked")
			} else {
				canUpdate = true
			}
		}

	} else if invitation.InvitationStatus == COMPETITION_INVITATION_STATUS_ACCEPTED {
		if currentUser.ID == invitation.Recipient.ID {
			if response != COMPETITION_INVITATION_STATUS_REJECTED {
				return errors.New("The invitation can only be rejected.")
			}
		} else if currentUser.ID == invitation.Sender.ID {
			if response != COMPETITION_INVITATION_STATUS_REVOKED {
				return errors.New("The invitation can only be revoked")
			} else {
				canUpdate = true
			}
		}
	}

	if canUpdate {
		invitation.InvitationStatus = response
		invitation.DateTimeUpdated = time.Now()
		invitation.UpdateUserId = currentUser.ID
		return service.invitationRepo.UpdateCompetitionOfficialInvitationRepository(invitation)
	}
	return errors.New("An unknown error occurred while processing this invitation. Please report this incident to site administrator.")
}
