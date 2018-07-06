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
