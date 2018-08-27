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

package referencedal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/Masterminds/squirrel"
)

const (
	DAS_AGE_TABLE           = "DAS.AGE"
	DAS_AGE_COL_ENFORCED    = "ENFORCED"
	DAS_AGE_COL_MINIMUM_AGE = "MINIMUM_AGE"
	DAS_AGE_COL_MAXIMUM_AGE = "MAXIMUM_AGE"
)

type PostgresAgeRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresAgeRepository) CreateAge(age *reference.Age) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresAgeRepository is not specified")
	}
	stmt := repo.SqlBuilder.Insert("").Into(DAS_AGE_TABLE).Columns(
		common.COL_NAME,
		common.COL_DESCRIPTION,
		common.COL_DIVISION_ID,
		DAS_AGE_COL_ENFORCED,
		DAS_AGE_COL_MINIMUM_AGE,
		DAS_AGE_COL_MAXIMUM_AGE,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated,
	).Values(
		age.Name,
		age.Description,
		age.DivisionID,
		age.Enforced,
		age.AgeMinimum,
		age.AgeMaximum,
		age.CreateUserID,
		age.DateTimeCreated,
		age.UpdateUserID,
		age.DateTimeUpdated,
	).Suffix("RETURNING ID")
	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&age.ID)
		tx.Commit()
	}
	return err
}

func (repo PostgresAgeRepository) DeleteAge(age reference.Age) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresAgeRepository is not specified")
	}
	stmt := repo.SqlBuilder.Delete("").From(DAS_AGE_TABLE).
		Where(squirrel.Eq{common.ColumnPrimaryKey: age.ID})
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		tx.Commit()
	}
	return err
}

func (repo PostgresAgeRepository) SearchAge(criteria reference.SearchAgeCriteria) ([]reference.Age, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresAgeRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
			common.ColumnPrimaryKey,
			common.COL_NAME,
			common.COL_DESCRIPTION,
			common.COL_DIVISION_ID,
			DAS_AGE_COL_ENFORCED,
			DAS_AGE_COL_MINIMUM_AGE,
			DAS_AGE_COL_MAXIMUM_AGE,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated)).
		From(DAS_AGE_TABLE).
		OrderBy(common.ColumnPrimaryKey)
	if criteria.DivisionID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_DIVISION_ID: criteria.DivisionID})
	}
	if criteria.AgeID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.AgeID})
	}
	rows, err := stmt.RunWith(repo.Database).Query()
	output := make([]reference.Age, 0)
	if err != nil {
		return output, err
	}
	for rows.Next() {
		age := reference.Age{}
		rows.Scan(
			&age.ID,
			&age.Name,
			&age.Description,
			&age.DivisionID,
			&age.Enforced,
			&age.AgeMinimum,
			&age.AgeMaximum,
			&age.CreateUserID,
			&age.DateTimeCreated,
			&age.UpdateUserID,
			&age.DateTimeUpdated,
		)
		output = append(output, age)
	}
	rows.Close()
	return output, err
}

func (repo PostgresAgeRepository) UpdateAge(age reference.Age) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresAgeRepository is not specified")
	}
	stmt := repo.SqlBuilder.Update("").Table(DAS_AGE_TABLE)
	if age.ID > 0 {
		stmt = stmt.Set(common.COL_NAME, age.Name).
			Set(common.COL_DESCRIPTION, age.Description).
			Set(common.COL_DIVISION_ID, age.DivisionID).
			Set(DAS_AGE_COL_MINIMUM_AGE, age.AgeMinimum).
			Set(DAS_AGE_COL_MAXIMUM_AGE, age.AgeMaximum).
			Set(DAS_AGE_COL_ENFORCED, age.Enforced).
			Set(common.ColumnUpdateUserID, age.UpdateUserID).
			Set(common.ColumnDateTimeUpdated, age.DateTimeUpdated)
	}
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		tx.Commit()
	}
	return err
}
