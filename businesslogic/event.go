package businesslogic

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"time"
)

const (
	// EVENT_STATUS_DRAFT is the stage where it is newly created not open to registration or competition
	EVENT_STATUS_DRAFT = 1
	// EVENT_STATUS_OPEN is the stage where it is open to registration, but competition has not started
	EVENT_STATUS_OPEN = 2
	// EVENT_STATUS_RUNNING is the stage where the event is currently danced. Adding entries is prohibited but dropping entries is okay.
	EVENT_STATUS_RUNNING = 3
	// EVENT_STATUS_CLOSED is the stage where the final round of the event is danced (Regardless of the posting of placement)
	EVENT_STATUS_CLOSED = 4
	// EVENT_STATUS_CANCELED is the stage where the event is canceled due to non-technical reasons while the competition is running
	EVENT_STATUS_CANCELED = 5
)

// EventStatus defines status of an event
type EventStatus struct {
	ID              int
	Name            string
	Abbreviation    string
	Description     string
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}

// IEventStatusRepository defines the method that a EventStatus Repository should implement.
type IEventStatusRepository interface {
	// GetEventStatus should return *all* the stored event status in the repository
	GetEventStatus() ([]EventStatus, error)
}

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
	CompetitionID int `schema:"competitionId,required"`
	CategoryID    int `schema:"category"`
	FederationID  int `schema:"federationId"`
	DivisionID    int `schema:"divisionId"`
	AgeID         int `schema:"ageId"`
	ProficiencyID int `schema:"proficiencyId"`
	StyleID       int `schema:"styleId"`
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
	Prefix          string // Organizer can use prefix to customize event
	Suffix          string // Organizer can use suffix to customize event
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

// ToString returns the String representation of an event
func (event Event) ToString() string {
	output := fmt.Sprintf("%v %v %v %v %v",
		event.Federation.Name,
		event.Division.Name,
		event.Age.Name,
		event.Proficiency.Name,
		event.Style.Name,
	)
	return output
}

