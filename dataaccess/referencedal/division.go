package referencedal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/DancesportSoftware/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
)

const (
	DAS_DIVISION_TABLE = "DAS.DIVISION"
)

type PostgresDivisionRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresDivisionRepository) SearchDivision(criteria businesslogic.SearchDivisionCriteria) ([]businesslogic.Division, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s",
			common.ColumnPrimaryKey,
			common.COL_NAME,
			common.ColumnAbbreviation,
			common.COL_DESCRIPTION,
			common.COL_NOTE,
			common.COL_FEDERATION_ID,
			common.ColumnDateTimeCreated,
			common.ColumnDateTimeUpdated)).
		From(DAS_DIVISION_TABLE).
		OrderBy(common.ColumnPrimaryKey)
	if criteria.FederationID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_FEDERATION_ID: criteria.FederationID})
	}
	if criteria.ID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.ID})
	}
	rows, err := stmt.RunWith(repo.Database).Query()
	divisions := make([]businesslogic.Division, 0)
	if err != nil {
		return divisions, err
	}
	for rows.Next() {
		each := businesslogic.Division{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Abbreviation,
			&each.Description,
			&each.Note,
			&each.FederationID,
			&each.DateTimeCreated,
			&each.DateTimeUpdated,
		)
		divisions = append(divisions, each)
	}
	rows.Close()
	return divisions, err
}

func (repo PostgresDivisionRepository) CreateDivision(division *businesslogic.Division) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Insert("").Into(DAS_DIVISION_TABLE).Columns(
		common.COL_NAME,
		common.ColumnAbbreviation,
		common.COL_DESCRIPTION,
		common.COL_NOTE,
		common.COL_FEDERATION_ID,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated,
	).Values(
		division.Name,
		division.Abbreviation,
		division.Description,
		division.Note,
		division.FederationID,
		division.CreateUserID,
		division.DateTimeCreated,
		division.UpdateUserID,
		division.DateTimeUpdated,
	).Suffix(
		"RETURNING ID",
	)

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&division.ID)
		tx.Commit()
	}
	return err
}

func (repo PostgresDivisionRepository) UpdateDivision(division businesslogic.Division) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Update("").Table(DAS_DIVISION_TABLE)
	if division.ID > 0 {
		stmt = stmt.Set(common.COL_NAME, division.Name).
			Set(common.ColumnAbbreviation, division.Abbreviation).
			Set(common.COL_DESCRIPTION, division.Description).
			Set(common.COL_NOTE, division.Note).
			Set(common.COL_FEDERATION_ID, division.FederationID).
			Set(common.ColumnUpdateUserID, division.UpdateUserID).
			Set(common.ColumnDateTimeUpdated, division.DateTimeUpdated)

		var err error
		if tx, txErr := repo.Database.Begin(); txErr != nil {
			return txErr
		} else {
			_, err = stmt.RunWith(repo.Database).Exec()
			tx.Commit()
		}
		return err
	} else {
		return errors.New("division is not specified")
	}
}

func (repo PostgresDivisionRepository) DeleteDivision(division businesslogic.Division) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.
		Delete("").
		From(DAS_DIVISION_TABLE).
		Where(squirrel.Eq{common.ColumnPrimaryKey: division.ID})
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		tx.Commit()
	}
	return err
}
