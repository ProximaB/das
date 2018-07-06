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
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/dataaccess/referencedal"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var studioRepository = referencedal.PostgresStudioRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresStudioRepository_SearchStudio(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	studioRepository.Database = db

	rows := sqlmock.NewRows(
		[]string{
			"ID", "NAME", "ABBREVIATION", "COUNTRY_ID", "CREATE_USER_ID", "DATETIME_CREATED", "UPDATE_USER_ID", "DATETIME_UPDATED",
		},
	).AddRow(1, "Wisconsin", "WI", 8, 3, time.Now(), 4, time.Now())

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	studios, _ := studioRepository.SearchStudio(referencebll.SearchStudioCriteria{})

	assert.NotZero(t, len(studios))
}

func TestPostgresStudioRepository_CreateStudio(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	studioRepository.Database = db

	args := referencebll.Studio{Name: "Super Dance Studio", CityID: 34}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO DAS.STUDIO (NAME, ADDRESS, CITY_ID, WEBSITE, CREATE_USER_ID, DATETIME_CREATED, 
		UPDATE_USER_ID, DATETIME_UPDATED)`)
	mock.ExpectCommit()

	err = studioRepository.CreateStudio(&args)
	assert.Nil(t, err, "should insert new studio without error")
}

func TestPostgresStudioRepository_DeleteStudio(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	studioRepository.Database = db

	args := referencebll.Studio{ID: 29, Name: "Super Dance Studio", CityID: 34}

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM DAS.STUDIO`).WillReturnResult(sqlmock.NewResult(29, 1))
	mock.ExpectCommit()

	err = studioRepository.DeleteStudio(args)
	assert.Nil(t, err, "should delete new studio without error")
}

func TestPostgresStudioRepository_UpdateStudio(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	studioRepository.Database = db

	args := referencebll.Studio{ID: 29, Name: "Super Dance Studio", CityID: 34}

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE DAS.STUDIO`).WillReturnResult(sqlmock.NewResult(29, 1))
	mock.ExpectCommit()

	err = studioRepository.UpdateStudio(args)
	assert.Nil(t, err, "should update new studio without error")
}
