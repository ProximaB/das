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
	"errors"
	"time"
)

// PartnershipCompetitionRepresentation specifies the Country, State, Studio, School that a partnership represents
// at a Competition
type PartnershipCompetitionRepresentation struct {
	ID                            int
	PartnershipCompetitionEntryID int
	CountryID                     *int
	StateID                       *int
	StudioID                      *int
	SchoolID                      *int
	CreateUserID                  int
	DateTimeCreated               time.Time
	UpdateUserID                  int
	DateTimeUpdated               time.Time
}

// EventRegistration specifies the data needed to create/update/drop event registration
type EventRegistration struct {
	CompetitionID      int   `json:"competition"`
	PartnershipID      int   `json:"partnership"`
	EventsAdded        []int `json:"added"`   // event id, should not be competitive ballroom event id
	EventsDropped      []int `json:"dropped"` // event id, should not be competitive ballroom event id
	CountryRepresented int   `json:"country"`
	StateRepresented   int   `json:"state"`
	SchoolRepresented  int   `json:"school"`
	StudioRepresented  int   `json:"studio"`
}

// CompetitionRegistrationService provides a high level operation for Competition Registration
type CompetitionRegistrationService struct {
	IAccountRepository
	IPartnershipRepository
	ICompetitionRepository
	IEventRepository
	IAthleteCompetitionEntryRepository
	IPartnershipCompetitionEntryRepository
	IPartnershipEventEntryRepository
}

// ValidateEventRegistration validates if the registration data is valid. This does not create the registration
func (service CompetitionRegistrationService) ValidateEventRegistration(currentUser Account, registration EventRegistration) error {
	partnership, err := GetPartnershipByID(registration.PartnershipID, service.IPartnershipRepository)
	if err != nil {
		return err
	}

	competition, err := GetCompetitionByID(registration.CompetitionID, service.ICompetitionRepository)
	if err != nil {
		return nil
	}

	if currentUser.AccountTypeID == AccountTypeAthlete && competition.GetStatus() != CompetitionStatusOpenRegistration {
		return errors.New("registration is no longer open")
	}

	// check if organizer is authorized to change this partnership's registration
	organizer := GetAccountByID(competition.CreateUserID, service.IAccountRepository) // creator may not be the organizer of specified competition
	if currentUser.AccountTypeID == AccountTypeOrganizer && organizer.ID != currentUser.ID {
		return errors.New("not an authorized organizer to update the registration")
	}

	// check if the creator of the entry is competitor
	if currentUser.AccountTypeID == AccountTypeAthlete && (!partnership.HasAthlete(currentUser.ID)) {
		// request was sent by people who are neither the lead or the follow of this partnership
		return errors.New("not an authorized athlete to update the registration")
	}

	// check if those events are valid and open for registration
	for _, each := range registration.EventsAdded {
		cbe, err := GetEventByID(each, service.IEventRepository)
		if err != nil || cbe.ID == 0 {
			return errors.New("competitive ballroom event does not exist")
		}

		event, searchErr := service.IEventRepository.SearchEvent(SearchEventCriteria{EventID: cbe.ID})
		if searchErr != nil || len(event) != 1 {
			return errors.New("event does not exist")
		} else if event[0].StatusID != EVENT_STATUS_OPEN {
			return errors.New("event is not open for registration")
		}
	}

	// create competition entry for the lead, if the entry has not been created yet
	entries, hasEntryErr := service.IAthleteCompetitionEntryRepository.SearchAthleteCompetitionEntry(SearchAthleteCompetitionEntryCriteria{
		CompetitionID: registration.CompetitionID,
		AthleteID:     partnership.LeadID,
	})
	if len(entries) != 1 || hasEntryErr != nil {
		service.IAthleteCompetitionEntryRepository.CreateAthleteCompetitionEntry(&AthleteCompetitionEntry{
			CompetitionEntry: CompetitionEntry{
				CompetitionID: registration.CompetitionID,
			},
			//AthleteID:     partnership.LeadID,
		})
	}

	// create competition entry for the follow, if the entry has not been created yet
	entries, hasEntryErr = service.IAthleteCompetitionEntryRepository.SearchAthleteCompetitionEntry(SearchAthleteCompetitionEntryCriteria{
		CompetitionID: registration.CompetitionID,
		AthleteID:     partnership.FollowID,
	})
	if len(entries) != 1 || hasEntryErr != nil {
		service.IAthleteCompetitionEntryRepository.CreateAthleteCompetitionEntry(&AthleteCompetitionEntry{
			CompetitionEntry: CompetitionEntry{
				CompetitionID: registration.CompetitionID,
			},
			//AthleteID:     partnership.FollowID,
		})
	}

	// check if added events are open
	for _, each := range registration.EventsAdded {
		cbe, findErr := GetEventByID(each, service.IEventRepository)
		if findErr != nil {
			return errors.New("a competitive ballroom event does not exist")
		}
		events, eventErr := service.IEventRepository.SearchEvent(SearchEventCriteria{
			EventID: cbe.ID,
		})
		if eventErr != nil || len(events) != 1 {
			return errors.New("a competitive ballroom event is invalid")
		}
		if events[0].StatusID != EVENT_STATUS_OPEN {
			return errors.New("event is no longer open for registration")
		}
	}

	// check if dropped events are open

	// check if added events are already added

	// check if dropped events are added

	// check event entries, and see if this partnership is still eligible for entering these events
	// TODO: eligibility check
	for _, each := range registration.EventsAdded {
		eventEntry := PartnershipEventEntry{
			PartnershipID: registration.PartnershipID,
			EventEntry: EventEntry{
				EventID:         each,
				Mask:            0,
				CreateUserID:    currentUser.ID,
				DateTimeCreated: time.Now(),
				UpdateUserID:    currentUser.ID,
				DateTimeUpdated: time.Now(),
			},
		}
		eligibilityErr := checkEventEligibility(eventEntry)
		if eligibilityErr != nil {
			return eligibilityErr
		}
	}
	return nil
}

