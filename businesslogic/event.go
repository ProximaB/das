package businesslogic

import (
	"errors"
	"log"
	"sort"
	"time"
)

const (
	// EventCategoryCompetitiveBallroom is a constant for Competitive Ballroom events
	EventCategoryCompetitiveBallroom = 1
	// EventCategoryShowDance is a constant for Show Dance events
	EventCategoryShowDance = 2
	// EventCategoryCabaret is a constant for Cabaret events
	EventCategoryCabaret = 3
	// EventCategoryTheatreArt is a constant for Theatre Art events
	EventCategoryTheatreArt = 4
)

// SearchEventCriteria specifies the parameters that can be used to search events
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
	OrganizerID   int `schema:"organizerID"`
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
	Federation      Federation
	DivisionID      int
	Division        Division
	AgeID           int
	Age             Age
	ProficiencyID   int
	Proficiency     Proficiency
	StyleID         int
	Style           Style
	dances          map[int]bool
	eventDances     map[int]EventDance
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
	e.eventDances = make(map[int]EventDance)
	return &e
}

// IEventRepository specifies the interface that a struct need to implement to function as a repository for businesslogic
type IEventRepository interface {
	SearchEvent(criteria SearchEventCriteria) ([]Event, error)
	CreateEvent(event *Event) error
	UpdateEvent(event Event) error
	DeleteEvent(event Event) error
}

