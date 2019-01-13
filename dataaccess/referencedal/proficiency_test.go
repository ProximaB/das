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

var proficiencyRepo = referencedal.PostgresProficiencyRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresProficiencyRepository_SearchProficiency(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	proficiencyRepo.Database = db

	rows := sqlmock.NewRows(
		[]string{
			"ID", "NAME", "DIVISION_ID", "DESCRIPTION", "CREATE_USER_ID", "DATETIME_CREATED", "UPDATE_USER_ID", "DATETIME_UPDATED",
		},
	).AddRow(1, "Gold", 3, "USA DANCE Gold", 3, time.Now(), 4, time.Now())

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	proficienies, _ := proficiencyRepo.SearchProficiency(businesslogic.SearchProficiencyCriteria{})

	assert.EqualValues(t, 1, len(proficienies))
}

func TestPostgresProficiencyRepository_CreateProficiency(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	proficiencyRepo.Database = db

	args := businesslogic.Proficiency{Name: "Gold"}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO DAS.PROFICIENCY (NAME, DIVISION_ID, DESCRIPTION, CREATE_USER_ID,
		DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED)`)
	mock.ExpectCommit()

	err = proficiencyRepo.CreateProficiency(&args)
	assert.Nil(t, err, "should insert new proficiency without error")
}

func TestPostgresProficiencyRepository_DeleteProficiency(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	proficiencyRepo.Database = db

	args := businesslogic.Proficiency{ID: 12, Name: "Gold"}

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM DAS.PROFICIENCY`).WillReturnResult(sqlmock.NewResult(12, 1))
	mock.ExpectCommit()

	err = proficiencyRepo.DeleteProficiency(args)
	assert.Nil(t, err, "should delete proficiency without error")
}

func TestPostgresProficiencyRepository_UpdateProficiency(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	proficiencyRepo.Database = db

	args := businesslogic.Proficiency{ID: 12, Name: "Gold"}

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE DAS.PROFICIENCY`).WillReturnResult(sqlmock.NewResult(12, 1))
	mock.ExpectCommit()

	err = proficiencyRepo.UpdateProficiency(args)
	assert.Nil(t, err, "should update proficiency without error")
}
