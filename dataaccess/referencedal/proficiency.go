package referencedal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/dataaccess/common"
	"github.com/ProximaB/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
	"log"
)

const (
	DAS_PROFICIENCY_TABLE = "DAS.PROFICIENCY"
)

type PostgresProficiencyRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresProficiencyRepository) CreateProficiency(proficiency *businesslogic.Proficiency) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Insert("").Into(DAS_PROFICIENCY_TABLE).Columns(
		common.COL_NAME,
		common.COL_DIVISION_ID,
		common.COL_DESCRIPTION,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated,
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

func (repo PostgresProficiencyRepository) UpdateProficiency(proficiency businesslogic.Proficiency) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Update("").Table(DAS_PROFICIENCY_TABLE)
	if proficiency.ID > 0 {
		stmt = stmt.Set(common.COL_NAME, proficiency.Name).
			Set(common.COL_DIVISION_ID, proficiency.DivisionID).
			Set(common.COL_DESCRIPTION, proficiency.Description).
			Set(common.ColumnUpdateUserID, proficiency.UpdateUserID).
			Set(common.ColumnDateTimeUpdated, proficiency.DateTImeUpdated)
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

func (repo PostgresProficiencyRepository) DeleteProficiency(proficiency businesslogic.Proficiency) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.
		Delete("").
		From(DAS_PROFICIENCY_TABLE).
		Where(squirrel.Eq{common.ColumnPrimaryKey: proficiency.ID})
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

func (repo PostgresProficiencyRepository) SearchProficiency(criteria businesslogic.SearchProficiencyCriteria) ([]businesslogic.Proficiency, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s",
		common.ColumnPrimaryKey,
		common.COL_NAME,
		common.COL_DIVISION_ID,
		common.COL_DESCRIPTION,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated)).
		From(DAS_PROFICIENCY_TABLE)

	if criteria.DivisionID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_DIVISION_ID: criteria.DivisionID})
	}
	if criteria.ProficiencyID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.ProficiencyID})
	}
	if len(criteria.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_NAME: criteria.Name})
	}
	rows, err := stmt.RunWith(repo.Database).Query()
	proficiencies := make([]businesslogic.Proficiency, 0)
	if err != nil {
		return proficiencies, err
	}
	for rows.Next() {
		each := businesslogic.Proficiency{}
		scanErr := rows.Scan(
			&each.ID,
			&each.Name,
			&each.DivisionID,
			&each.Description,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTImeUpdated,
		)
		if scanErr != nil {
			log.Printf("[error] scanning proficiency: %v", scanErr)
			return proficiencies, scanErr
		}
		proficiencies = append(proficiencies, each)
	}
	return proficiencies, rows.Close()
}