// CreateAthleteCompetitionEntry takes the current user and the registration data and create new Competition Entry for
// each of the athlete
func (service CompetitionRegistrationService) CreateAthleteCompetitionEventEntry(currentUser Account, registration EventRegistration) error {
	partnership, findPartnershipErr := GetPartnershipByID(registration.PartnershipID, service.IPartnershipRepository)
	if findPartnershipErr != nil {
		return findPartnershipErr
	}
	leadCompEntry := AthleteCompetitionEntry{
		CompetitionEntry: CompetitionEntry{
			CompetitionID:    registration.CompetitionID,
			CheckInIndicator: false,
			CreateUserID:     currentUser.ID,
			DateTimeCreated:  time.Now(),
			UpdateUserID:     currentUser.ID,
			DateTimeUpdated:  time.Now(),
		},
		AthleteID:                partnership.LeadID,
		PaymentReceivedIndicator: false,
	}
	followCompEntry := AthleteCompetitionEntry{
		CompetitionEntry: CompetitionEntry{
			CompetitionID:    registration.CompetitionID,
			CheckInIndicator: false,
			CreateUserID:     currentUser.ID,
			DateTimeCreated:  time.Now(),
			UpdateUserID:     currentUser.ID,
			DateTimeUpdated:  time.Now(),
		},
		AthleteID:                partnership.FollowID,
		PaymentReceivedIndicator: false,
	}

	leadCompEntry.createAthleteCompetitionEntry(service.ICompetitionRepository, service.IAthleteCompetitionEntryRepository)
	followCompEntry.createAthleteCompetitionEntry(service.ICompetitionRepository, service.IAthleteCompetitionEntryRepository)
	return nil
}

// CreatePartnershipCompetitionEntry takes the current user and registration data and create a Competition Entry for
// this Partnership
func (service CompetitionRegistrationService) CreatePartnershipCompetitionEntry(currentUser Account, registration EventRegistration) error {
	partnership, findPartnershipErr := GetPartnershipByID(registration.PartnershipID, service.IPartnershipRepository)
	if findPartnershipErr != nil {
		return findPartnershipErr
	}

	partnershipEntry := PartnershipCompetitionEntry{
		PartnershipID: partnership.PartnershipID,
		CompetitionEntry: CompetitionEntry{
			CompetitionID:    registration.CompetitionID,
			CheckInIndicator: false,
			CreateUserID:     currentUser.ID,
			DateTimeCreated:  time.Now(),
			UpdateUserID:     currentUser.ID,
			DateTimeUpdated:  time.Now(),
		},
	}
	partnershipEntry.createPartnershipCompetitionEntry(service.ICompetitionRepository, service.IPartnershipCompetitionEntryRepository)
	return nil
}

// CreatePartnershipEventEntries takes the current user and registration data to create a new Event Entry for this partnership
func (service CompetitionRegistrationService) CreatePartnershipEventEntries(currentUser Account, registration EventRegistration) error {
	for _, each := range registration.EventsAdded {
		eventEntry := PartnershipEventEntry{

			PartnershipID: registration.PartnershipID,
			EventEntry: EventEntry{
				EventID:         each,
				Mask:            0,
				CreateUserID:    currentUser.ID,
				DateTimeCreated: time.Now(),
				UpdateUserID:    currentUser.ID,
				DateTimeUpdated: time.Now(),
			},
		}
		createErr := service.IPartnershipEventEntryRepository.CreatePartnershipEventEntry(&eventEntry)
		if createErr != nil {
			return createErr
		}
	}
	return nil
}

// DropPartnershipEventEntries takes the current user and registration data and removes specified entries from the event
func (service CompetitionRegistrationService) DropPartnershipEventEntries(currentUser Account, registration EventRegistration) error {
	for _, each := range registration.EventsDropped {
		eventEntry := PartnershipEventEntry{
			EventEntry: EventEntry{
				EventID:      each,
				UpdateUserID: currentUser.ID,
			},
			PartnershipID: registration.PartnershipID,
		}
		dropErr := service.IPartnershipEventEntryRepository.DeletePartnershipEventEntry(eventEntry)
		if dropErr != nil {
			return dropErr
		}
	}
	return nil
}

func checkEventEligibility(entry PartnershipEventEntry) error {
	// TODO: implement this method
	return errors.New("not implemented")
}

// GetEventRegistration get event registration for the provided competition and partnership
func GetEventRegistration(competitionID int, partnershipID int, user *Account, partnershipRepo IPartnershipRepository) (EventRegistration, error) {
	// check if user is part of the partnership
	results, err := partnershipRepo.SearchPartnership(SearchPartnershipCriteria{PartnershipID: partnershipID})
	if err != nil {
		return EventRegistration{}, errors.New("cannot find requested partnership")
	}
	if results == nil || len(results) != 1 {
		return EventRegistration{}, errors.New("cannot find partnership for registration")
	}
	partnership := results[0]
	if user.ID == 0 || user.AccountTypeID != AccountTypeAthlete || (user.ID != partnership.LeadID && user.ID != partnership.FollowID) {
		return EventRegistration{}, errors.New("not authorized to request this information")
	}

	//return dataaccess.GetCompetitiveBallroomEventRegistration(dataaccess.DATABASE, competitionID, partnershipID)
	return EventRegistration{}, errors.New("not implemented")
}
