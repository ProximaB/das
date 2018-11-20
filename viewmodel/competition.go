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

package viewmodel

import (
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"strings"
	"time"
)

type Competition struct {
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

func CompetitionDataModelToViewModel(competition businesslogic.Competition, accountType int) Competition {
	view := Competition{
		ID:         competition.ID,
		Federation: competition.FederationID,
		Name:       competition.Name,
		Website:    competition.Website,
		CountryID:  competition.Country.ID,
		StateID:    competition.GetStatus(),
		CityID:     competition.City.ID,
		Address:    competition.Street,
		StartDate:  competition.StartDateTime,
		EndDate:    competition.EndDateTime,
		Attendance: competition.Attendance,
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
	FederationID   int             `json:"federationId"`
	Name           string          `json:"name"`
	Start          CompetitionDate `json:"start"`
	End            CompetitionDate `json:"end"`
	Status         int             `json:"statusId"`
	WebsiteUrl     string          `json:"website"`
	VenueStreet    string          `json:"address"`
	VenueCityID    int             `json:"cityId"`
	VenueStateID   int             `json:"stateId"`
	VenueCountryID int             `json:"countryId"`
	ContactName    string          `json:"contact"`
	ContactPhone   string          `json:"phone"`
	ContactEmail   string          `json:"email"`
	CreateUserID   string          `json:"createdby,omitempty"`
}

func (createDTO CreateCompetition) ToCompetitionDataModel(user businesslogic.Account) businesslogic.Competition {

	competition := businesslogic.Competition{
		FederationID: createDTO.FederationID,
		Name:         createDTO.Name,
		Website:      createDTO.WebsiteUrl,

		Country: reference.Country{},
		State:   reference.State{},
		City:    reference.City{},
		Street:  createDTO.VenueStreet,

		ContactName:  createDTO.ContactName,
		ContactPhone: createDTO.ContactPhone,
		ContactEmail: createDTO.ContactEmail,

		StartDateTime: createDTO.Start.Time,
		EndDateTime:   createDTO.End.Time,

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
