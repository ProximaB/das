package viewmodel

import (
	"fmt"
	"github.com/ProximaB/das/businesslogic"
	"strings"
	"time"
)

type CompetitionViewModel struct {
	ID           int       `json:"id"`
	Federation   int       `json:"federationId"`
	Name         string    `json:"name"`
	Website      string    `json:"website"`
	Status       int       `json:"statusId"`
	CountryID    int       `json:"countryId"`
	StateID      int       `json:"stateId"`
	CityID       int       `json:"cityId"`
	Address      string    `json:"address"`
	StartDate    time.Time `json:"start"`
	EndDate      time.Time `json:"end"`
	Attendance   int       `json:"attendance"`
	ContactName  string    `json:"contact"`
	ContactPhone string    `json:"phone"`
	ContactEmail string    `json:"email"`
}

func CompetitionDataModelToViewModel(competition businesslogic.Competition, accountType int) CompetitionViewModel {
	view := CompetitionViewModel{
		ID:         competition.ID,
		Federation: competition.FederationID,
		Name:       competition.Name,
		Website:    competition.Website,
		CountryID:  competition.Country.ID,
		StateID:    competition.State.ID,
		CityID:     competition.City.ID,
		Address:    competition.Street,
		StartDate:  competition.StartDateTime,
		EndDate:    competition.EndDateTime,
		Attendance: competition.Attendance,
		Status:     competition.GetStatus(),
	}

	if accountType == businesslogic.AccountTypeOrganizer {
		view.ContactName = competition.ContactName
		view.ContactEmail = competition.ContactEmail
		view.ContactPhone = competition.ContactPhone
	}
	return view
}

type CompetitionDate struct {
	time.Time
}

func (cd *CompetitionDate) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.Trim(strInput, "\"")
	newTime, err := time.Parse("2006-01-02", strInput)
	if err != nil {
		return err
	}
	cd.Time = newTime
	return nil
}
func (cd *CompetitionDate) MarshalJSON() ([]byte, error) {
	if cd.Time.UnixNano() == (time.Time{}).UnixNano() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", cd.Time.Format("2006-01-02"))), nil
}

// CreateCompetition defines the JSON payload for creating a competition
type CreateCompetition struct {
	FederationID   int       `json:"federationId" validate:"min=1"`
	Name           string    `json:"name" validate:"min=3"`
	Start          time.Time `json:"start"`
	End            time.Time `json:"end"`
	Status         int       `json:"statusId"`
	Website        string    `json:"website" validate:"min=10"` // TODO: still requires regex check
	VenueStreet    string    `json:"address" validate:"min=4"`
	VenueCityID    int       `json:"cityId" validate:"min=1"`
	VenueStateID   int       `json:"stateId" validate:"min=1"`
	VenueCountryID int       `json:"countryId" validate:"min=1"`
	ContactName    string    `json:"contact" validate:"min=3"`
	ContactPhone   string    `json:"phone" validate:"min=5"`
	ContactEmail   string    `json:"email" validate:"min=9"`
	CreateUserID   string    `json:"createdby,omitempty"`
}

func (createDTO CreateCompetition) ToCompetitionDataModel(user businesslogic.Account) businesslogic.Competition {

	competition := businesslogic.Competition{
		FederationID: createDTO.FederationID,
		Name:         createDTO.Name,
		Website:      createDTO.Website,

		Country: businesslogic.Country{},
		State:   businesslogic.State{},
		City:    businesslogic.City{},
		Street:  createDTO.VenueStreet,

		ContactName:  createDTO.ContactName,
		ContactPhone: createDTO.ContactPhone,
		ContactEmail: createDTO.ContactEmail,

		StartDateTime: createDTO.Start,
		EndDateTime:   createDTO.End,

		CreateUserID:    user.ID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    user.ID,
		DateTimeUpdated: time.Now(),
	}
	competition.Country.ID = createDTO.VenueCountryID
	competition.State.ID = createDTO.VenueStateID
	competition.City.ID = createDTO.VenueCityID
	competition.UpdateStatus(createDTO.Status)
	return competition
}

type UpdateCompetitionDTO struct {
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
