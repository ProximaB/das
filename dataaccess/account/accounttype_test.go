// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

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
