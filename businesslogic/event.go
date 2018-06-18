package businesslogic

import (
	"errors"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"sort"
	"time"
)

const (
	EVENT_CATEGORY_COMPETITIVE_BALLROOM = 1
	EVENT_CATEGORY_SHOWDANCE            = 2
	EVENT_CATEGORY_CABARET              = 3
	EVENT_CATEGORY_THEATRE_ART          = 4
)

type SearchEventCriteria struct {
	EventID       int `schema:"id"`
	CompetitionID int `schema:"competition"`
	CategoryID    int `schema:"category"`
	FederationID  int `schema:"federation"`
	DivisionID    int `schema:"division"`
	AgeID         int `schema:"age"`
	ProficiencyID int `schema:"proficiency"`
	StyleID       int `schema:"style"`
	StatusID      int `schema:"status"`
}

// Event contains data that are used for a generic competitive ballroom event, though it can be used for
// theatre art or cabaret events as well by leaving unnecessary fields empty or with default values.
type Event struct {
	ID              int
	CompetitionID   int
	CategoryID      int // ballroom, cabaret, theater art
	Description     string
	StatusID        int
	FederationID    int
	DivisionID      int
	AgeID           int
	ProficiencyID   int
	StyleID         int
	dances          map[int]bool
	Rounds          []int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

// NewEvent create a new
func NewEvent() *Event {
	e := Event{}
	e.dances = make(map[int]bool)
	return &e
}

type IEventRepository interface {
	SearchEvent(criteria SearchEventCriteria) ([]Event, error)
	CreateEvent(event *Event) error
	UpdateEvent(event Event) error
	DeleteEvent(event Event) error
}

func (event Event) GetDances() []int {
	keys := make([]int, 0)
	for k := range event.dances {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func (event *Event) AddDance(dance int) {
	if !event.dances[dance] {
		event.dances[dance] = true
	}
}

func (event *Event) RemoveDance(dance int) {
	if event.dances[dance] {
		delete(event.dances, dance)
	}
}

func (event *Event) SetDances(dances []int) {
	event.dances = make(map[int]bool)
	for _, each := range dances {
		event.dances[each] = true
	}
}

func (event Event) HasDance(danceID int) bool {
	return event.dances[danceID]
}

func (event Event) EquivalentTo(other Event) bool {
	if event.FederationID != other.FederationID {
		return false
	}
	if event.DivisionID != other.DivisionID {
		return false
	}
	if len(event.dances) != len(other.dances) {
		return false
	}
	sameDances := true
	for k, _ := range event.dances {
		if !(other.dances[k]) {
			sameDances = false
			break
		}
	}
	return sameDances
}

func GetEventByID(id int, repo IEventRepository) (Event, error) {
	results, err := repo.SearchEvent(SearchEventCriteria{EventID: id})
	return results[0], err
}

func CreateEvent(event Event, compRepo ICompetitionRepository, eventRepo IEventRepository, eventDanceRepo IEventDanceRepository) error {

	// check if competition is still at the right status
	compSearchResults, _ := compRepo.SearchCompetition(SearchCompetitionCriteria{ID: event.CompetitionID})
	if len(compSearchResults) != 1 {
		return errors.New("cannot find competition")
	}

	competition := compSearchResults[0]

	if competition.GetStatus() != COMPETITION_STATUS_PRE_REGISTRATION {
		return errors.New("events can only be added when competition is in pre-registration")
	} else if competition.CreateUserID != event.CreateUserID {
		return errors.New("not authorized to create event for this competition")
	}

	// check if specified events were created
	similarEvents, _ := eventRepo.SearchEvent(SearchEventCriteria{
		CompetitionID: event.CompetitionID,
		CategoryID:    event.CategoryID,
		FederationID:  event.FederationID,
		DivisionID:    event.DivisionID,
		AgeID:         event.AgeID,
		ProficiencyID: event.ProficiencyID,
		StyleID:       event.StyleID,
	})

	// for each similar event, check if they have the dance included
	for _, eachEvent := range similarEvents {
		for _, eachDance := range event.GetDances() {
			if eachEvent.HasDance(eachDance) {
				return errors.New("specified dance is already in this event")
			}
		}
	}

	// if no errors, create the event
	// step 1: create an event
	createEventErr := eventRepo.CreateEvent(&event)
	if createEventErr != nil {
		return createEventErr
	} else if event.ID == 0 {
		return errors.New("event was not created on time")
	}

	// step 2: create all the eventDances.
	for _, each := range event.GetDances() {
		eventDance := NewEventDance(event, each)
		createDancesErr := eventDanceRepo.CreateEventDance(eventDance)
		if createDancesErr != nil {
			return createDancesErr
		}
	}

	// Step 2 requires primary key returned from the previous step
	return nil
}

func (event Event) validate(dances []EventDance,
	federationRepo referencebll.IFederationRepository,
	divisionRepo referencebll.IDivisionRepository,
	ageRepo referencebll.IAgeRepository,
	proficiencyRepo referencebll.IProficiencyRepository,
	styleRepo referencebll.IStyleRepository,
	danceRepo referencebll.IDanceRepository) error {
	// check if federation exists
	targetFederations, err := federationRepo.SearchFederation(referencebll.SearchFederationCriteria{ID: event.FederationID})
	if err != nil {
		return err
	}

	// check if division exists
	divisions, err := divisionRepo.SearchDivision(referencebll.SearchDivisionCriteria{ID: event.DivisionID})
	if err != nil {
		return err
	}
	targetDivision := divisions[0]

	// check if division is part of this federation
	if targetDivision.FederationID != targetFederations[0].ID {
		return errors.New("specified division is not part of this federation")
	}

	// check if age category exists
	targetAges, err := ageRepo.SearchAge(referencebll.SearchAgeCriteria{AgeID: event.AgeID})
	if err != nil {
		return err
	}

	// check if age category is part of this division
	if targetAges[0].DivisionID != targetDivision.ID {
		return errors.New("specified age category is not part of this division")
	}

	// check if proficiency is part of this division
	targetSkills, err := proficiencyRepo.SearchProficiency(referencebll.SearchProficiencyCriteria{ProficiencyID: event.ProficiencyID})
	if targetSkills[0].DivisionID != targetDivision.ID {
		return errors.New("specified proficiency is not part of this division")
	}

	// check if style exists
	targetStyles, err := styleRepo.SearchStyle(referencebll.SearchStyleCriteria{StyleID: event.StyleID})
	if err != nil {
		return errors.New("specified style does not exist")
	}

	// check if there are duplicated dance
	unique := map[int]bool{}
	result := make([]EventDance, 0)
	for _, each := range dances {
		if unique[each.DanceID] == false {
			// check if dance exists
			dances, err := danceRepo.SearchDance(referencebll.SearchDanceCriteria{DanceID: each.DanceID})
			if err != nil {
				return err
			}
			targetDance := dances[0]
			if targetDance.StyleID != targetStyles[0].ID {
				return errors.New("specified dance is not part of this style")
			}
			unique[each.DanceID] = true
			result = append(result, each)
		}
	}
	if len(result) != len(dances) {
		return errors.New("selected dances contain duplicates")
	}

	// check if there are enough dances
	if len(dances) < 1 || len(event.GetDances()) < 1 {
		return errors.New("not enough dance specified")
	}

	return nil
}
