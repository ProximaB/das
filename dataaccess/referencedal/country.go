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
	"time"
)

const (
	DAS_COUNTRY_TABLE = "DAS.COUNTRY"
)

type PostgresCountryRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresCountryRepository) CreateCountry(country *referencebll.Country) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresCountryRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Insert("").
		Into(DAS_COUNTRY_TABLE).
		Columns(common.COL_NAME,
			common.COL_ABBREVIATION,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED).
		Values(
			country.Name,
			country.Abbreviation,
			country.CreateUserID,
			country.DateTimeCreated,
			country.CreateUserID,
			country.DateTimeUpdated,
		).Suffix(
		"RETURNING ID",
	)
	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&country.ID)
		tx.Commit()
	}
	return err
}

func (repo PostgresCountryRepository) DeleteCountry(country referencebll.Country) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresCountryRepository is not specified")
	}
	stmt := repo.SqlBuilder.Delete("").From(DAS_COUNTRY_TABLE).
		Where(squirrel.Eq{common.PRIMARY_KEY: country.ID})
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		tx.Commit()
	}
	return err
}

func (repo PostgresCountryRepository) UpdateCountry(country referencebll.Country) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresCountryRepository is not specified")
	}
	stmt := repo.SqlBuilder.Update("").Table(DAS_COUNTRY_TABLE)
	if country.ID > 0 {
		stmt = stmt.Set(common.COL_NAME, country.Name).
			Set(common.COL_ABBREVIATION, country.Abbreviation).
			Set(common.COL_DATETIME_UPDATED, time.Now()).
			Where(squirrel.Eq{common.PRIMARY_KEY: country.ID})
		if country.UpdateUserID != nil {
			stmt = stmt.Set(common.COL_UPDATE_USER_ID, country.UpdateUserID)
		}

		var err error
		if tx, txErr := repo.Database.Begin(); txErr != nil {
			return txErr
		} else {
			_, err = stmt.RunWith(repo.Database).Exec()
			tx.Commit()
		}

		return err
	} else {
		return errors.New("country is not specified")
	}
}

func (repo PostgresCountryRepository) SearchCountry(criteria referencebll.SearchCountryCriteria) ([]referencebll.Country, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresCountryRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s",
			common.PRIMARY_KEY,
			common.COL_NAME,
			common.COL_ABBREVIATION,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED)).
		From(DAS_COUNTRY_TABLE).
		OrderBy(common.PRIMARY_KEY)
	if criteria.CountryID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.PRIMARY_KEY: criteria.CountryID})
	}
	if len(criteria.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_NAME: criteria.Name})
	}
	if len(criteria.Abbreviation) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_ABBREVIATION: criteria.Abbreviation})
	}

	rows, err := stmt.RunWith(repo.Database).Query()
	countries := make([]referencebll.Country, 0)
	if err != nil {
		return countries, err
	}
	for rows.Next() {
		each := referencebll.Country{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Abbreviation,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		countries = append(countries, each)
	}
	rows.Close()
	return countries, err
}