// GetDances returns the ID of dances of the caller event
func (event Event) GetDances() []int {
	keys := make([]int, 0)
	for k := range event.dances {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

// AddDance adds a dance's ID
func (event *Event) AddDance(danceID int) {
	if !event.dances[danceID] {
		event.dances[danceID] = true
	}
}

// RemoveDance removes the dance of the provided ID from the event
func (event *Event) RemoveDance(danceID int) {
	if event.dances[danceID] {
		delete(event.dances, danceID)
	}
}

// SetDances replaces the dances of the event with the new dances
func (event *Event) SetDances(dances []int) {
	event.dances = make(map[int]bool)
	for _, each := range dances {
		event.dances[each] = true
	}
}

// HasDance checks if a dance of the provided ID is in the event
func (event Event) HasDance(danceID int) bool {
	return event.dances[danceID]
}

func (event *Event) AddEventDance(eveDance EventDance) {
	if _, has := event.eventDances[eveDance.ID]; !has {
		event.eventDances[eveDance.ID] = eveDance
	}
}

func (event Event) GetEventDances() []EventDance {
	output := make([]EventDance, 0)
	for _, v := range event.eventDances {
		output = append(output, v)
	}
	return output
}

// EquivalentTo checks if two events are equivalent in Federation, Division, Age, Proficiency, Style, and dances
func (event Event) EquivalentTo(other Event) bool {
	if event.FederationID != other.FederationID {
		return false
	}
	if event.DivisionID != other.DivisionID {
		return false
	}
	if event.AgeID != other.AgeID {
		return false
	}
	if event.ProficiencyID != other.ProficiencyID {
		return false
	}
	if event.StyleID != other.StyleID {
		return false
	}
	if len(event.dances) != len(other.dances) {
		return false
	}
	sameDances := true
	for k := range event.dances {
		if !(other.dances[k]) {
			sameDances = false
			break
		}
	}
	return sameDances
}

// GetEventByID retrieves an existing Event from the provided repository by its ID
func GetEventByID(id int, repo IEventRepository) (Event, error) {
	results, err := repo.SearchEvent(SearchEventCriteria{EventID: id})
	return results[0], err
}

// OrganizerEventService provides a layer of abstraction of services used by organizers to manage events of competitions
type OrganizerEventService struct {
	accountRepo       IAccountRepository
	roleRepo          IAccountRoleRepository
	competitionRepo   ICompetitionRepository
	eventRepo         IEventRepository
	eventDanceRepo    IEventDanceRepository
	eventTemplateRepo ICompetitionEventTemplateRepository
	factory           CompetitionEventFactory
}

func NewOrganizerEventService(accountRepo IAccountRepository,
	roleRepo IAccountRoleRepository,
	compRepo ICompetitionRepository,
	eventRepo IEventRepository,
	eventDanceRepo IEventDanceRepository,
	eventTemplateRepo ICompetitionEventTemplateRepository,
	factory CompetitionEventFactory) OrganizerEventService {

	return OrganizerEventService{
		accountRepo,
		roleRepo,
		compRepo,
		eventRepo,
		eventDanceRepo,
		eventTemplateRepo,
		factory}
}

func (service OrganizerEventService) generateTemplateEvents(templateID int) ([]Event, error) {
	events := make([]Event, 0)
	results, err := service.eventTemplateRepo.SearchCompetitionEventTemplates(SearchCompetitionEventTemplateCriteria{
		ID: templateID,
	})
	if err != nil {
		return events, err
	}

	for _, each := range results[0].TemplateEvents {
		item, genErr := service.factory.GenerateEvent(each.Federation, each.Division, each.Age, each.Proficiency, each.Style, each.Dances)
		if genErr != nil {
			return events, genErr
		}
		events = append(events, item)
	}
	return events, nil
}

func (service OrganizerEventService) GenerateEventsFromTemplate(competitionID int, templateID int) error {
	events, err := service.generateTemplateEvents(templateID)
	if err != nil {
		return err
	}
	for i := 0; i < len(events); i++ {
		genErr := service.CreateEvent(&events[i])
		if genErr != nil {
			return genErr
		}
	}
	return nil
}

func (service OrganizerEventService) SearchCompetitionEventTemplate(criteria SearchCompetitionEventTemplateCriteria) ([]CompetitionEventTemplate, error) {
	return service.eventTemplateRepo.SearchCompetitionEventTemplates(criteria)
}

func (service OrganizerEventService) CreateEvent(event *Event) error {
	competition, _ := GetCompetitionByID(event.CompetitionID, service.competitionRepo)

	// check if competition is still at the right status
	if competition.GetStatus() != CompetitionStatusPreRegistration {
		return errors.New("events can only be added when competition is in pre-registration")
	} else if competition.CreateUserID != event.CreateUserID {
		return errors.New("not authorized to create event for this competition")
	}

	// check if specified events were created
	similarEvents, _ := service.eventRepo.SearchEvent(SearchEventCriteria{
		CompetitionID: event.CompetitionID,
		CategoryID:    event.CategoryID,
		FederationID:  event.FederationID,
		DivisionID:    event.DivisionID,
		AgeID:         event.AgeID,
		ProficiencyID: event.ProficiencyID,
		StyleID:       event.StyleID,
	})

	// for each similar event, check if they share dances
	for _, eachEvent := range similarEvents {
		for _, eachDance := range event.GetDances() {
			if eachEvent.HasDance(eachDance) {
				return errors.New("specified dance is already in this event")
			}
		}
	}

	// if no errors, create the event
	// step 1: create an event
	createEventErr := service.eventRepo.CreateEvent(event)
	if createEventErr != nil {
		return createEventErr
	}
	if event.ID == 0 {
		log.Printf("[error] creating event %v returned with ID of 0", *event)
		return errors.New("event could not be created")
	}

	// step 2: create all the eventDances. requires primary key returned from the previous step
	for _, each := range event.GetDances() {
		eventDance := NewEventDance(*event, each)
		if createDancesErr := service.eventDanceRepo.CreateEventDance(eventDance); createDancesErr != nil {
			return createDancesErr
		}
	}
	return nil
}

func (service OrganizerEventService) SearchEvents(criteria SearchEventCriteria) ([]Event, error) {
	return service.eventRepo.SearchEvent(criteria)
}

func (service OrganizerEventService) ApplyCollegiateTemplate(competitionID int) error {
	return errors.New("Not implemented")
}

func (service OrganizerEventService) ApplyNDCATemplate(competitionID int) error {
	return errors.New("Not implemented")
}

// CreateEvent will check if event is valid, and create the in the provided IEventRepository. If competition
func CreateEvent(event Event, compRepo ICompetitionRepository, eventRepo IEventRepository, eventDanceRepo IEventDanceRepository) error {

	competition, _ := GetCompetitionByID(event.CompetitionID, compRepo)

	// check if competition is still at the right status
	if competition.GetStatus() != CompetitionStatusPreRegistration {
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

	// for each similar event, check if they share dances
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
	}
	if event.ID == 0 {
		return errors.New("event could not be created")
	}

	// step 2: create all the eventDances. requires primary key returned from the previous step
	for _, each := range event.GetDances() {
		eventDance := NewEventDance(event, each)
		if createDancesErr := eventDanceRepo.CreateEventDance(eventDance); createDancesErr != nil {
			return createDancesErr
		}
	}
	return nil
}

func (event Event) validate(dances []EventDance,
	federationRepo IFederationRepository,
	divisionRepo IDivisionRepository,
	ageRepo IAgeRepository,
	proficiencyRepo IProficiencyRepository,
	styleRepo IStyleRepository,
	danceRepo IDanceRepository) error {
	// check if federation exists
	targetFederations, err := federationRepo.SearchFederation(SearchFederationCriteria{ID: event.FederationID})
	if err != nil {
		return err
	}

	// check if division exists
	divisions, err := divisionRepo.SearchDivision(SearchDivisionCriteria{ID: event.DivisionID})
	if err != nil {
		return err
	}
	targetDivision := divisions[0]

	// check if division is part of this federation
	if targetDivision.FederationID != targetFederations[0].ID {
		return errors.New("specified division is not part of this federation")
	}

	// check if age category exists
	targetAges, err := ageRepo.SearchAge(SearchAgeCriteria{AgeID: event.AgeID})
	if err != nil {
		return err
	}

	// check if age category is part of this division
	if targetAges[0].DivisionID != targetDivision.ID {
		return errors.New("specified age category is not part of this division")
	}

	// check if proficiency is part of this division
	targetSkills, err := proficiencyRepo.SearchProficiency(SearchProficiencyCriteria{ProficiencyID: event.ProficiencyID})
	if targetSkills[0].DivisionID != targetDivision.ID {
		return errors.New("specified proficiency is not part of this division")
	}

	// check if style exists
	targetStyles, err := styleRepo.SearchStyle(SearchStyleCriteria{StyleID: event.StyleID})
	if err != nil {
		return errors.New("specified style does not exist")
	}

	// check if there are duplicated dance
	unique := map[int]bool{}
	result := make([]EventDance, 0)
	for _, each := range dances {
		if unique[each.DanceID] == false {
			// check if dance exists
			dances, err := danceRepo.SearchDance(SearchDanceCriteria{DanceID: each.DanceID})
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
