// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

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
	dasCityTable = "DAS.CITY"
)

// PostgresCityRepository implements ICityRepository and provides CRUD operations
// in PostgreSQL database
type PostgresCityRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

// CreateCity inserts a new City record in the database and updates the ID key of city
func (repo PostgresCityRepository) CreateCity(city *referencebll.City) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresCityRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Insert("").
		Into(dasCityTable).
		Columns(common.COL_NAME,
			common.COL_STATE_ID,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED).
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
func (repo PostgresCityRepository) DeleteCity(city referencebll.City) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresCityRepository is not specified")
	}
	stmt := repo.SqlBuilder.Delete("").From(dasCityTable)
	if city.ID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.PRIMARY_KEY: city.ID})
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
func (repo PostgresCityRepository) UpdateCity(city referencebll.City) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresCityRepository is not specified")
	}
	stmt := repo.SqlBuilder.Update("").Table(dasCityTable).
		SetMap(squirrel.Eq{common.COL_NAME: city.Name, common.COL_STATE_ID: city.StateID}).
		SetMap(squirrel.Eq{common.COL_DATETIME_UPDATED: city.DateTimeUpdated}).Where(squirrel.Eq{common.PRIMARY_KEY: city.ID})

	if city.UpdateUserID != nil {
		stmt = stmt.SetMap(squirrel.Eq{common.COL_UPDATE_USER_ID: city.UpdateUserID})
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
func (repo PostgresCityRepository) SearchCity(criteria referencebll.SearchCityCriteria) ([]referencebll.City, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresCityRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s",
			common.PRIMARY_KEY,
			common.COL_NAME,
			common.COL_STATE_ID,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED)).
		From(dasCityTable).OrderBy(common.PRIMARY_KEY)
	if len(criteria.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_NAME: criteria.Name})
	}
	if criteria.StateID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_STATE_ID: criteria.StateID})
	}
	if criteria.CityID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.PRIMARY_KEY: criteria.CityID})
	}

	rows, err := stmt.RunWith(repo.Database).Query()
	cities := make([]referencebll.City, 0)
	if err != nil {
		return cities, err
	}
	for rows.Next() {
		each := referencebll.City{}
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
