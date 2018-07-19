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
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"time"
)

type Competition struct {
	ID           int       `json:"id"`
	Federation   int       `json:"federation"`
	Name         string    `json:"name"`
	Website      string    `json:"website"`
	Status       int       `json:"status"`
	CountryID    int       `json:"country"`
	StateID      int       `json:"state"`
	CityID       int       `json:"city"`
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

type CreateCompetition struct {
	FederationID   int       `json:"federation"`
	Name           string    `json:"name"`
	Start          time.Time `json:"start"`
	End            time.Time `json:"end"`
	Status         int       `json:"status"`
	WebsiteUrl     string    `json:"website"`
	VenueStreet    string    `json:"address"`
	VenueCityID    int       `json:"city"`
	VenueStateID   int       `json:"state"`
	VenueCountryID int       `json:"country"`
	ContactName    string    `json:"contact"`
	ContactPhone   string    `json:"phone"`
	ContactEmail   string    `json:"email"`
	CreateUserID   string    `json:"createdby,omitempty"`
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
