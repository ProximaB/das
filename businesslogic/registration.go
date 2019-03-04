package businesslogic

import (
	"errors"
	"fmt"
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
	Competition        Competition
	Couple             Partnership
	EventsAdded        []Event
	EventsDropped      []Event
	CountryRepresented Country
	StateRepresented   State
	SchoolRepresented  School
	StudioRepresented  Studio
}

// CompetitionRegistrationService provides a high level operation for Competition Registration
type CompetitionRegistrationService struct {
	AccountRepository                  IAccountRepository
	PartnershipRepository              IPartnershipRepository
	CompetitionRepository              ICompetitionRepository
	EventRepository                    IEventRepository
	AthleteCompetitionEntryRepo        IAthleteCompetitionEntryRepository
	PartnershipCompetitionEntryRepo    IPartnershipCompetitionEntryRepository
	athleteEventEntryRepo              IAthleteEventEntryRepository
	PartnershipEventEntryRepo          IPartnershipEventEntryRepository
	AthleteCompetitionEntryService     AthleteCompetitionEntryService
	partnershipCompetitionEntryService PartnershipCompetitionEntryService
	athleteEventEntryService           AthleteEventEntryService
	coupleEventEntryService            PartnershipEventEntryService
}

func NewCompetitionRegistrationService(
	accountRepo IAccountRepository,
	partnershipRepo IPartnershipRepository,
	competitionRepo ICompetitionRepository,
	eventRepo IEventRepository,
	athleteCompetitionEntryRepo IAthleteCompetitionEntryRepository,
	athleteEventEntryRepo IAthleteEventEntryRepository,
	coupleCompetitionEntryRepo IPartnershipCompetitionEntryRepository,
	coupleEventEntryRepo IPartnershipEventEntryRepository) CompetitionRegistrationService {
	service := CompetitionRegistrationService{}
	service.AccountRepository = accountRepo
	service.PartnershipRepository = partnershipRepo
	service.CompetitionRepository = competitionRepo
	service.EventRepository = eventRepo
	service.athleteEventEntryRepo = athleteEventEntryRepo
	service.AthleteCompetitionEntryRepo = athleteCompetitionEntryRepo
	service.athleteEventEntryRepo = athleteEventEntryRepo
	service.PartnershipCompetitionEntryRepo = coupleCompetitionEntryRepo
	service.PartnershipEventEntryRepo = coupleEventEntryRepo
	service.AthleteCompetitionEntryService = NewAthleteCompetitionEntryService(accountRepo, competitionRepo, athleteCompetitionEntryRepo)
	return service
}

