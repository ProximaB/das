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
	DAS_STUDIO_TABLE = "DAS.STUDIO"
)

type PostgresStudioRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresStudioRepository) SearchStudio(criteria reference.SearchStudioCriteria) ([]reference.Studio, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresStudioRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf(`DAS.STUDIO.%s, DAS.STUDIO.%s, DAS.STUDIO.%s, DAS.STUDIO.%s, 
		DAS.STUDIO.%s, DAS.STUDIO.%s, DAS.STUDIO.%s, DAS.STUDIO.%s, DAS.STUDIO.%s`,
			common.ColumnPrimaryKey,
			common.COL_NAME,
			common.COL_ADDRESS,
			common.COL_CITY_ID,
			common.COL_WEBSITE,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated)).
		From(DAS_STUDIO_TABLE).OrderBy(common.ColumnPrimaryKey)
	if len(criteria.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_NAME: criteria.Name})
	}
	if criteria.ID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.ID})
	}
	if criteria.CityID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_CITY_ID: criteria.CityID})
	}
	if criteria.StateID > 0 {
		stmt = stmt.Join(`DAS.CITY C ON C.ID = DAS.STUDIO.CITY_ID`).
			Join(`DAS.STATE S ON S.ID = C.STATE_ID`).Where(squirrel.Eq{`S.ID`: criteria.StateID})
	}
	rows, err := stmt.RunWith(repo.Database).Query()
	studios := make([]reference.Studio, 0)
	if err != nil {
		return studios, err
	}

	for rows.Next() {
		each := reference.Studio{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Address,
			&each.CityID,
			&each.Website,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		studios = append(studios, each)
	}
	rows.Close()
	return studios, err
}

func (repo PostgresStudioRepository) CreateStudio(studio *reference.Studio) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresStudioRepository is not specified")
	}
	stmt := repo.SqlBuilder.Insert("").Into(DAS_STUDIO_TABLE).Columns(
		common.COL_NAME,
		common.COL_ADDRESS,
		common.COL_CITY_ID,
		common.COL_WEBSITE,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated,
	).Values(
		studio.Name,
		studio.Address,
		studio.CityID,
		studio.Website,
		studio.CreateUserID,
		studio.DateTimeCreated,
		studio.UpdateUserID,
		studio.DateTimeUpdated,
	).Suffix(
		"RETURNING ID",
	)

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&studio.ID)
		tx.Commit()
	}
	return err
}

func (repo PostgresStudioRepository) UpdateStudio(studio reference.Studio) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresStudioRepository is not specified")
	}
	stmt := repo.SqlBuilder.Update("").Table(DAS_STUDIO_TABLE)
	if studio.ID > 0 {
		stmt = stmt.Set(common.COL_NAME, studio.Name).
			Set(common.COL_ADDRESS, studio.Address).
			Set(common.COL_CITY_ID, studio.CityID).
			Set(common.COL_WEBSITE, studio.Website).
			Set(common.ColumnUpdateUserID, studio.UpdateUserID).
			Set(common.ColumnDateTimeUpdated, studio.DateTimeUpdated)
		var err error
		if tx, txErr := repo.Database.Begin(); txErr != nil {
			return txErr
		} else {
			_, err = stmt.RunWith(repo.Database).Exec()
			tx.Commit()
		}
		return err
	} else {
		return errors.New("studio is not specified")
	}
}

func (repo PostgresStudioRepository) DeleteStudio(studio reference.Studio) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresStudioRepository is not specified")
	}
	stmt := repo.SqlBuilder.Delete("").From(DAS_STUDIO_TABLE).Where(squirrel.Eq{common.ColumnPrimaryKey: studio.ID})
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		tx.Commit()
	}
	return err
}
