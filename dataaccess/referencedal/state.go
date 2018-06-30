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
	DAS_STATE_TABLE = "DAS.STATE"
)

type PostgresStateRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresStateRepository) SearchState(criteria referencebll.SearchStateCriteria) ([]referencebll.State, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresStateRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s",
			common.PRIMARY_KEY,
			common.COL_NAME,
			common.COL_ABBREVIATION,
			common.COL_COUNTRY_ID,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED)).
		From(DAS_STATE_TABLE)
	if criteria.CountryID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_COUNTRY_ID: criteria.CountryID})
	}
	if len(criteria.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_NAME: criteria.Name})
	}
	if criteria.StateID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.PRIMARY_KEY: criteria.StateID})
	}
	stmt = stmt.OrderBy(common.PRIMARY_KEY,
		common.COL_NAME)

	states := make([]referencebll.State, 0)
	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return states, err
	}

	for rows.Next() {
		each := referencebll.State{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Abbreviation,
			&each.CountryID,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		states = append(states, each)
	}
	if err != nil {
		return nil, err
	}
	return states, nil
}

func (repo PostgresStateRepository) CreateState(state *referencebll.State) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresStateRepository is not specified")
	}
	stmt := repo.SqlBuilder.Insert("").Into(DAS_STATE_TABLE).Columns(
		common.COL_NAME,
		common.COL_ABBREVIATION,
		common.COL_COUNTRY_ID,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED,
	).Values(
		state.Name,
		state.Abbreviation,
		state.CountryID,
		state.CreateUserID,
		state.DateTimeCreated,
		state.UpdateUserID,
		state.DateTimeUpdated,
	).Suffix(
		"RETURNING ID",
	)

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&state.ID)
		tx.Commit()
	}
	return err
}

func (repo PostgresStateRepository) UpdateState(state referencebll.State) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresStateRepository is not specified")
	}
	stmt := repo.SqlBuilder.Update("").Table(DAS_STATE_TABLE)
	if state.ID > 0 {
		stmt = stmt.Set(common.COL_NAME, state.Name).
			Set(common.COL_ABBREVIATION, state.Abbreviation).
			Set(common.COL_COUNTRY_ID, state.CountryID).
			Set(common.COL_UPDATE_USER_ID, state.UpdateUserID).
			Set(common.COL_DATETIME_UPDATED, state.DateTimeUpdated)

		var err error
		if tx, txErr := repo.Database.Begin(); txErr != nil {
			return txErr
		} else {
			_, err = stmt.RunWith(repo.Database).Exec()
			if err = tx.Commit(); err != nil {
				tx.Rollback()
			}
		}
		return err
	} else {
		return errors.New("state is not specified")
	}
}

func (repo PostgresStateRepository) DeleteState(state referencebll.State) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresStateRepository is not specified")
	}
	stmt := repo.SqlBuilder.Delete("").From(DAS_STATE_TABLE).Where(squirrel.Eq{common.PRIMARY_KEY: state.ID})
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		if err = tx.Commit(); err != nil {
			tx.Rollback()
		}
	}
	return err
}
