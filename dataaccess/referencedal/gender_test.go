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
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/DancesportSoftware/das/dataaccess/referencedal"
	"github.com/Masterminds/squirrel"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var genderRepository = referencedal.PostgresGenderRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresGenderRepository_GetAllGenders(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows(
		[]string{
			common.ColumnPrimaryKey,
			common.COL_NAME,
			common.ColumnAbbreviation,
			common.COL_DESCRIPTION,
			common.COL_DATETIME_CREATED,
			common.COL_DATETIME_UPDATED,
		},
	).AddRow(
		1, "Female", "F", "Biologicially female", time.Now(), time.Now(),
	).AddRow(
		2, "Male", "M", "Biologically male", time.Now(), time.Now(),
	)

	mock.ExpectQuery(`SELECT ID, NAME, ABBREVIATION, DESCRIPTION, DATETIME_CREATED, DATETIME_UPDATED
FROM DAS.GENDER`).WillReturnRows(rows)
	genderRepository.Database = db
	genderRepository.GetAllGenders()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s\n", err)
	}
}
