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

var styleRepository = referencedal.PostgresStyleRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresStyleRepository_SearchStyle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	styleRepository.Database = db

	rows := sqlmock.NewRows(
		[]string{
			"ID", "NAME", "DESCRIPTION", "CREATE_USER_ID", "DATETIME_CREATED", "UPDATE_USER_ID", "DATETIME_UPDATED",
		},
	).AddRow(1, "Standard", "International Standard", 3, time.Now(), 4, time.Now())

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	styles, _ := styleRepository.SearchStyle(reference.SearchStyleCriteria{})
	assert.NotZero(t, len(styles))
}

func TestPostgresStyleRepository_CreateStyle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	styleRepository.Database = db

	args := reference.Style{Name: "Rhythm", DateTimeCreated: time.Now(), DateTimeUpdated: time.Now()}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO DAS.STYLE (NAME, DESCRIPTION, CREATE_USER_ID, 
		DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED)`)
	mock.ExpectCommit()
	err = styleRepository.CreateStyle(&args)

	assert.Nil(t, err, "should insert a new style without error")
}

func TestPostgresStyleRepository_DeleteStyle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	styleRepository.Database = db

	args := reference.Style{ID: 2, Name: "Rhythm"}

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM DAS.STYLE WHERE`).WillReturnResult(sqlmock.NewResult(2, 1))
	mock.ExpectCommit()

	err = styleRepository.DeleteStyle(args)
	assert.Nil(t, err, "should delete style without error")
}

func TestPostgresStyleRepository_UpdateStyle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	styleRepository.Database = db

	args := reference.Style{ID: 2, Name: "Rhythm"}

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE DAS.STYLE`).WillReturnResult(sqlmock.NewResult(2, 1))
	mock.ExpectCommit()

	err = styleRepository.UpdateStyle(args)
	assert.Nil(t, err, "should update style with error")
}
