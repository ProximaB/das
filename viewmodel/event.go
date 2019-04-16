package viewmodel

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"time"
)

// OrganizerSearchEventCriteria defines the query string that Organizer can submit to search
// events that the organizer created
type OrganizerSearchEventCriteria struct {
	FederationID  int `schema:"federationId"`
	CompetitionID int `schema:"competitionId"`
	EventID       int `schema:"eventId"`
	DivisionID    int `schema:"divisionId"`
	AgeID         int `schema:"ageId"`
	ProficiencyID int `schema:"proficiencyId"`
	StyleID       int `schema:"styleId"`
	OrganizerID   int `schema:"organizerId,omitempty"`
}

func (criteria OrganizerSearchEventCriteria) ToBusinessModel() businesslogic.SearchEventCriteria {
	return businesslogic.SearchEventCriteria{
		CompetitionID: criteria.CompetitionID,
		EventID:       criteria.EventID,
		FederationID:  criteria.FederationID,
		DivisionID:    criteria.DivisionID,
		AgeID:         criteria.AgeID,
		ProficiencyID: criteria.ProficiencyID,
		StyleID:       criteria.StyleID,
		OrganizerID:   criteria.OrganizerID,
	}
}

type SearchCompetitionEventTemplateForm struct {
	ID         int    `schema:"templateId"`
	Name       string `schema:"name"`
	Federation string `schema:"federation"`
	OwnerID    int
}

// CreateEventForm defines the payload for creating an Event
type CreateEventForm struct {
	CompetitionID   int   `json:"competition" validate:"min=1"`
	EventCategoryID int   `json:"category" validate:"min=1"`
	FederationID    int   `json:"federation" validate:"min=1"`
	DivisionID      int   `json:"division" validate:"min=1"`
	AgeID           int   `json:"age" validate:"min=1"`
	ProficiencyID   int   `json:"proficiency" validate:"min=1"`
	StyleID         int   `json:"style" validate:"min=1"`
	Dances          []int `json:"dances"`
	Template        int   `json:"template,omitempty"`
}

// ToDomainModel converts the caller CreateEventForm to the Event domain model
func (dto CreateEventForm) ToDomainModel(user businesslogic.Account) *businesslogic.Event {
	event := businesslogic.NewEvent()
	event.CompetitionID = dto.CompetitionID
	event.CategoryID = businesslogic.EventCategoryCompetitiveBallroom
	event.StatusID = businesslogic.EVENT_STATUS_DRAFT
	event.FederationID = dto.FederationID
	event.DivisionID = dto.DivisionID
	event.AgeID = dto.AgeID
	event.ProficiencyID = dto.ProficiencyID
	event.StyleID = dto.StyleID

	dances := make([]int, 0)
	for _, each := range dto.Dances {
		dances = append(dances, each)
	}

	event.SetDances(dances)
	event.CreateUserID = user.ID
	event.DateTimeCreated = time.Now()
	event.UpdateUserID = user.ID
	event.DateTimeUpdated = time.Now()

	return event
}

// EventDanceViewModel defines the JSON structure of EventDance
type EventDanceViewModel struct {
	ID      int `json:"eventDanceId"`
	EventId int `json:"eventId"`
	DanceId int `json:"danceId"`
}

func (view *EventDanceViewModel) PopulateViewModel(model businesslogic.EventDance) {
	view.ID = model.ID
	view.EventId = model.EventID
	view.DanceId = model.DanceID
}

// EventViewModel defines the JSON structure of Event which is used in outbound API
type EventViewModel struct {
	ID            int                   `json:"eventId"`
	CompetitionID int                   `json:"competitionId"`
	FederationID  int                   `json:"federationId"`
	DivisionID    int                   `json:"divisionId"`
	AgeID         int                   `json:"ageId"`
	ProficiencyID int                   `json:"proficiencyId"`
	StyleID       int                   `json:"styleId"`
	Dances        []int                 `json:"dances"`
	EventDances   []EventDanceViewModel `json:"eventDances"`
}

// PopulateViewModel populates the caller EventViewModel data fields with data from business logic Event
func (view *EventViewModel) PopulateViewModel(model businesslogic.Event) {
	view.ID = model.ID
	view.CompetitionID = model.CompetitionID
	view.FederationID = model.FederationID
	view.DivisionID = model.DivisionID
	view.AgeID = model.AgeID
	view.ProficiencyID = model.ProficiencyID
	view.StyleID = model.StyleID
	view.Dances = model.GetDances()
	view.EventDances = make([]EventDanceViewModel, 0)
	for _, each := range model.GetEventDances() {
		item := EventDanceViewModel{}
		item.PopulateViewModel(each)
		view.EventDances = append(view.EventDances, item)
	}
}

type DeleteEventForm struct {
	ID int `json:"eventId";validate:"min=1"`
}