// UpdateRegistration is the only method that handles registration. This methods handles the following situations
// - Create new entries, if entries are not created
// - Delete entries, if exists
func (service CompetitionRegistrationService) UpdateRegistration(currentUser Account, registration EventRegistrationForm) error {
	// data access control: current user must be one of the following:
	// - Athlete: competition is still in: Open Registration
	// - Scrutineer: competition is in progress
	// - Organizer: competition is in 1) closed registration, 2) in progress

	// TODO: implement role and ownership check. This will be dependent on the implementation of competition officials

	canChange := false
	if currentUser.HasRole(AccountTypeAthlete) && registration.Competition.GetStatus() == CompetitionStatusOpenRegistration {
		canChange = true
	}

	if !canChange {
		return errors.New("registration can no longer be updated or you are not authorized")
	}

	var err error

	// Get current PartnershipCompetitionEntry, if not exists, then create a new one
	currentCoupleCompEntries, err := service.PartnershipCompetitionEntryRepo.SearchEntry(SearchPartnershipCompetitionEntryCriteria{
		PartnershipID: registration.Couple.ID,
		CompetitionID: registration.Competition.ID,
	})
	if err != nil {
		log.Printf("[error] searching competition entry of partnership ID = %v: %v", registration.Couple.ID, err)
	}
	if len(currentCoupleCompEntries) == 0 {
		entry := PartnershipCompetitionEntry{
			Competition:     registration.Competition,
			Couple:          registration.Couple,
			CreateUserID:    currentUser.ID,
			DateTimeCreated: time.Now(),
			UpdateUserID:    currentUser.ID,
			DateTimeUpdated: time.Now(),
		}
		if err = service.PartnershipCompetitionEntryRepo.CreateEntry(&entry); err != nil {
			return err
		}
	}

	currentLeadCompEntries, err := service.AthleteCompetitionEntryRepo.SearchEntry(SearchAthleteCompetitionEntryCriteria{CompetitionID: registration.Competition.ID, AthleteID: registration.Couple.Lead.ID})
	if err != nil {
		log.Printf("[error] searching athlete competition entry of lead %v: %v", registration.Couple.Lead.FullName(), err)
		return err
	}
	if len(currentLeadCompEntries) == 0 {
		entry := AthleteCompetitionEntry{
			Competition:              registration.Competition,
			Athlete:                  registration.Couple.Lead,
			CheckedIn:                false,
			PaymentReceivedIndicator: false,
			CreateUserID:             currentUser.ID,
			DateTimeCreated:          time.Now(),
			UpdateUserID:             currentUser.ID,
			DateTimeUpdated:          time.Now(),
		}
		if err = service.AthleteCompetitionEntryRepo.CreateEntry(&entry); err != nil {
			return err
		}
	}

	currentFollowCompEntries, err := service.AthleteCompetitionEntryRepo.SearchEntry(SearchAthleteCompetitionEntryCriteria{CompetitionID: registration.Competition.ID, AthleteID: registration.Couple.Follow.ID})
	if err != nil {
		log.Printf("[error] searching athlete competition entry of follow %v: %v", registration.Couple.Follow.FullName(), err)
		return err
	}
	if len(currentFollowCompEntries) == 0 {
		entry := AthleteCompetitionEntry{
			Competition:              registration.Competition,
			Athlete:                  registration.Couple.Follow,
			CheckedIn:                false,
			PaymentReceivedIndicator: false,
			CreateUserID:             currentUser.ID,
			DateTimeCreated:          time.Now(),
			UpdateUserID:             currentUser.ID,
			DateTimeUpdated:          time.Now(),
		}
		if err = service.AthleteCompetitionEntryRepo.CreateEntry(&entry); err != nil {
			return err
		}
	}

	// TODO: add logic for dropping competition entries.

	existingEntries, err := service.PartnershipEventEntryRepo.SearchPartnershipEventEntry(SearchPartnershipEventEntryCriteria{PartnershipID: registration.Couple.ID, CompetitionID: registration.Competition.ID})
	for _, each := range registration.EventsAdded {
		hasEntry := false
		for _, existing := range existingEntries {
			if each.ID == existing.Event.ID {
				hasEntry = true
				break
			}
		}
		if !hasEntry {
			coupleEntry := PartnershipEventEntry{
				Event:           each,
				Couple:          registration.Couple,
				CreateUserID:    currentUser.ID,
				DateTimeCreated: time.Now(),
				UpdateUserID:    currentUser.ID,
				DateTimeUpdated: time.Now(),
			}
			if err = service.PartnershipEventEntryRepo.CreatePartnershipEventEntry(&coupleEntry); err != nil {
				return err
			}
		}
	}

	for _, each := range registration.EventsDropped {
		idx := -1
		for i := 0; i < len(existingEntries); i++ {
			if each.ID == existingEntries[i].Event.ID {
				idx = i
				break
			}
		}
		if idx >= 0 {
			if err = service.PartnershipEventEntryRepo.DeletePartnershipEventEntry(existingEntries[idx]); err != nil {
				return err
			}
		}
	}

	// TODO: if partnership has no remaining event entries, then drop the competition entry of this couple

	// TODO: update attendance of a competition based on athlete competition entry

	// for scrutineer and organizer, check if they are either invited officials of the competition
	// Scrutineer/Organizer of one competitions should not be able to update the registration forms of other competitions

	// check if current user already has registration

	// if has existing registration, get the existing registration
	//
	return nil
}

type SearchEntryCriteria struct {
	CompetitionID int
	EventID       int
	FederationID  int
	DivisionID    int
	AgeID         int
	ProficiencyID int
	StyleID       int
	AthleteID     int
	PartnershipID int
}

