package viewmodel

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/businesslogic/reference"
	"time"
)

func GetEventDanceIDs(event businesslogic.Event) []int {
	output := make([]int, 0)
	for _, each := range event.GetDances() {
		output = append(output, each)
	}
	return output
}

type CreateEventViewModel struct {
	CompetitionID   int   `json:"competition"`
	EventCategoryID int   `json:"category"`
	FederationID    int   `json:"federation"`
	DivisionID      int   `json:"division"`
	AgeID           int   `json:"age"`
	ProficiencyID   int   `json:"proficiency"`
	StyleID         int   `json:"style"`
	Dances          []int `json:"dances"`
}

func (dto CreateEventViewModel) ToDomainModel(user businesslogic.Account, repo reference.IDanceRepository) businesslogic.Event {
	event := businesslogic.NewEvent()
	event.CompetitionID = dto.CompetitionID
	event.CategoryID = businesslogic.EVENT_CATEGORY_COMPETITIVE_BALLROOM
	event.StatusID = businesslogic.EVENT_STATUS_DRAFT
	event.FederationID = dto.FederationID
	event.DivisionID = dto.DivisionID
	event.AgeID = dto.AgeID
	event.ProficiencyID = dto.ProficiencyID
	event.StyleID = dto.StyleID

	dances := make([]int, 0)
	for _, each := range dto.Dances {
		results, _ := repo.SearchDance(&reference.SearchDanceCriteria{DanceID: each})
		dances = append(dances, results[0].ID)
	}
	event.SetDances(dances)
	event.CreateUserID = user.ID
	event.DateTimeCreated = time.Now()
	event.UpdateUserID = user.ID
	event.DateTimeUpdated = time.Now()

	return event
}
