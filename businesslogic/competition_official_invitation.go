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
	COMPETITION_INVITATION_STATUS_Revoked  = "Revoked"
)

type CompetitionOfficialInvitation struct {
	ID                 int
	Sender             Account
	Recipient          Account
	ServiceCompetition Competition // the competition that the recipient will serve at if accepted
	AssignedRoleID     int         // only allow Adjudicator, Scrutineer, Deck Captain, Emcee
	InvitationStatus   string
	CreateUserId       int
	DateTimeCreated    time.Time
	UpdateUserId       int
	DateTimeUpdated    time.Time
}

type SearchCompetitionOfficialInvitationCriteria struct {
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

func (service CompetitionOfficialInvitationService) NewCompetitionOfficialInvitation(sender, recipient Account, serviceRole int, comp Competition) (CompetitionOfficialInvitation, error) {
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

	// create the role
	createErr := service.invitationRepo.CreateCompetitionOfficialInvitationRepository(&invitation)

	// TODO: send notification to recipient (requires notification)

	return invitation, createErr
}
