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
	DAS_STYLE_TABLE = "DAS.STYLE"
)

// PostgresStyleRepository implements IStyleRepository and feeds data to
// businesslogic from a PostgreSQL database
type PostgresStyleRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresStyleRepository) CreateStyle(style *referencebll.Style) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresStyleRepository is not specified")
	}
	stmt := repo.SqlBuilder.Insert("").Into(DAS_STYLE_TABLE).Columns(
		common.COL_NAME,
		common.COL_DESCRIPTION,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED,
	).Values(
		style.Name,
		style.Description,
		style.CreateUserID,
		style.DateTimeCreated,
		style.UpdateUserID,
		style.DateTimeUpdated,
	).Suffix(
		"RETURNING ID",
	)

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&style.ID)
		tx.Commit()
	}
	return err
}

func (repo PostgresStyleRepository) DeleteStyle(style referencebll.Style) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresStyleRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Delete("").
		From(DAS_STYLE_TABLE).
		Where(squirrel.Eq{common.PRIMARY_KEY: style.ID})
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		tx.Commit()
	}
	return err
}

func (repo PostgresStyleRepository) SearchStyle(criteria referencebll.SearchStyleCriteria) ([]referencebll.Style, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresStyleRepository is not specified")
	}
	stmt := repo.SqlBuilder.Select(
		fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s",
			common.PRIMARY_KEY,
			common.COL_NAME,
			common.COL_DESCRIPTION,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED)).
		From(DAS_STYLE_TABLE).
		OrderBy(common.PRIMARY_KEY)
	if criteria.StyleID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.PRIMARY_KEY: criteria.StyleID})
	}
	if len(criteria.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_NAME: criteria.Name})
	}
	rows, err := stmt.RunWith(repo.Database).Query()
	styles := make([]referencebll.Style, 0)
	if err != nil {
		return styles, err
	}
	for rows.Next() {
		each := referencebll.Style{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Description,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		styles = append(styles, each)
	}
	rows.Close()
	return styles, err
}

func (repo PostgresStyleRepository) UpdateStyle(style referencebll.Style) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresStyleRepository is not specified")
	}
	stmt := repo.SqlBuilder.Update("").Table(DAS_STYLE_TABLE)
	if style.ID > 0 {
		stmt = stmt.Set(common.COL_NAME, style.Name).
			Set(common.COL_DESCRIPTION, style.Description).
			Set(common.COL_UPDATE_USER_ID, style.UpdateUserID).
			Set(common.COL_DATETIME_UPDATED, style.DateTimeUpdated)
		var err error
		if tx, txErr := repo.Database.Begin(); txErr != nil {
			return txErr
		} else {
			_, err = stmt.RunWith(repo.Database).Exec()
			if err != nil {
				tx.Rollback()
			} else {
				tx.Commit()
			}
		}
		return err
	} else {
		return errors.New("style is not specified")
	}
}