// EventDance represents the many-to-many relationship between competition event and dance references.
type EventDance struct {
	ID              int
	EventID         int
	DanceID         int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

type SearchEventDanceCriteria struct {
	EventDanceID  int
	CompetitionID int
	EventID       int
}

type IEventDanceRepository interface {
	SearchEventDance(criteria SearchEventDanceCriteria) ([]EventDance, error)
	CreateEventDance(eventDance *EventDance) error
	DeleteEventDance(eventDance EventDance) error
	UpdateEventDance(eventDance EventDance) error
}

func NewEventDance(event Event, danceID int) *EventDance {
	return &EventDance{
		EventID:         event.ID,
		DanceID:         danceID,
		CreateUserID:    event.CreateUserID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    event.UpdateUserID,
		DateTimeUpdated: time.Now(),
	}
}

// OrganizerEventService provides a layer of abstraction of services used by organizers to manage events of competitions
type OrganizerEventService struct {
	accountRepo       IAccountRepository
	roleRepo          IAccountRoleRepository
	competitionRepo   ICompetitionRepository
	eventRepo         IEventRepository
	eventDanceRepo    IEventDanceRepository
	eventTemplateRepo ICompetitionEventTemplateRepository
	federationRepo    IFederationRepository
	divisionRepo      IDivisionRepository
	ageRepo           IAgeRepository
	proficiencyRepo   IProficiencyRepository
	styleRepo         IStyleRepository
	danceRepo         IDanceRepository
	factory           CompetitionEventFactory
}

func NewOrganizerEventService(accountRepo IAccountRepository,
	roleRepo IAccountRoleRepository,
	compRepo ICompetitionRepository,
	eventRepo IEventRepository,
	eventDanceRepo IEventDanceRepository,
	eventTemplateRepo ICompetitionEventTemplateRepository,
	federationRepo IFederationRepository,
	divisionRepo IDivisionRepository,
	ageRepo IAgeRepository,
	proficiencyRepo IProficiencyRepository,
	styleRepo IStyleRepository,
	danceRepo IDanceRepository) OrganizerEventService {
	eventFactory := CompetitionEventFactory{
		FederationRepo:  federationRepo,
		DivisionRepo:    divisionRepo,
		AgeRepo:         ageRepo,
		ProficiencyRepo: proficiencyRepo,
		StyleRepo:       styleRepo,
		DanceRepo:       danceRepo,
	}
	return OrganizerEventService{
		accountRepo,
		roleRepo,
		compRepo,
		eventRepo,
		eventDanceRepo,
		eventTemplateRepo,
		federationRepo,
		divisionRepo,
		ageRepo,
		proficiencyRepo,
		styleRepo,
		danceRepo,
		eventFactory}
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
	searchResults, searchErr := service.competitionRepo.SearchCompetition(SearchCompetitionCriteria{ID: event.CompetitionID})
	if len(searchResults) != 1 {
		return errors.New(fmt.Sprintf("cannot find competition with ID = %d", event.CompetitionID))
	}
	if searchErr != nil {
		return searchErr
	}

	competition := searchResults[0]

	// check if competition is still at the right status
	if competition.GetStatus() != CompetitionStatusPreRegistration {
		return errors.New("events can only be added when competition is in pre-registration")
	} else if competition.CreateUserID != event.CreateUserID {
		// Only the creator/owner of the competition can create events for the competition.
		return errors.New("not authorized to create event for this competition")
	}

	validationErr := service.ValidateEvent(*event)
	if validationErr != nil {
		log.Printf("[error] event %v is not valid: %v", *event, validationErr)
		return validationErr
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

func (service OrganizerEventService) ValidateEvent(event Event) error {
	// check if federation exists
	if targetFederations, err := service.federationRepo.SearchFederation(SearchFederationCriteria{
		ID: event.FederationID,
	}); err != nil {
		return err
	} else if len(targetFederations) == 0 {
		return errors.New(fmt.Sprintf("cannot find Federation with ID  = %d", event.FederationID))
	} else {
		event.Federation = targetFederations[0]
	}

	// check if division exists
	if divisions, err := service.divisionRepo.SearchDivision(SearchDivisionCriteria{
		ID:           event.DivisionID,
		FederationID: event.FederationID,
	}); err != nil {
		return err
	} else if len(divisions) == 0 {
		return errors.New(fmt.Sprintf("cannot find Division with ID = %d and Federation ID = %d", event.DivisionID, event.FederationID))
	} else {
		event.Division = divisions[0]
	}

	// check if age category exists
	if targetAges, err := service.ageRepo.SearchAge(SearchAgeCriteria{
		AgeID:      event.AgeID,
		DivisionID: event.DivisionID,
	}); err != nil {
		return err
	} else if len(targetAges) == 0 {
		return errors.New(fmt.Sprintf("cannot find Age with ID = %d and Division ID = %d", event.AgeID, event.DivisionID))
	} else {
		event.Age = targetAges[0]
	}

	// check if proficiency is part of this division
	if targetSkills, err := service.proficiencyRepo.SearchProficiency(SearchProficiencyCriteria{
		ProficiencyID: event.ProficiencyID,
		DivisionID:    event.DivisionID,
	}); err != nil {
		return err
	} else if len(targetSkills) == 0 {
		return errors.New(fmt.Sprintf("cannot find Proficiency with ID = %d and Division ID = %d", event.ProficiencyID, event.DivisionID))
	}

	// check if style exists
	if targetStyles, err := service.styleRepo.SearchStyle(SearchStyleCriteria{StyleID: event.StyleID}); err != nil {
		return err
	} else if len(targetStyles) != 1 {
		return errors.New(fmt.Sprintf("cannot find Style with ID = %d", event.StyleID))
	}

	// check if there are duplicated dance
	uniqueDanceIDs := map[int]bool{}
	uniqueEventDances := make([]Dance, 0)
	for _, each := range event.GetDances() {
		if !uniqueDanceIDs[each] {
			// check if dance exists
			if dances, err := service.danceRepo.SearchDance(SearchDanceCriteria{DanceID: each, StyleID: event.StyleID}); err != nil {
				return err
			} else if len(dances) != 1 {
				return errors.New(fmt.Sprintf("cannot find Dance with ID = %d and Style ID = %d", each, event.StyleID))
			} else {
				targetDance := dances[0]
				uniqueDanceIDs[each] = true
				uniqueEventDances = append(uniqueEventDances, targetDance)
			}
		}
	}
	if len(uniqueEventDances) != len(uniqueEventDances) {
		return errors.New("selected dances contain duplicates")
	}

	// check if there are enough dances
	if len(event.GetDances()) < 1 {
		return errors.New("not enough dance specified")
	}

	// check if specified events were created
	similarEvents, _ := service.eventRepo.SearchEvent(SearchEventCriteria{
		CompetitionID: event.CompetitionID,
		// CategoryID:    event.CategoryID, // YH: as of 2019-03-03, category of event is not used and should not be a factor
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

	return nil
}

func (service OrganizerEventService) DeleteEvent(event Event, currentUser Account) error {
	competitions, searchCompErr := service.competitionRepo.SearchCompetition(SearchCompetitionCriteria{ID: event.CompetitionID})
	if searchCompErr != nil {
		return searchCompErr
	}
	if len(competitions) != 1 {
		return errors.New(fmt.Sprintf("cannot find the competition of event %v", event.ToString()))
	}
	if competitions[0].GetStatus() == CompetitionStatusClosed {
		return errors.New("the competition is concluded")
	}
	if event.CreateUserID != currentUser.ID {
		return errors.New("not authorized to delete this event")
	}
	if event.StatusID == EVENT_STATUS_RUNNING {
		return errors.New("event is running and cannot be deleted")
	}
	if event.StatusID == EVENT_STATUS_CLOSED {
		return errors.New("event is closed and cannot be deleted")
	}
	return service.eventRepo.DeleteEvent(event)
}

// CompetitionEventFactory can generate event based on specification of event attributes
type CompetitionEventFactory struct {
	FederationRepo  IFederationRepository
	DivisionRepo    IDivisionRepository
	AgeRepo         IAgeRepository
	ProficiencyRepo IProficiencyRepository
	StyleRepo       IStyleRepository
	DanceRepo       IDanceRepository
}

func (factory CompetitionEventFactory) GenerateEvent(federationName, divisionName, ageName, proficiencyName, styleName string, danceNames []string) (Event, error) {
	event := NewEvent()
	if federationName != "" {
		searchResults, err := factory.FederationRepo.SearchFederation(SearchFederationCriteria{Name: federationName})
		if len(searchResults) == 1 && err == nil {
			event.FederationID = searchResults[0].ID
			event.Federation = searchResults[0]
		} else {
			return *event, errors.New(fmt.Sprintf("cannot find federation \"%s\"", federationName))
		}
	} else {
		return *event, errors.New("federation is required")
	}

	if divisionName != "" {
		searchResults, err := factory.DivisionRepo.SearchDivision(SearchDivisionCriteria{
			Name:         divisionName,
			FederationID: event.FederationID,
		})
		if len(searchResults) == 1 && err == nil {
			event.DivisionID = searchResults[0].ID
			event.Division = searchResults[0]
		} else {
			return *event, errors.New(fmt.Sprintf("cannot find division \"%s\"", divisionName))
		}
	} else {
		return *event, errors.New("division is required")
	}

	if ageName != "" {
		searchResults, err := factory.AgeRepo.SearchAge(SearchAgeCriteria{
			Name:       ageName,
			DivisionID: event.DivisionID,
		})
		if len(searchResults) == 1 && err == nil {
			event.AgeID = searchResults[0].ID
			event.Age = searchResults[0]
		} else {
			return *event, errors.New(fmt.Sprintf("cannot find age \"%s\"", ageName))
		}
	} else {
		return *event, errors.New("age is required")
	}

	if proficiencyName != "" {
		searchResults, err := factory.ProficiencyRepo.SearchProficiency(SearchProficiencyCriteria{
			Name:       proficiencyName,
			DivisionID: event.DivisionID,
		})
		if len(searchResults) == 1 && err == nil {
			event.ProficiencyID = searchResults[0].ID
			event.Proficiency = searchResults[0]
		} else {
			return *event, errors.New(fmt.Sprintf("cannot find proficiency \"%s\"", proficiencyName))
		}
	} else {
		return *event, errors.New("proficiency is required")
	}

	if styleName != "" {
		searchResults, err := factory.StyleRepo.SearchStyle(SearchStyleCriteria{
			Name: styleName,
		})
		if len(searchResults) == 1 && err == nil {
			event.StyleID = searchResults[0].ID
			event.Style = searchResults[0]
		} else {
			return *event, errors.New(fmt.Sprintf("cannot find style \"%s\"", styleName))
		}
	} else {
		return *event, errors.New("style is required")
	}

	if len(danceNames) != 0 {
		for _, each := range danceNames {
			searchResults, err := factory.DanceRepo.SearchDance(SearchDanceCriteria{
				StyleID: event.StyleID,
				Name:    each,
			})
			if len(searchResults) == 1 && err == nil {
				event.AddDance(searchResults[0].ID)
			}
		}
	} else {
		return *event, errors.New("dances are required")
	}
	return *event, nil
}

type EventTemplate struct {
	Federation  string   `json:"federation"`
	Division    string   `json:"division"`
	Age         string   `json:"age"`
	Proficiency string   `json:"proficiency"`
	Style       string   `json:"style"`
	Dances      []string `json:"dances"`
}

type CompetitionEventTemplate struct {
	ID               int
	Name             string
	Description      string
	TargetFederation Federation
	TemplateEvents   []EventTemplate
	DateTimeCreate   time.Time
}

type SearchCompetitionEventTemplateCriteria struct {
	ID           int
	Name         string
	CreateUserID int
}

type ICompetitionEventTemplateRepository interface {
	SearchCompetitionEventTemplates(criteria SearchCompetitionEventTemplateCriteria) ([]CompetitionEventTemplate, error)
}
