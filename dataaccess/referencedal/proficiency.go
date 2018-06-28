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
	DAS_PROFICIENCY_TABLE = "DAS.PROFICIENCY"
)

type PostgresProficiencyRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresProficiencyRepository) CreateProficiency(proficiency *referencebll.Proficiency) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresProficiencyRepository is not specified")
	}
	stmt := repo.SqlBuilder.Insert("").Into(DAS_PROFICIENCY_TABLE).Columns(
		common.COL_NAME,
		common.COL_DIVISION_ID,
		common.COL_DESCRIPTION,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED,
	).Values(
		proficiency.Name,
		proficiency.DivisionID,
		proficiency.Description,
		proficiency.CreateUserID,
		proficiency.DateTimeCreated,
		proficiency.UpdateUserID,
		proficiency.DateTImeUpdated,
	).Suffix(
		"RETURNING ID",
	)

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&proficiency.ID)
		tx.Commit()
	}
	return err
}

func (repo PostgresProficiencyRepository) UpdateProficiency(proficiency referencebll.Proficiency) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresProficiencyRepository is not specified")
	}
	stmt := repo.SqlBuilder.Update("").Table(DAS_PROFICIENCY_TABLE)
	if proficiency.ID > 0 {
		stmt = stmt.Set(common.COL_NAME, proficiency.Name).
			Set(common.COL_DIVISION_ID, proficiency.DivisionID).
			Set(common.COL_DESCRIPTION, proficiency.Description).
			Set(common.COL_UPDATE_USER_ID, proficiency.UpdateUserID).
			Set(common.COL_DATETIME_UPDATED, proficiency.DateTImeUpdated)
		var err error
		if tx, txErr := repo.Database.Begin(); txErr != nil {
			return txErr
		} else {
			_, err = stmt.RunWith(repo.Database).Exec()
			err = tx.Commit()
			if err != nil {
				tx.Rollback()
			}
		}
		return err
	} else {
		return errors.New("proficiency is not specified")
	}
}

func (repo PostgresProficiencyRepository) DeleteProficiency(proficiency referencebll.Proficiency) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresProficiencyRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Delete("").
		From(DAS_PROFICIENCY_TABLE).
		Where(squirrel.Eq{common.PRIMARY_KEY: proficiency.ID})
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		if err = tx.Commit(); err != nil {
			tx.Rollback()
		}
		return err
	}
}

func (repo PostgresProficiencyRepository) SearchProficiency(criteria referencebll.SearchProficiencyCriteria) ([]referencebll.Proficiency, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresProficiencyRepository is not specified")
	}
	stmt := repo.SqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s",
		common.PRIMARY_KEY,
		common.COL_NAME,
		common.COL_DIVISION_ID,
		common.COL_DESCRIPTION,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED)).
		From(DAS_PROFICIENCY_TABLE)

	if criteria.DivisionID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_DIVISION_ID: criteria.DivisionID})
	}
	if criteria.ProficiencyID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.PRIMARY_KEY: criteria.ProficiencyID})
	}
	rows, err := stmt.RunWith(repo.Database).Query()
	proficiencies := make([]referencebll.Proficiency, 0)
	if err != nil {
		return proficiencies, err
	}
	for rows.Next() {
		each := referencebll.Proficiency{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.DivisionID,
			&each.Description,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTImeUpdated,
		)
		proficiencies = append(proficiencies, each)
	}
	rows.Close()
	return proficiencies, err
}
