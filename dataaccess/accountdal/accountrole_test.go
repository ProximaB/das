// Dancesport Application System (DAS)
// Copyright (C) 2018 Yubing Hou
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
	"time"
)

var accountRoleRepo = accountdal.PostgresAccountRoleRepository{
	Database:   nil,
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresAccountRoleRepository_CreateAccountRole(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	role := businesslogic.AccountRole{
		AccountTypeID:   businesslogic.AccountTypeAthlete,
		AccountID:       12,
		CreateUserID:    12,
		DateTimeCreated: time.Now(),
		UpdateUserID:    12,
		DateTimeUpdated: time.Now(),
	}

	missingDBErr := accountRoleRepo.CreateAccountRole(&role)
	assert.NotNil(t, missingDBErr, "should return an error if database connection is not specified")

	accountRoleRepo.Database = db
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO DAS.ACCOUNT_ROLE (ACCOUNT_ID, ACCOUNT_TYPE_ID, CREATE_USER_ID, DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED) VALUES`)
	mock.ExpectCommit()

	createErr := accountRoleRepo.CreateAccountRole(&role)
	assert.Nil(t, createErr, "should insert legitimate AccountRole data without error")
}

func TestPostgresAccountRoleRepository_SearchAccountRole(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	accountRoleRepo.Database = db

	rows1 := sqlmock.NewRows(
		[]string{
			"ID", "ACCOUNT_ID", "ACCOUNT_TYPE_ID", "CREATE_USER_ID", "DATETIME_CREATED", "UPDATE_USER_ID", "DATETIME_UPDATED",
		},
	).AddRow(
		1, 2, 4, 5, time.Now(), 5, time.Now(),
	).AddRow(
		2, 2, 6, 5, time.Now(), 5, time.Now(),
	).AddRow(
		3, 3, 4, 5, time.Now(), 5, time.Now(),
	)

	rows2 := sqlmock.NewRows(
		[]string{
			"ID", "ACCOUNT_ID", "ACCOUNT_TYPE_ID", "CREATE_USER_ID", "DATETIME_CREATED", "UPDATE_USER_ID", "DATETIME_UPDATED",
		},
	).AddRow(
		1, 2, 4, 5, time.Now(), 5, time.Now(),
	).AddRow(
		2, 2, 6, 5, time.Now(), 5, time.Now(),
	)

	rows3 := sqlmock.NewRows(
		[]string{
			"ID", "ACCOUNT_ID", "ACCOUNT_TYPE_ID", "CREATE_USER_ID", "DATETIME_CREATED", "UPDATE_USER_ID", "DATETIME_UPDATED",
		},
	).AddRow(
		1, 2, 4, 5, time.Now(), 5, time.Now(),
	).AddRow(
		3, 3, 4, 5, time.Now(), 5, time.Now(),
	)

	mock.ExpectQuery(`SELECT ID, ACCOUNT_ID, ACCOUNT_TYPE_ID, CREATE_USER_ID, DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED FROM DAS.ACCOUNT_ROLE`).WillReturnRows(rows1)
	mock.ExpectQuery(`SELECT ID, ACCOUNT_ID, ACCOUNT_TYPE_ID, CREATE_USER_ID, DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED FROM DAS.ACCOUNT_ROLE WHERE ACCOUNT_ID = `).WillReturnRows(rows2)
	mock.ExpectQuery(`SELECT ID, ACCOUNT_ID, ACCOUNT_TYPE_ID, CREATE_USER_ID, DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED FROM DAS.ACCOUNT_ROLE WHERE ACCOUNT_TYPE_ID = `).WillReturnRows(rows3)

	res1, err1 := accountRoleRepo.SearchAccountRole(businesslogic.SearchAccountRoleCriteria{})
	res2, err2 := accountRoleRepo.SearchAccountRole(businesslogic.SearchAccountRoleCriteria{AccountID: 2})
	res3, err3 := accountRoleRepo.SearchAccountRole(businesslogic.SearchAccountRoleCriteria{AccountTypeID: 4})

	assert.Equal(t, 3, len(res1), "should not return empty roles when there are valid records")
	assert.Equal(t, 2, len(res2), "should not return empty roles when there are valid records")
	assert.Equal(t, 2, len(res3), "should not return empty roles when there are valid records")

	assert.Nil(t, err1, "should not return error when criteria data is valid")
	assert.Nil(t, err2, "should not return error when criteria data is valid")
	assert.Nil(t, err3, "should not return error when criteria data is valid")
}
