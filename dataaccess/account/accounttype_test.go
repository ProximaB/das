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

package account

import (
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var accountTypeRepository = PostgresAccountTypeRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresAccountTypeRepository_GetAccountTypes(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	accountTypeRepository.Database = db
	rows := sqlmock.NewRows(
		[]string{
			"ID",
			"NAME",
			"DESCRIPTION",
			"DATETIME_CREATED",
			"DATETIME_UPDATED",
		},
	).AddRow(
		1, "Athlete", "Athlete or Competitor", time.Now(), time.Now(),
	).AddRow(
		2, "Adjudicator", "Judges", time.Now(), time.Now(),
	).AddRow(
		3, "Scrutineer", "Scrutineer or chairperson of judge", time.Now(), time.Now(),
	).AddRow(
		4, "Organizer", "Competition organizer", time.Now(), time.Now(),
	).AddRow(
		5, "Deck Captain", "Deck Captain", time.Now(), time.Now(),
	).AddRow(
		6, "Emcee", "Emcee view", time.Now(), time.Now(),
	)
	mock.ExpectQuery(`SELECT ID, NAME, DESCRIPTION, DATETIME_CREATED, DATETIME_UPDATED FROM 
			DAS.ACCOUNT_TYPE`).WillReturnRows(rows)
	types, _ := accountTypeRepository.GetAccountTypes()
	assert.EqualValues(t, 6, len(types))
}
