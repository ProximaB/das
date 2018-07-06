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
	"log"
)

const (
	DAS_FEDERATION_TABLE            = "DAS.FEDERATION"
	DAS_FEDERATION_COL_YEAR_FOUNDED = "YEAR_FOUNDED"
)

type PostgresFederationRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresFederationRepository) CreateFederation(federation *referencebll.Federation) error {
	if repo.Database == nil {
		log.Println(common.ErrorMessageEmptyDatabase)
	}
	stmt := repo.SqlBuilder.Insert("").
		Into(DAS_FEDERATION_TABLE).
		Columns(
			common.COL_NAME,
			common.COL_ABBREVIATION,
			common.COL_DESCRIPTION,
			DAS_FEDERATION_COL_YEAR_FOUNDED,
			common.COL_COUNTRY_ID,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED,
		).Values(
		federation.Name,
		federation.Abbreviation,
		federation.Description,
		federation.YearFounded,
		federation.YearFounded,
		federation.CreateUserID,
		federation.DateTimeCreated,
		federation.UpdateUserID,
		federation.DateTimeUpdated,
	).Suffix("RETURNING ID")

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&federation.ID)
		tx.Commit()
	}

	return err
}

func (repo PostgresFederationRepository) SearchFederation(criteria referencebll.SearchFederationCriteria) ([]referencebll.Federation, error) {
	if repo.Database == nil {
		log.Println(common.ErrorMessageEmptyDatabase)
	}
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s",
			common.PRIMARY_KEY,
			common.COL_NAME,
			common.COL_ABBREVIATION,
			DAS_FEDERATION_COL_YEAR_FOUNDED,
			common.COL_COUNTRY_ID,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED)).
		From(DAS_FEDERATION_TABLE).OrderBy(common.PRIMARY_KEY)
	if criteria.CountryID > 0 {
		stmt = stmt.Where(squirrel.Eq{
			common.COL_COUNTRY_ID: criteria.CountryID})
	}
	if len(criteria.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_NAME: criteria.Name})
	}
	if criteria.ID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.PRIMARY_KEY: criteria.ID})
	}

	federations := make([]referencebll.Federation, 0)
	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return federations, err
	}
	for rows.Next() {
		each := referencebll.Federation{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Abbreviation,
			&each.YearFounded,
			&each.CountryID,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		federations = append(federations, each)
	}
	rows.Close()
	return federations, err
}

func (repo PostgresFederationRepository) DeleteFederation(federation referencebll.Federation) error {
	if repo.Database == nil {
		log.Println(common.ErrorMessageEmptyDatabase)
	}
	stmt := repo.SqlBuilder.Delete("").From(DAS_FEDERATION_TABLE).Where(squirrel.Eq{common.PRIMARY_KEY: federation.ID})

	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		tx.Commit()
	}
	return err
}

func (repo PostgresFederationRepository) UpdateFederation(federation referencebll.Federation) error {
	if repo.Database == nil {
		log.Println(common.ErrorMessageEmptyDatabase)
	}
	stmt := repo.SqlBuilder.Update("").Table(DAS_FEDERATION_TABLE)
	if federation.ID > 0 {
		stmt = stmt.Set(common.COL_NAME, federation.Name).
			Set(common.COL_ABBREVIATION, federation.Abbreviation).
			Set(common.COL_DESCRIPTION, federation.Description).
			Set(DAS_FEDERATION_COL_YEAR_FOUNDED, federation.YearFounded).
			Set(common.COL_COUNTRY_ID, federation.CountryID).
			Set(common.COL_UPDATE_USER_ID, federation.UpdateUserID).
			Set(common.COL_DATETIME_UPDATED, federation.DateTimeUpdated)
		var err error
		if tx, txErr := repo.Database.Begin(); txErr != nil {
			return txErr
		} else {
			_, err = stmt.RunWith(repo.Database).Exec()
			tx.Commit()
		}
		return err
	} else {
		return errors.New("federation is not specified")
	}
}
