package businesslogic

import (
	"errors"
	"fmt"
	"time"
)

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
