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
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"time"
)

// Competition provides the base data structure for a competitive ballroom dance. All competitions in
// DAS must have some affiliation with a dancesport federation (Not Affiliated/Independent is also a Federation,)
type Competition struct {
	ID              int
	FederationID    int
	Name            string
	Street          string
	City            referencebll.City
	State           referencebll.State
	Country         referencebll.Country
	StartDateTime   time.Time
	EndDateTime     time.Time
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
	ContactName     string
	ContactEmail    string
	ContactPhone    string
	statusID        int
	Website         string
	Attendance      int
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
	if validationErr := validateCreateCompetition(competition); validationErr != nil {
		return validationErr
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
	historyEntry := newProvisionHistory(newProvision, competition)
	updateOrganizerProvision(newProvision, historyEntry, provisionRepo, historyRepo)

	err := competitionRepo.CreateCompetition(&competition)

	return err
}

func validateCreateCompetition(competition Competition) error {
	if competition.FederationID < 1 {
		return errors.New("invalid federation")
	}
	if len(competition.Name) < 3 {
		return errors.New("competition name is too short")
	}
	if len(competition.Website) < 7 { // requires "http://"
		return errors.New("official competition website is required")
	}
	if competition.GetStatus() > CompetitionStatusClosedRegistration {
		return errors.New("cannot create competition that no longer allows new registration")
	}
	if competition.StartDateTime.After(competition.EndDateTime) {
		return errors.New("start date must be ahead of end date")
	}
	if competition.StartDateTime.Before(time.Now()) {
		return errors.New("competition must starts in a future time")
	}
	if competition.StartDateTime.After(time.Now().AddDate(1, 0, 0)) {
		return errors.New("cannot create far-future competition")
	}
	if len(competition.ContactName) < 3 {
		return errors.New("contact name is too short")
	}
	if len(competition.ContactEmail) < 5 {
		return errors.New("contact email is too short")
	}
	if len(competition.ContactPhone) < 9 {
		return errors.New("contact phone is too short")
	}
	if competition.City.ID < 1 {
		return errors.New("city is required")
	}
	if competition.State.ID < 1 {
		return errors.New("state is required")
	}
	if competition.Country.ID < 1 {
		return errors.New("country is required")
	}
	if competition.CreateUserID < 1 {
		return errors.New("unauthorized")
	}
	if competition.UpdateUserID < 1 {
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
		// TODO: reimplement eventdal update
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
	GetEventUniqueFederations(competition Competition) ([]referencebll.Federation, error)
	GetEventUniqueDivisions(competition Competition) ([]referencebll.Division, error)
	GetEventUniqueAges(competition Competition) ([]referencebll.Age, error)
	GetEventUniqueProficiencies(competition Competition) ([]referencebll.Proficiency, error)
	GetEventUniqueStyles(competition Competition) ([]referencebll.Style, error)
}

// Get a list of unique federations that a competition has
func (competition Competition) GetEventUniqueFederations(eventRepository IEventMetaRepository) ([]referencebll.Federation, error) {
	return eventRepository.GetEventUniqueFederations(competition)
}
func (competition Competition) GetEventUniqueDivisions(eventRepository IEventMetaRepository) ([]referencebll.Division, error) {
	return eventRepository.GetEventUniqueDivisions(competition)
}
func (competition Competition) GetEventUniqueAges(eventRepository IEventMetaRepository) ([]referencebll.Age, error) {
	return eventRepository.GetEventUniqueAges(competition)
}
func (competition Competition) GetEventUniqueProficiencies(eventRepository IEventMetaRepository) ([]referencebll.Proficiency, error) {
	return eventRepository.GetEventUniqueProficiencies(competition)
}
func (competition Competition) GetEventUniqueStyles(eventRepository IEventMetaRepository) ([]referencebll.Style, error) {
	return eventRepository.GetEventUniqueStyles(competition)
}
