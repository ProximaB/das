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
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/Masterminds/squirrel"
)

const (
	dasCityTable = "DAS.CITY"
)

// PostgresCityRepository implements ICityRepository and provides CRUD operations
// in PostgreSQL database
type PostgresCityRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

// CreateCity inserts a new City record in the database and updates the ID key of city
func (repo PostgresCityRepository) CreateCity(city *businesslogic.City) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresCityRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Insert("").
		Into(dasCityTable).
		Columns(common.COL_NAME,
			common.COL_STATE_ID,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated).
		Values(
			city.Name,
			city.StateID,
			city.CreateUserID,
			city.DateTimeCreated,
			city.UpdateUserID,
			city.DateTimeUpdated).Suffix("RETURNING ID")

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&city.ID)
		tx.Commit()
	}
	return err
}

// DeleteCity removes the City record from the database
func (repo PostgresCityRepository) DeleteCity(city businesslogic.City) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresCityRepository is not specified")
	}
	stmt := repo.SqlBuilder.Delete("").From(dasCityTable)
	if city.ID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: city.ID})
	}
	if len(city.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_NAME: city.Name})
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

// UpdateCity updates the value in a City record
func (repo PostgresCityRepository) UpdateCity(city businesslogic.City) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresCityRepository is not specified")
	}
	stmt := repo.SqlBuilder.Update("").Table(dasCityTable).
		SetMap(squirrel.Eq{common.COL_NAME: city.Name, common.COL_STATE_ID: city.StateID}).
		SetMap(squirrel.Eq{common.ColumnDateTimeUpdated: city.DateTimeUpdated}).Where(squirrel.Eq{common.ColumnPrimaryKey: city.ID})

	if city.UpdateUserID != nil {
		stmt = stmt.SetMap(squirrel.Eq{common.ColumnUpdateUserID: city.UpdateUserID})
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

// SearchCity selects cityes
func (repo PostgresCityRepository) SearchCity(criteria businesslogic.SearchCityCriteria) ([]businesslogic.City, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresCityRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s",
			common.ColumnPrimaryKey,
			common.COL_NAME,
			common.COL_STATE_ID,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated)).
		From(dasCityTable).OrderBy(common.ColumnPrimaryKey)
	if len(criteria.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_NAME: criteria.Name})
	}
	if criteria.StateID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_STATE_ID: criteria.StateID})
	}
	if criteria.CityID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.CityID})
	}

	rows, err := stmt.RunWith(repo.Database).Query()
	cities := make([]businesslogic.City, 0)
	if err != nil {
		return cities, err
	}
	for rows.Next() {
		each := businesslogic.City{}
		scanErr := rows.Scan(
			&each.ID,
			&each.Name,
			&each.StateID,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		if scanErr != nil {
			rows.Close()
			return cities, scanErr
		}
		cities = append(cities, each)
	}
	rows.Close()
	return cities, nil
}
