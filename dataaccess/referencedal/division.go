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
	DAS_DIVISION_TABLE = "DAS.DIVISION"
)

type PostgresDivisionRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresDivisionRepository) SearchDivision(criteria referencebll.SearchDivisionCriteria) ([]referencebll.Division, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresDivisionRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s",
			common.PRIMARY_KEY,
			common.COL_NAME,
			common.COL_ABBREVIATION,
			common.COL_DESCRIPTION,
			common.COL_NOTE,
			common.COL_FEDERATION_ID,
			common.COL_DATETIME_CREATED,
			common.COL_DATETIME_UPDATED)).
		From(DAS_DIVISION_TABLE).
		OrderBy(common.PRIMARY_KEY)
	if criteria.FederationID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_FEDERATION_ID: criteria.FederationID})
	}
	if criteria.ID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.PRIMARY_KEY: criteria.ID})
	}
	rows, err := stmt.RunWith(repo.Database).Query()
	divisions := make([]referencebll.Division, 0)
	if err != nil {
		return divisions, err
	}
	for rows.Next() {
		each := referencebll.Division{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Abbreviation,
			&each.Description,
			&each.Note,
			&each.FederationID,
			&each.DateTimeCreated,
			&each.DateTimeUpdated,
		)
		divisions = append(divisions, each)
	}
	rows.Close()
	return divisions, err
}

func (repo PostgresDivisionRepository) CreateDivision(division *referencebll.Division) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresDivisionRepository is not specified")
	}
	stmt := repo.SqlBuilder.Insert("").Into(DAS_DIVISION_TABLE).Columns(
		common.COL_NAME,
		common.COL_ABBREVIATION,
		common.COL_DESCRIPTION,
		common.COL_NOTE,
		common.COL_FEDERATION_ID,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED,
	).Values(
		division.Name,
		division.Abbreviation,
		division.Description,
		division.Note,
		division.FederationID,
		division.CreateUserID,
		division.DateTimeCreated,
		division.UpdateUserID,
		division.DateTimeUpdated,
	).Suffix(
		"RETURNING ID",
	)

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&division.ID)
		tx.Commit()
	}
	return err
}

func (repo PostgresDivisionRepository) UpdateDivision(division referencebll.Division) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresDivisionRepository is not specified")
	}
	stmt := repo.SqlBuilder.Update("").Table(DAS_DIVISION_TABLE)
	if division.ID > 0 {
		stmt = stmt.Set(common.COL_NAME, division.Name).
			Set(common.COL_ABBREVIATION, division.Abbreviation).
			Set(common.COL_DESCRIPTION, division.Description).
			Set(common.COL_NOTE, division.Note).
			Set(common.COL_FEDERATION_ID, division.FederationID).
			Set(common.COL_UPDATE_USER_ID, division.UpdateUserID).
			Set(common.COL_DATETIME_UPDATED, division.DateTimeUpdated)

		var err error
		if tx, txErr := repo.Database.Begin(); txErr != nil {
			return txErr
		} else {
			_, err = stmt.RunWith(repo.Database).Exec()
			tx.Commit()
		}
		return err
	} else {
		return errors.New("division is not specified")
	}
}

func (repo PostgresDivisionRepository) DeleteDivision(division referencebll.Division) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresDivisionRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Delete("").
		From(DAS_DIVISION_TABLE).
		Where(squirrel.Eq{common.PRIMARY_KEY: division.ID})
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		tx.Commit()
	}
	return err
}
