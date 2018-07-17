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

package accountdal_test

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/accountdal"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

var accountRepository = accountdal.PostgresAccountRepository{
	Database:   nil,
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresAccountRepository_SearchAccount(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	res1, err := accountRepository.SearchAccount(businesslogic.SearchAccountCriteria{})
	assert.NotNil(t, err, "should return an error when database connection is not specified")
	assert.Nil(t, res1, "should not return a concrete object if database connection does not even exist")

	accountRepository.Database = db
	rows := sqlmock.NewRows(
		[]string{
			"ID",
			"UUID",
			"ACCOUNT_STATUS_ID",
			"USER_GENDER_ID",
			"LAST_NAME",
			"MIDDLE_NAMES",
			"FIRST_NAME",
			"DATE_OF_BIRTH",
			"EMAIl",
			"PHONE",
			"EMAIL_VERIFIED",
			"PHONE_VERIFIED",
			"HASH_ALGORITHM",
			"PASSWORD_SALT",
			"PASSWORD_HASH",
			"DATETIME_CREATED",
			"DATETIME_UPDATED",
			"TOS_ACCEPTED",
			"PP_ACCEPTED",
			"BY_GUARDIAN",
			"GUARDIAN_SIGNATURE",
		},
	)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	results, err := accountRepository.SearchAccount(businesslogic.SearchAccountCriteria{
		Email: "test",
	})

	assert.Nil(t, err, "Database schema should match")
	assert.NotNil(t, results, "should return at least empty slice of accounts")
}

func TestPostgresAccountRepository_CreateAccount(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	accountRepository.Database = db
	rows := sqlmock.NewRows(
		[]string{
			"ID",
			"UUID",
			"ACCOUNT_STATUS_ID",
			"USER_GENDER_ID",
			"LAST_NAME",
			"MIDDLE_NAMES",
			"FIRST_NAME",
			"DATE_OF_BIRTH",
			"EMAIl",
			"PHONE",
			"EMAIL_VERIFIED",
			"PHONE_VERIFIED",
			"HASH_ALGORITHM",
			"PASSWORD_SALT",
			"PASSWORD_HASH",
			"DATETIME_CREATED",
			"DATETIME_UPDATED",
			"TOS_ACCEPTED",
			"PP_ACCEPTED",
			"BY_GUARDIAN",
			"GUARDIAN_SIGNATURE",
		},
	)

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	account := businesslogic.Account{}
	results := accountRepository.CreateAccount(&account)

	assert.Nil(t, results)
	assert.NotEqual(t, account.ID, 0)
}