// ValidateEventRegistration validates if the registration data is valid. This does not create the registration
func (service CompetitionRegistrationService) ValidateEventRegistration(currentUser Account, registration EventRegistrationForm) error {
	if registration.Couple.ID < 1 {
		return errors.New("partnership should be specified")
	}
	if registration.Competition.ID < 1 {
		return errors.New("competition should be specified")
	}

	if currentUser.HasRole(AccountTypeAthlete) && registration.Competition.GetStatus() != CompetitionStatusOpenRegistration {
		return errors.New("registration is no longer open")
	}

	// check if organizer is authorized to change this partnership's registration
	organizer := GetAccountByID(registration.Competition.CreateUserID, service.AccountRepository) // creator may not be the organizer of specified competition
	if currentUser.HasRole(AccountTypeOrganizer) && organizer.ID != currentUser.ID {
		return errors.New("not an authorized organizer to update the registration")
	}

	// check if the creator of the entry is competitor
	if currentUser.HasRole(AccountTypeAthlete) && (!registration.Couple.HasAthlete(currentUser.ID)) {
		// request was sent by people who are neither the lead or the follow of this partnership
		return errors.New("not an authorized athlete to update the registration")
	}

	// check if those events are valid and open for registration
	for _, each := range registration.EventsAdded {
		if each.StatusID != EVENT_STATUS_OPEN {
			return errors.New("event is not open for registration")
		}
	}

	// create competition entry for the lead, if the entry has not been created yet
	entries, hasEntryErr := service.AthleteCompetitionEntryRepo.SearchEntry(SearchAthleteCompetitionEntryCriteria{
		CompetitionID: registration.Competition.ID,
		AthleteID:     registration.Couple.Lead.ID,
	})
	if len(entries) != 1 || hasEntryErr != nil {
		service.AthleteCompetitionEntryRepo.CreateEntry(&AthleteCompetitionEntry{
			Competition: Competition{
				ID: registration.Competition.ID,
			},
			Athlete: registration.Couple.Lead,
		})
	}

	// create competition entry for the follow, if the entry has not been created yet
	entries, hasEntryErr = service.AthleteCompetitionEntryRepo.SearchEntry(SearchAthleteCompetitionEntryCriteria{
		CompetitionID: registration.Competition.ID,
		AthleteID:     registration.Couple.Lead.ID,
	})
	if len(entries) != 1 || hasEntryErr != nil {
		service.AthleteCompetitionEntryRepo.CreateEntry(&AthleteCompetitionEntry{
			Competition: Competition{
				ID: registration.Competition.ID,
			},
			Athlete: registration.Couple.Follow,
		})
	}

	// check if dropped events are open

	// check if added events are already added

	// check if dropped events are added

	// check event entries, and see if this partnership is still eligible for entering these events
	// TODO: eligibility check
	for _, each := range registration.EventsAdded {
		eventEntry := PartnershipEventEntry{
			Couple: registration.Couple,
			Event:  each,
		}
		eligibilityErr := checkEventEligibility(eventEntry)
		if eligibilityErr != nil {
			return eligibilityErr
		}
	}
	return nil
}

func (service CompetitionRegistrationService) SearchCompetitionEntries(criteria SearchEntryCriteria) (CompetitionEntryList, error) {
	var err error
	entries := CompetitionEntryList{}

	athleteEntries, err := service.AthleteCompetitionEntryRepo.SearchEntry(SearchAthleteCompetitionEntryCriteria{
		CompetitionID: criteria.CompetitionID,
		AthleteID:     criteria.AthleteID,
	})
	if err != nil {
		return entries, err
	}

	partnershipEntries, err := service.PartnershipCompetitionEntryRepo.SearchEntry(SearchPartnershipCompetitionEntryCriteria{
		CompetitionID: criteria.CompetitionID,
		PartnershipID: criteria.PartnershipID,
	})
	if err != nil {
		return entries, err
	}

	competitions, err := service.CompetitionRepository.SearchCompetition(SearchCompetitionCriteria{ID: criteria.CompetitionID})
	if err != nil || len(competitions) != 1 {
		return entries, errors.New(fmt.Sprintf("cannot find competition with ID = %v", criteria.CompetitionID))
	}

	entries.Competition = competitions[0]
	entries.AthleteEntries = athleteEntries
	entries.CoupleEntries = partnershipEntries

	return entries, nil
}

