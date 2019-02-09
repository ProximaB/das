package businesslogic

import (
	"errors"
	"fmt"
	"time"
)

// Competition provides the base data structure for a competitive ballroom dance. All competitions in
// DAS must have some affiliation with a dancesport federation (Not Affiliated/Independent is also a Federation,)
type Competition struct {
	ID                        int
	FederationID              int
	Name                      string
	Street                    string
	City                      City
	State                     State
	Country                   Country
	StartDateTime             time.Time
	EndDateTime               time.Time
	CreateUserID              int
	DateTimeCreated           time.Time
	UpdateUserID              int
	DateTimeUpdated           time.Time
	ContactName               string
	ContactEmail              string
	ContactPhone              string
	statusID                  int
	Website                   string
	Attendance                int
	RegistrationOpenDateTime  time.Time
	RegistrationCloseDateTime time.Time
	officials                 map[int][]Account
}

// UpdateStatus will attempt to change the status of the caller competition to statusID, if the change is in logical order
func (comp *Competition) UpdateStatus(statusID int) error {
	if comp.statusID >= statusID && comp.statusID != 0 {
		return errors.New("cannot revert competition status")
	}
	if comp.statusID == CompetitionStatusClosed || comp.statusID == CompetitionStatusCancelled {
		return errors.New("competition status is locked")
	}
	comp.statusID = statusID
	return nil
}

func (comp *Competition) AddOfficial(role int, official Account) error {
	if comp.officials == nil {
		comp.officials = make(map[int][]Account)
	}

	if !official.HasRole(role) {
		return errors.New("This user does not have the appointed role provisioned.")
	}

	if comp.officials[role] == nil {
		comp.officials[role] = make([]Account, 0)
	}

	// check if this official is already added
	for _, each := range comp.officials[role] {
		if each.ID == official.ID {
			return errors.New("This user is already added to the list of officials")
		}
	}
	comp.officials[role] = append(comp.officials[role], official)

	return nil
}

func (comp Competition) GetOfficials(role int) ([]Account, error) {
	if comp.officials == nil {
		return nil, errors.New("No official has been assigned to this competition")
	}
	return comp.officials[role], nil
}

func (comp Competition) GetStatus() int {
	return comp.statusID
}

// SearchCompetitionCriteria specifies the parameters that can be used to search a Competition
type SearchCompetitionCriteria struct {
	ID            int       `schema:"id"`
	Name          string    `schema:"name"`
	FederationID  int       `schema:"federation"`
	StateID       int       `schema:"state"`
	CountryID     int       `schema:"country"`
	StartDateTime time.Time `schema:"start"`
	OrganizerID   int
	StatusID      int `schema:"status"`
}

type OrganizerUpdateCompetition struct {
	CompetitionID int       `json:"competition"`
	Name          string    `json:"name"`
	Website       string    `json:"website"`
	Status        int       `json:"status"`
	Address       string    `json:"street"`
	ContactName   string    `json:"contact"`
	ContactEmail  string    `json:"email"`
	ContactPhone  string    `json:"phone"`
	StartDate     time.Time `json:"start"`
	EndDate       time.Time `json:"end"`
	UpdateUserID  int
}

// ICompetitionRepository specifies the interface that a competition repository needs to implement to provide CRUD
// operations in the data repository
type ICompetitionRepository interface {
	CreateCompetition(competition *Competition) error
	SearchCompetition(criteria SearchCompetitionCriteria) ([]Competition, error)
	UpdateCompetition(competition Competition) error
	DeleteCompetition(competition Competition) error
}

// GetCompetitionByID guarantees getting a competition from the provided repository. In case failure happens,
// panic() will be invoked
func GetCompetitionByID(id int, repo ICompetitionRepository) (Competition, error) {
	searchResults, err := repo.SearchCompetition(SearchCompetitionCriteria{ID: id})
	if err != nil || searchResults == nil || len(searchResults) != 1 {
		return Competition{}, err
	}
	return searchResults[0], err
}

// CreateCompetition creates competition in competitionRepo, update records in provisionRepo, and
// add a new record to historyRepo
func CreateCompetition(competition Competition, competitionRepo ICompetitionRepository,
	provisionRepo IOrganizerProvisionRepository, historyRepo IOrganizerProvisionHistoryRepository) error {
	// check if data received is validationErr
	if validationErr := competition.validateCreateCompetition(); validationErr != nil {
		return validationErr
	}

	if competition.statusID == 0 {
		competition.statusID = CompetitionStatusPreRegistration
	}

	// check if organizer is provisioned with available competitions
	provisions, _ := provisionRepo.SearchOrganizerProvision(SearchOrganizerProvisionCriteria{
		OrganizerID: competition.CreateUserID,
	})
	if len(provisions) != 1 {
		return errors.New("no organizer record is found")
	}
	provision := provisions[0]
	if provision.Available < 1 {
		return errors.New("no available competition slot")
	}

	newProvision := provision.updateForCreateCompetition(competition)
	historyEntry := newProvisionHistoryEntry(newProvision, competition)
	updateOrganizerProvision(newProvision, historyEntry, provisionRepo, historyRepo)

	err := competitionRepo.CreateCompetition(&competition)
	if err != nil {
		// refund competition organizer's provision
		refundProvision := newProvision
		refundProvision.Available += 1
		refundProvision.Hosted -= 1

		refundEntry := OrganizerProvisionHistoryEntry{
			OrganizerRoleID: refundProvision.OrganizerRoleID,
			Amount:          1,
			Note:            fmt.Sprintf("Refund for failing in creating competition %v %v", competition.Name, competition.StartDateTime),
			CreateUserID:    competition.CreateUserID,
			DateTimeCreated: time.Now(),
			UpdateUserID:    competition.UpdateUserID,
			DateTimeUpdated: time.Now(),
		}
		updateOrganizerProvision(refundProvision, refundEntry, provisionRepo, historyRepo)
	}

	return err
}

