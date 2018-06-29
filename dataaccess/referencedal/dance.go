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
	"log"
)

const (
	DAS_DANCE_TABLE = "DAS.DANCE"
)

type PostgresDanceRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresDanceRepository) SearchDance(criteria referencebll.SearchDanceCriteria) ([]referencebll.Dance, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresDanceRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s",
			common.PRIMARY_KEY,
			common.COL_NAME,
			common.COL_ABBREVIATION,
			common.COL_DESCRIPTION,
			common.COL_STYLE_ID,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED)).
		From(DAS_DANCE_TABLE).OrderBy(common.PRIMARY_KEY)
	if len(criteria.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_NAME: criteria.Name})
	}
	if criteria.StyleID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_STYLE_ID: criteria.StyleID})
	}
	if criteria.DanceID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.PRIMARY_KEY: criteria.DanceID})
	}
	rows, err := stmt.RunWith(repo.Database).Query()
	dances := make([]referencebll.Dance, 0)
	if err != nil {
		return dances, err
	}

	for rows.Next() {
		each := referencebll.Dance{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Abbreviation,
			&each.Description,
			&each.StyleID,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		dances = append(dances, each)
	}
	rows.Close()
	return dances, err
}

func (repo PostgresDanceRepository) CreateDance(dance *referencebll.Dance) error {
	stmt := repo.SqlBuilder.Insert("").Into(DAS_DANCE_TABLE).Columns(
		common.COL_NAME,
		common.COL_ABBREVIATION,
		common.COL_DESCRIPTION,
		common.COL_STYLE_ID,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED,
	).Values(
		dance.Name,
		dance.Abbreviation,
		dance.Description,
		dance.StyleID,
		dance.CreateUserID,
		dance.DateTimeCreated,
		dance.UpdateUserID,
		dance.DateTimeUpdated,
	).Suffix(
		"RETURNING ID",
	)

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&dance.ID)
		tx.Commit()
	}
	return err
}

func (repo PostgresDanceRepository) UpdateDance(dance referencebll.Dance) error {
	stmt := repo.SqlBuilder.Update("").Table(DAS_DANCE_TABLE)
	if dance.ID > 0 {
		stmt = stmt.Set(common.COL_NAME, dance.Name).
			Set(common.COL_ABBREVIATION, dance.Abbreviation).
			Set(common.COL_DESCRIPTION, dance.Description).
			Set(common.COL_STYLE_ID, dance.StyleID).
			Set(common.COL_UPDATE_USER_ID, dance.UpdateUserID).
			Set(common.COL_DATETIME_UPDATED, dance.DateTimeUpdated)

		var err error
		if tx, txErr := repo.Database.Begin(); txErr != nil {
			return txErr
		} else {
			_, err = stmt.RunWith(repo.Database).Exec()
			tx.Commit()
		}
		return err
	}
	return errors.New("not implemented")
}

func (repo PostgresDanceRepository) DeleteDance(dance referencebll.Dance) error {
	if repo.Database == nil {
		log.Println(common.ErrorMessageEmptyDatabase)
	}
	stmt := repo.SqlBuilder.Delete("").From(DAS_DANCE_TABLE).Where(
		squirrel.Eq{common.PRIMARY_KEY: dance.ID},
	)
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		tx.Commit()
	}
	return err
}
