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

package referencedal_test

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/referencedal"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var cityRepository = referencedal.PostgresCityRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

var city = businesslogic.City{
	Name:            "Test City",
	StateID:         1,
	DateTimeCreated: time.Now(),
	DateTimeUpdated: time.Now(),
}

func TestPostgresCityRepository_SearchCity(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	cityRepository.Database = db

	rows := sqlmock.NewRows(
		[]string{"ID", "NAME", "STATE_ID", "CREATE_USER_ID", "DATETIME_CREATED", "UPDATE_USER_ID", "DATETIME_UPDATED"},
	).AddRow(
		1, "Madison", 2, 3, time.Now(), 3, time.Now(),
	).AddRow(
		2, "Milwaukee", 4, 5, time.Now(), 4, time.Now(),
	)

	mock.ExpectQuery(`SELECT ID, NAME, STATE_ID, CREATE_USER_ID, DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED FROM DAS.CITY`).WillReturnRows(rows)
	cities, err := cityRepository.SearchCity(businesslogic.SearchCityCriteria{})

	assert.NotZero(t, len(cities), "should retrieve cities that were populated to Database")
	assert.Nil(t, err, "schema for DAS.CITY should be up to date")
}

func TestPostgresCityRepository_CreateCity(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	cityRepository.Database = db

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO DAS.CITY (NAME, STATE_ID, CREATE_USER_ID, DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED) `)
	mock.ExpectCommit()
	err = cityRepository.CreateCity(&city)

	assert.Nil(t, err, "should be able to create a new city")
}

func TestPostgresCityRepository_DeleteCity(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	cityRepository.Database = db

	mock.ExpectBegin()
	mock.ExpectExec(`^DELETE FROM DAS.CITY`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = cityRepository.DeleteCity(businesslogic.City{ID: 1, Name: "Shenzhen"})

	assert.Nil(t, err, "should delete city without error")
}

func TestPostgresCityRepository_UpdateCity(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	cityRepository.Database = db

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE DAS.CITY").WillReturnResult(sqlmock.NewResult(12, 1))
	mock.ExpectCommit()

	args := businesslogic.City{ID: 12, Name: "New City", StateID: 77}

	err = cityRepository.UpdateCity(args)

	assert.Nil(t, err, "should update city without error")

}