func (comp Competition) validateCreateCompetition() error {
	if comp.FederationID < 1 {
		return errors.New("invalid federation")
	}
	if len(comp.Name) < 3 {
		return errors.New("comp name is too short")
	}
	if len(comp.Website) < 7 { // requires "http://"
		return errors.New("official comp website is required")
	}
	if comp.GetStatus() > CompetitionStatusClosedRegistration {
		return errors.New("cannot create comp that no longer allows new registration")
	}
	if comp.StartDateTime.After(comp.EndDateTime) {
		return errors.New("start date must be ahead of end date")
	}
	if comp.StartDateTime.Before(time.Now()) {
		return errors.New("comp must starts in a future time")
	}
	if comp.StartDateTime.After(time.Now().AddDate(1, 0, 0)) {
		return errors.New("cannot create far-future comp")
	}
	if len(comp.ContactName) < 3 {
		return errors.New("contact name is too short")
	}
	if len(comp.ContactEmail) < 5 {
		return errors.New("contact email is too short")
	}
	if len(comp.ContactPhone) < 9 {
		return errors.New("contact phone is too short")
	}
	if comp.City.ID < 1 {
		return errors.New("city is required")
	}
	if comp.State.ID < 1 {
		return errors.New("state is required")
	}
	if comp.Country.ID < 1 {
		return errors.New("country is required")
	}
	if comp.CreateUserID < 1 {
		return errors.New("unauthorized")
	}
	if comp.UpdateUserID < 1 {
		return errors.New("unauthorized")
	}
	return nil
}

// UpdateCompetition updates the existing competition with the information provided in OrganizerUpdateCompetition
func UpdateCompetition(user *Account, competition OrganizerUpdateCompetition, repo ICompetitionRepository) error {
	// check if user is the owner of the original competition
	competitions, err := repo.SearchCompetition(SearchCompetitionCriteria{ID: competition.CompetitionID})
	if err != nil {
		return err
	}
	if competitions == nil || len(competitions) != 1 || competitions[0].ID == 0 {
		return errors.New("cannot find this competition")
	}
	if validationErr := validateUpdateCompetition(user, competitions[0], &competition, repo); validationErr != nil {
		return validationErr
	}

	if competitions[0].GetStatus() == CompetitionStatusOpenRegistration ||
		competitions[0].GetStatus() == CompetitionStatusClosedRegistration {
		// TODO: reimplement event update
		/*if updateEventErr := dataaccess.UpdateCompetitionEventStatus(dataaccess.DATABASE, competition.ID, competitions[0].StatusID); updateEventErr != nil {
			return updateEventErr
		}*/
	}

	// if competition is cancelled, refund the slot. competition cannot be cancelled unless it is done by site administrator

	return repo.UpdateCompetition(competitions[0])
}

func validateUpdateCompetition(user *Account,
	competition Competition,
	updateDTO *OrganizerUpdateCompetition,
	repo ICompetitionRepository) error {
	if user.ID != competition.CreateUserID {
		return errors.New("not authorized to update this competition")
	}

	if competition.GetStatus() > updateDTO.Status {
		return errors.New("cannot change competition status back")
	}

	if competition.GetStatus() == CompetitionStatusClosed {
		return errors.New("competition is closed")
	}
	if len(updateDTO.Name) < 3 {
		return errors.New("invalid competition name")
	}
	if len(updateDTO.Website) < 3 {
		// TODO: need a better url validation mechanics
		return errors.New("website link is too short")
	}
	if updateDTO.StartDate.After(updateDTO.EndDate) {
		return errors.New("competition must start before it ends")
	}

	if updateDTO.StartDate.Before(time.Now()) {
		return errors.New("cannot start competition in the past")
	}

	if updateDTO.StartDate.After(time.Now().AddDate(1, 0, 0)) {
		return errors.New("cannot create competition that starts a year later")
	}

	return nil
}

type IEventMetaRepository interface {
	GetEventUniqueFederations(competition Competition) ([]Federation, error)
	GetEventUniqueDivisions(competition Competition) ([]Division, error)
	GetEventUniqueAges(competition Competition) ([]Age, error)
	GetEventUniqueProficiencies(competition Competition) ([]Proficiency, error)
	GetEventUniqueStyles(competition Competition) ([]Style, error)
}

// Get a list of unique federations that a competition has
func (competition Competition) GetEventUniqueFederations(eventRepository IEventMetaRepository) ([]Federation, error) {
	return eventRepository.GetEventUniqueFederations(competition)
}
func (competition Competition) GetEventUniqueDivisions(eventRepository IEventMetaRepository) ([]Division, error) {
	return eventRepository.GetEventUniqueDivisions(competition)
}
func (competition Competition) GetEventUniqueAges(eventRepository IEventMetaRepository) ([]Age, error) {
	return eventRepository.GetEventUniqueAges(competition)
}
func (competition Competition) GetEventUniqueProficiencies(eventRepository IEventMetaRepository) ([]Proficiency, error) {
	return eventRepository.GetEventUniqueProficiencies(competition)
}
func (competition Competition) GetEventUniqueStyles(eventRepository IEventMetaRepository) ([]Style, error) {
	return eventRepository.GetEventUniqueStyles(competition)
}
