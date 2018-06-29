// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package account

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var accountRepository = PostgresAccountRepository{
	Database:   nil,
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresAccountRepository_SearchAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	accountRepository.Database = db
	rows := sqlmock.NewRows(
		[]string{
			"ID",
			"NAME",
			"ABBREVIATION",
			"YEAR_FOUNDED",
			"COUNTRY_ID",
			"CREATE_USER_ID",
			"DATETIME_CREATED",
			"UPDATE_USER_ID",
			"DATETIME_UPDATED",
		},
	).AddRow(1,
		"Pre Teen I",
		"Pre",
		1233,
		8,
		2,
		time.Now(),
		3,
		time.Now())
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	results, err := accountRepository.SearchAccount(businesslogic.SearchAccountCriteria{
		Email: "test",
	})

	assert.Nil(t, err, "Database schema should match")
	assert.NotNil(t, results, "should return at least empty slice of accounts")
}
