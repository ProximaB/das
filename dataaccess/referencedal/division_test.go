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

var divisionRepo = referencedal.PostgresDivisionRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresDivisionRepository_SearchDivision(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows(
		[]string{
			"ID", "NAME", "ABBREVIATION", "DESCRIPTION", "NOTE", "FEDERATION_ID", "DATETIME_CREATED", "DATETIME_UPDATED",
		}).
		AddRow(
			1, "Amateur", "A", "Amateur division of Federation F", "", 3, time.Now(), time.Now()).
		AddRow(
			2, "Professional", "P", "Professional division of Federation F", "", 3, time.Now(), time.Now())

	divisionRepo.Database = db
	mock.ExpectQuery(`SELECT ID, NAME, ABBREVIATION, DESCRIPTION, NOTE, FEDERATION_ID, DATETIME_CREATED, 
DATETIME_UPDATED FROM DAS.DIVISION ORDER BY ID`).WillReturnRows(rows)
	divisions, err := divisionRepo.SearchDivision(referencebll.SearchDivisionCriteria{})

	assert.Nil(t, err, "should search divisions without error")
	assert.EqualValues(t, 2, len(divisions), "should return all divisions when search with empty criteria")
}

func TestPostgresDivisionRepository_CreateDivision(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	divisionRepo.Database = db

	args := referencebll.Division{Name: "Amateur", FederationID: 33}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO DAS.DIVISION (NAME, ABBREVIATION, DESCRIPTION, NOTE,
		FEDERATION_ID, DATETIME_CREATED, DATETIME_UPDATED)`)
	mock.ExpectCommit()
	err = divisionRepo.CreateDivision(&args)

	assert.Nil(t, err, "should insert new division without error")
}

func TestPostgresDivisionRepository_DeleteDivision(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	divisionRepo.Database = db

	args := referencebll.Division{ID: 17, Name: "Professional", FederationID: 12}

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM DAS.DIVISION`).WillReturnResult(sqlmock.NewResult(17, 1))
	mock.ExpectCommit()

	err = divisionRepo.DeleteDivision(args)
	assert.Nil(t, err, "should delete division without error")
}

func TestPostgresDivisionRepository_UpdateDivision(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	divisionRepo.Database = db

	args := referencebll.Division{ID: 17, Name: "Professional", FederationID: 12}

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE DAS.DIVISION`).WillReturnResult(sqlmock.NewResult(17, 1))
	mock.ExpectCommit()

	err = divisionRepo.UpdateDivision(args)
	assert.Nil(t, err, "should delete division without error")

}
