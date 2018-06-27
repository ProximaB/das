// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

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

type Registration interface {
}

func ValidateCompetitiveBallroomEventRegistration(creator *Account,
	registration *EventRegistration,
	competitionRepo ICompetitionRepository,
	eventRepo IEventRepository,
	repo IAthleteCompetitionEntryRepository,
	accountRepo IAccountRepository,
	partnershipRepo IPartnershipRepository) error {
	// check if partnership exists
	results, partnershipErr := partnershipRepo.SearchPartnership(SearchPartnershipCriteria{PartnershipID: registration.PartnershipID})
	if results == nil || partnershipErr != nil {
		return errors.New("partnership does not exist")
	}
	if len(results) != 1 {
		return errors.New("cannot find partnership")
	}
	partnership := results[0]

	// check if competition exists
	competitions, _ := competitionRepo.SearchCompetition(SearchCompetitionCriteria{ID: registration.CompetitionID})
	if len(competitions) != 1 {
		return errors.New("competition does not exists")
	}

	// check if competition still allow registration
	// for competitor: only change registration if registration is open
	// for organizer, only change registration if competition is
	if creator.AccountTypeID == ACCOUNT_TYPE_ATHLETE && competitions[0].GetStatus() != COMPETITION_STATUS_OPEN_REGISTRATION {
		return errors.New("registration is no longer open")
	}

	// check if organizer is authorized to change this partnership's reigstration
	organizer := GetAccountByID(competitions[0].CreateUserID, accountRepo) // creator may not be the organizer of specified competition
	if creator.AccountTypeID == ACCOUNT_TYPE_ORGANIZER && organizer.ID != creator.ID {
		return errors.New("not the organizer of specified competition")
	}

	// check if the creator of the entry is competitor
	if creator.AccountTypeID == ACCOUNT_TYPE_ATHLETE && partnership.LeadID != creator.ID && partnership.FollowID != creator.ID {
		// request was sent by people who are neither the lead or the follow of this partnership
		return errors.New("not the lead or the follow of specified partnership")
	}

	// check if those events are valid and open for registration
	for _, each := range registration.EventsAdded {
		cbe, err := GetEventByID(each, eventRepo)
		if err != nil || cbe.ID == 0 {
			return errors.New("competitive ballroom event does not exist")
		}

		event, searchErr := eventRepo.SearchEvent(SearchEventCriteria{EventID: cbe.ID})
		if searchErr != nil || len(event) != 1 {
			return errors.New("event does not exist")
		} else if event[0].StatusID != EVENT_STATUS_OPEN {
			return errors.New("event is not open for registration")
		}
	}

	// create competition entry for the lead, if the entry has not been created yet
	entries, hasEntryErr := repo.SearchAthleteCompetitionEntry(SearchAthleteCompetitionEntryCriteria{
		CompetitionID: registration.CompetitionID,
		AthleteID:     partnership.LeadID,
	})
	if len(entries) != 1 || hasEntryErr != nil {
		repo.CreateAthleteCompetitionEntry(&AthleteCompetitionEntry{
			CompetitionEntry: CompetitionEntry{
				CompetitionID: registration.CompetitionID,
			},
			//AthleteID:     partnership.LeadID,
		})
	}

	// create competition entry for the follow, if the entry has not been created yet
	entries, hasEntryErr = repo.SearchAthleteCompetitionEntry(SearchAthleteCompetitionEntryCriteria{
		CompetitionID: registration.CompetitionID,
		AthleteID:     partnership.FollowID,
	})
	if len(entries) != 1 || hasEntryErr != nil {
		repo.CreateAthleteCompetitionEntry(&AthleteCompetitionEntry{
			CompetitionEntry: CompetitionEntry{
				CompetitionID: registration.CompetitionID,
			},
			//AthleteID:     partnership.FollowID,
		})
	}

	// check if added events are open
	for _, each := range registration.EventsAdded {
		cbe, findErr := GetEventByID(each, eventRepo)
		if findErr != nil {
			return errors.New("a competitive ballroom event does not exist")
		}
		events, eventErr := eventRepo.SearchEvent(SearchEventCriteria{
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
		eventEntry := EventEntry{
			EventID:         each,
			PartnershipID:   registration.PartnershipID,
			CompetitorTag:   0,
			CreateUserID:    creator.ID,
			DateTimeCreated: time.Now(),
			UpdateUserID:    creator.ID,
			DateTimeUpdated: time.Now(),
		}
		eligibilityErr := CheckCompetitiveBallroomEventEligibility(eventEntry)
		if eligibilityErr != nil {
			return eligibilityErr
		}
	}
	return nil
}

func CreateEventEntries(creator *Account,
	registration *EventRegistration,
	eventEntryRepo IPartnershipEventEntryRepository) error {
	for _, each := range registration.EventsAdded {
		eventEntry := PartnershipEventEntry{
			EventEntry: EventEntry{
				EventID:         each,
				PartnershipID:   registration.PartnershipID,
				CompetitorTag:   0,
				CreateUserID:    creator.ID,
				DateTimeCreated: time.Now(),
				UpdateUserID:    creator.ID,
				DateTimeUpdated: time.Now(),
			},
		}
		createErr := eventEntryRepo.CreatePartnershipEventEntry(&eventEntry)
		if createErr != nil {
			return createErr
		}
	}
	return nil
}

func DropEventEntries(creator *Account,
	registration *EventRegistration,
	eventEntryRepo IPartnershipEventEntryRepository) error {
	for _, each := range registration.EventsDropped {
		eventEntry := PartnershipEventEntry{
			EventEntry: EventEntry{
				EventID:      each,
				UpdateUserID: creator.ID,
			},
			PartnershipID: registration.PartnershipID,
		}
		dropErr := eventEntryRepo.DeletePartnershipEventEntry(eventEntry)
		if dropErr != nil {
			return dropErr
		}
	}
	return nil
}

func CheckCompetitiveBallroomEventEligibility(entry EventEntry) error {
	// TODO: implement this method
	return errors.New("not implemented")
}

func GetEventRegistration(competitionID int,
	partnershipID int,
	user *Account,
	partnershipRepo IPartnershipRepository,
) (Registration, error) {
	// check if user is part of the partnership
	results, err := partnershipRepo.SearchPartnership(SearchPartnershipCriteria{PartnershipID: partnershipID})
	if err != nil {
		return EventRegistration{}, errors.New("cannot find requested partnership")
	}
	if results == nil || len(results) != 1 {
		return nil, errors.New("cannot find partnership for registration")
	}
	partnership := results[0]
	if user.ID == 0 || user.AccountTypeID != ACCOUNT_TYPE_ATHLETE || (user.ID != partnership.LeadID && user.ID != partnership.FollowID) {
		return EventRegistration{}, errors.New("not authorized to request this information")
	}

	//return dataaccess.GetCompetitiveBallroomEventRegistration(dataaccess.DATABASE, competitionID, partnershipID)
	return nil, errors.New("not implemented")
}
