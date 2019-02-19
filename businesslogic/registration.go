package businesslogic

import (
	"errors"
	"log"
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

// EventRegistrationForm specifies the data needed to create/update/drop event registration
type EventRegistrationForm struct {
	CompetitionID      int
	PartnershipID      int
	EventsAdded        []int
	EventsDropped      []int
	CountryRepresented int
	StateRepresented   int
	SchoolRepresented  int
	StudioRepresented  int
}

// Validate performs sanity check of EventRegistrationForm
func (registration EventRegistrationForm) Validate() error {
	if registration.PartnershipID < 1 {
		return errors.New("partnership should be specified")
	}
	if registration.CompetitionID < 1 {
		return errors.New("competition should be specified")
	}
	return nil
}

// CompetitionRegistrationService provides a high level operation for Competition Registration
type CompetitionRegistrationService struct {
	AccountRepository               IAccountRepository
	PartnershipRepository           IPartnershipRepository
	CompetitionRepository           ICompetitionRepository
	EventRepository                 IEventRepository
	AthleteCompetitionEntryRepo     IAthleteCompetitionEntryRepository
	PartnershipCompetitionEntryRepo IPartnershipCompetitionEntryRepository
	AthleteEventEntryRepo           IAthleteEventEntryRepository
	PartnershipEventEntryRepo       IPartnershipEventEntryRepository
	AthleteCompetitionEntryService  AthleteCompetitionEntryService
}

func NewCompetitionRegistrationService(
	accountRepo IAccountRepository,
	partnershipRepo IPartnershipRepository,
	competitionRepo ICompetitionRepository,
	eventRepo IEventRepository,
	athleteCompetitionEntryRepo IAthleteCompetitionEntryRepository,
	athleteEventEntryRepo IAthleteEventEntryRepository) CompetitionRegistrationService {
	service := CompetitionRegistrationService{}
	service.AccountRepository = accountRepo
	service.PartnershipRepository = partnershipRepo
	service.CompetitionRepository = competitionRepo
	service.EventRepository = eventRepo
	service.AthleteCompetitionEntryRepo = athleteCompetitionEntryRepo
	service.AthleteCompetitionEntryService = NewAthleteCompetitionEntryService(accountRepo, competitionRepo, athleteCompetitionEntryRepo)
	return service
}

func (service CompetitionRegistrationService) UpdateRegistration(currentUser Account, form EventRegistrationForm) error {
	// data access control: current user must be one of the following:
	// - Athlete: competition is still in: Open Registration
	// - Scrutineer: competition is in progress
	// - Organizer: competition is in 1) closed registration, 2) in progress
	competitions, compErr := service.CompetitionRepository.SearchCompetition(SearchCompetitionCriteria{
		ID: form.CompetitionID,
	})
	if compErr != nil || len(competitions) != 1 {
		log.Printf(compErr.Error())
		return errors.New("Error in finding this competition")
	}
	competition := competitions[0]

	// TODO: implement role and ownership check. This will be dependent on the implementation of competition officials

	canChange := false
	if currentUser.HasRole(AccountTypeAthlete) && competition.GetStatus() == CompetitionStatusOpenRegistration {
		canChange = true
	}

	if !canChange {
		return errors.New("Registration can no longer be updated")
	}

	// for scrutineer and organizer, check if they are either invited officials of the competition
	// Scrutineer/Organizer of one competitions should not be able to update the registration forms of other competitions

	// check if current user already has registration

	// if has existing registration, get the existing registration
	//
	return errors.New("Not implementeds")
}

// ValidateEventRegistration validates if the registration data is valid. This does not create the registration
func (service CompetitionRegistrationService) ValidateEventRegistration(currentUser Account, registration EventRegistrationForm) error {
	if err := registration.Validate(); err != nil {
		return err
	}

	partnership, err := GetPartnershipByID(registration.PartnershipID, service.PartnershipRepository)
	if err != nil {
		return err
	}

	competition, err := GetCompetitionByID(registration.CompetitionID, service.CompetitionRepository)
	if err != nil {
		return nil
	}

	if currentUser.HasRole(AccountTypeAthlete) && competition.GetStatus() != CompetitionStatusOpenRegistration {
		return errors.New("registration is no longer open")
	}

	// check if organizer is authorized to change this partnership's registration
	organizer := GetAccountByID(competition.CreateUserID, service.AccountRepository) // creator may not be the organizer of specified competition
	if currentUser.HasRole(AccountTypeOrganizer) && organizer.ID != currentUser.ID {
		return errors.New("not an authorized organizer to update the registration")
	}

	// check if the creator of the entry is competitor
	if currentUser.HasRole(AccountTypeAthlete) && (!partnership.HasAthlete(currentUser.ID)) {
		// request was sent by people who are neither the lead or the follow of this partnership
		return errors.New("not an authorized athlete to update the registration")
	}

	// check if those events are valid and open for registration
	for _, each := range registration.EventsAdded {
		cbe, err := GetEventByID(each, service.EventRepository)
		if err != nil || cbe.ID == 0 {
			return errors.New("competitive ballroom event does not exist")
		}

		event, searchErr := service.EventRepository.SearchEvent(SearchEventCriteria{EventID: cbe.ID})
		if searchErr != nil || len(event) != 1 {
			return errors.New("event does not exist")
		} else if event[0].StatusID != EVENT_STATUS_OPEN {
			return errors.New("event is not open for registration")
		}
	}

	// create competition entry for the lead, if the entry has not been created yet
	entries, hasEntryErr := service.AthleteCompetitionEntryRepo.SearchEntry(SearchAthleteCompetitionEntryCriteria{
		CompetitionID: registration.CompetitionID,
		AthleteID:     partnership.Lead.ID,
	})
	if len(entries) != 1 || hasEntryErr != nil {
		service.AthleteCompetitionEntryRepo.CreateEntry(&AthleteCompetitionEntry{
			CompetitionEntry: BaseCompetitionEntry{
				CompetitionID: registration.CompetitionID,
			},
			//AthleteID:     partnership.LeadID,
		})
	}

	// create competition entry for the follow, if the entry has not been created yet
	entries, hasEntryErr = service.AthleteCompetitionEntryRepo.SearchEntry(SearchAthleteCompetitionEntryCriteria{
		CompetitionID: registration.CompetitionID,
		AthleteID:     partnership.Follow.ID,
	})
	if len(entries) != 1 || hasEntryErr != nil {
		service.AthleteCompetitionEntryRepo.CreateEntry(&AthleteCompetitionEntry{
			CompetitionEntry: BaseCompetitionEntry{
				CompetitionID: registration.CompetitionID,
			},
			//AthleteID:     partnership.FollowID,
		})
	}

	// check if added events are open
	for _, each := range registration.EventsAdded {
		cbe, findErr := GetEventByID(each, service.EventRepository)
		if findErr != nil {
			return errors.New("a competitive ballroom event does not exist")
		}
		events, eventErr := service.EventRepository.SearchEvent(SearchEventCriteria{
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
				CompetitorTag:   0,
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

// CreateEntry takes the current user and the registration data and create new Competition Entry for
// each of the athlete
func (service CompetitionRegistrationService) CreateAthleteCompetitionEntry(currentUser Account, registration EventRegistrationForm) error {
	partnership, findPartnershipErr := GetPartnershipByID(registration.PartnershipID, service.PartnershipRepository)
	if findPartnershipErr != nil {
		return findPartnershipErr
	}
	leadCompEntry := AthleteCompetitionEntry{
		CompetitionEntry: BaseCompetitionEntry{
			CompetitionID:    registration.CompetitionID,
			CheckInIndicator: false,
			CreateUserID:     currentUser.ID,
			DateTimeCreated:  time.Now(),
			UpdateUserID:     currentUser.ID,
			DateTimeUpdated:  time.Now(),
		},
		AthleteID:                partnership.Lead.ID,
		PaymentReceivedIndicator: false,
	}
	followCompEntry := AthleteCompetitionEntry{
		CompetitionEntry: BaseCompetitionEntry{
			CompetitionID:    registration.CompetitionID,
			CheckInIndicator: false,
			CreateUserID:     currentUser.ID,
			DateTimeCreated:  time.Now(),
			UpdateUserID:     currentUser.ID,
			DateTimeUpdated:  time.Now(),
		},
		AthleteID:                partnership.Follow.ID,
		PaymentReceivedIndicator: false,
	}

	service.AthleteCompetitionEntryService.CreateAthleteCompetitionEntry(&leadCompEntry)
	service.AthleteCompetitionEntryService.CreateAthleteCompetitionEntry(&followCompEntry)
	return nil
}

// CreateEntry takes the current user and registration data and create a Competition Entry for
// this Partnership
func (service CompetitionRegistrationService) CreatePartnershipCompetitionEntry(currentUser Account, registration EventRegistrationForm) error {
	partnership, findPartnershipErr := GetPartnershipByID(registration.PartnershipID, service.PartnershipRepository)
	if findPartnershipErr != nil {
		return findPartnershipErr
	}

	partnershipEntry := PartnershipCompetitionEntry{
		PartnershipID: partnership.ID,
		CompetitionEntry: BaseCompetitionEntry{
			CompetitionID:    registration.CompetitionID,
			CheckInIndicator: false,
			CreateUserID:     currentUser.ID,
			DateTimeCreated:  time.Now(),
			UpdateUserID:     currentUser.ID,
			DateTimeUpdated:  time.Now(),
		},
	}
	partnershipEntry.createPartnershipCompetitionEntry(service.CompetitionRepository, service.PartnershipCompetitionEntryRepo)
	return nil
}

// CreatePartnershipEventEntries takes the current user and registration data to create a new Event Entry for this partnership
func (service CompetitionRegistrationService) CreatePartnershipEventEntries(currentUser Account, registration EventRegistrationForm) error {
	for _, each := range registration.EventsAdded {
		eventEntry := PartnershipEventEntry{

			PartnershipID: registration.PartnershipID,
			EventEntry: EventEntry{
				EventID:         each,
				CompetitorTag:   0,
				CreateUserID:    currentUser.ID,
				DateTimeCreated: time.Now(),
				UpdateUserID:    currentUser.ID,
				DateTimeUpdated: time.Now(),
			},
		}
		createErr := service.PartnershipEventEntryRepo.CreatePartnershipEventEntry(&eventEntry)
		if createErr != nil {
			return createErr
		}
	}
	return nil
}

// DropPartnershipCompetitionEntry removes the competition entry of the specified partnership from the provided competition
// if that partnership, competition, or entry does not exist, and error will be thrown
func (service CompetitionRegistrationService) DropPartnershipCompetitionEntry(partnershipID, competitionID int) error {
	if results, err := service.PartnershipCompetitionEntryRepo.SearchEntry(SearchPartnershipCompetitionEntryCriteria{
		PartnershipID: partnershipID,
		CompetitionID: competitionID,
	}); err != nil {
		log.Printf("[error] cannot find competition entry for partnership ID = %d and competition ID = %d: %v", partnershipID, competitionID, err)
		return errors.New("an error occurred while searching for partnership competition entry")
	} else if len(results) != 1 {
		return errors.New("cannot find competition entry for this partnership")
	} else {
		return service.PartnershipCompetitionEntryRepo.DeleteEntry(results[0])
	}
}

// DropPartnershipEventEntries takes the current user and registration data and removes specified entries from the event
func (service CompetitionRegistrationService) DropPartnershipEventEntries(currentUser Account, registration EventRegistrationForm) error {
	for _, each := range registration.EventsDropped {
		eventEntry := PartnershipEventEntry{
			EventEntry: EventEntry{
				EventID:      each,
				UpdateUserID: currentUser.ID,
			},
			PartnershipID: registration.PartnershipID,
		}
		dropErr := service.PartnershipEventEntryRepo.DeletePartnershipEventEntry(eventEntry)
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
func GetEventRegistration(competitionID int, partnershipID int, user *Account, partnershipRepo IPartnershipRepository) (EventRegistrationForm, error) {
	// check if user is part of the partnership
	results, err := partnershipRepo.SearchPartnership(SearchPartnershipCriteria{PartnershipID: partnershipID})
	if err != nil {
		return EventRegistrationForm{}, errors.New("cannot find requested partnership")
	}
	if results == nil || len(results) != 1 {
		return EventRegistrationForm{}, errors.New("cannot find partnership for registration")
	}
	partnership := results[0]
	if user.ID == 0 || user.HasRole(AccountTypeAthlete) || (user.ID != partnership.Lead.ID && user.ID != partnership.Follow.ID) {
		return EventRegistrationForm{}, errors.New("not authorized to request this information")
	}

	//return dataaccess.GetCompetitiveBallroomEventRegistration(dataaccess.DATABASE, competitionID, partnershipID)
	return EventRegistrationForm{}, errors.New("not implemented")
}