// SearchEventEntries searches the entries of one event. Parameters specified in the criteria will be used to locate
// all eligible events. For each event, a list of entries will be returned
func (service CompetitionRegistrationService) SearchEventEntries(criteria SearchEntryCriteria) ([]EventEntryList, error) {
	events, _ := service.EventRepository.SearchEvent(SearchEventCriteria{
		CompetitionID: criteria.CompetitionID,
		EventID:       criteria.EventID,
		FederationID:  criteria.FederationID,
		DivisionID:    criteria.DivisionID,
		AgeID:         criteria.AgeID,
		ProficiencyID: criteria.ProficiencyID,
		StyleID:       criteria.StyleID,
	})
	entryLists := make([]EventEntryList, 0)
	for _, each := range events {
		entries := EventEntryList{}

		athleteEntries, err := service.athleteEventEntryRepo.SearchAthleteEventEntry(SearchAthleteEventEntryCriteria{
			CompetitionID: criteria.CompetitionID,
			EventID:       each.ID,
			AthleteID:     criteria.AthleteID,
		})
		if err != nil {
			return entryLists, err
		}

		partnershipEntries, err := service.PartnershipEventEntryRepo.SearchPartnershipEventEntry(SearchPartnershipEventEntryCriteria{
			CompetitionID: criteria.CompetitionID,
			EventID:       each.ID,
			PartnershipID: criteria.PartnershipID,
		})
		if err != nil {
			return entryLists, err
		}

		entries.Event = each
		entries.AthleteEntries = athleteEntries
		entries.CoupleEntries = partnershipEntries
		entryLists = append(entryLists, entries)
	}
	return entryLists, nil
}

// CreateEntry takes the current user and the registration data and create new Competition Entry for
// each of the athlete
func (service CompetitionRegistrationService) CreateAthleteCompetitionEntry(currentUser Account, registration EventRegistrationForm) error {
	leadCompEntry := AthleteCompetitionEntry{
		Competition:              registration.Competition,
		CheckedIn:                false,
		CreateUserID:             currentUser.ID,
		DateTimeCreated:          time.Now(),
		UpdateUserID:             currentUser.ID,
		DateTimeUpdated:          time.Now(),
		Athlete:                  registration.Couple.Lead,
		PaymentReceivedIndicator: false,
	}
	followCompEntry := AthleteCompetitionEntry{
		Competition:              registration.Competition,
		CheckedIn:                false,
		CreateUserID:             currentUser.ID,
		DateTimeCreated:          time.Now(),
		UpdateUserID:             currentUser.ID,
		DateTimeUpdated:          time.Now(),
		Athlete:                  registration.Couple.Follow,
		PaymentReceivedIndicator: false,
	}

	service.AthleteCompetitionEntryService.CreateAthleteCompetitionEntry(&leadCompEntry)
	service.AthleteCompetitionEntryService.CreateAthleteCompetitionEntry(&followCompEntry)
	return nil
}

// CreateEntry takes the current user and registration data and create a Competition Entry for
// this Partnership
func (service CompetitionRegistrationService) CreatePartnershipCompetitionEntry(currentUser Account, registration EventRegistrationForm) error {

	partnershipEntry := PartnershipCompetitionEntry{
		Couple:          registration.Couple,
		Competition:     registration.Competition,
		CheckedIn:       false,
		CreateUserID:    currentUser.ID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    currentUser.ID,
		DateTimeUpdated: time.Now(),
	}
	partnershipEntry.createPartnershipCompetitionEntry(service.CompetitionRepository, service.PartnershipCompetitionEntryRepo)
	return nil
}

// CreatePartnershipEventEntries takes the current user and registration data to create a new Event Entry for this partnership
func (service CompetitionRegistrationService) CreatePartnershipEventEntries(currentUser Account, registration EventRegistrationForm) error {
	for _, each := range registration.EventsAdded {
		eventEntry := PartnershipEventEntry{
			Couple: registration.Couple,
			Event:  each,
		}
		createErr := service.PartnershipEventEntryRepo.CreatePartnershipEventEntry(&eventEntry)
		if createErr != nil {
			return createErr
		}
	}
	return nil
}

func (service CompetitionRegistrationService) SearchPartnershipEventEntries(criteria SearchEntryCriteria) ([]PartnershipEventEntry, error) {
	return service.PartnershipEventEntryRepo.SearchPartnershipEventEntry(SearchPartnershipEventEntryCriteria{
		CompetitionID: criteria.CompetitionID,
		PartnershipID: criteria.PartnershipID,
	})
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
			Event:        each,
			UpdateUserID: currentUser.ID,
			Couple:       registration.Couple,
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
