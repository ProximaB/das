package referencedal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/dataaccess/common"
	"github.com/ProximaB/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
)

const (
	DAS_DANCE_TABLE = "DAS.DANCE"
)

type PostgresDanceRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresDanceRepository) SearchDance(criteria businesslogic.SearchDanceCriteria) ([]businesslogic.Dance, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s",
			common.ColumnPrimaryKey,
			common.COL_NAME,
			common.ColumnAbbreviation,
			common.COL_DESCRIPTION,
			common.COL_STYLE_ID,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated)).
		From(DAS_DANCE_TABLE).OrderBy(common.ColumnPrimaryKey)
	if len(criteria.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_NAME: criteria.Name})
	}
	if criteria.StyleID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_STYLE_ID: criteria.StyleID})
	}
	if criteria.DanceID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.DanceID})
	}
	rows, err := stmt.RunWith(repo.Database).Query()
	dances := make([]businesslogic.Dance, 0)
	if err != nil {
		return dances, err
	}

	for rows.Next() {
		each := businesslogic.Dance{}
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

func (repo PostgresDanceRepository) CreateDance(dance *businesslogic.Dance) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Insert("").Into(DAS_DANCE_TABLE).Columns(
		common.COL_NAME,
		common.ColumnAbbreviation,
		common.COL_DESCRIPTION,
		common.COL_STYLE_ID,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated,
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

func (repo PostgresDanceRepository) UpdateDance(dance businesslogic.Dance) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Update("").Table(DAS_DANCE_TABLE)
	if dance.ID > 0 {
		stmt = stmt.Set(common.COL_NAME, dance.Name).
			Set(common.ColumnAbbreviation, dance.Abbreviation).
			Set(common.COL_DESCRIPTION, dance.Description).
			Set(common.COL_STYLE_ID, dance.StyleID).
			Set(common.ColumnUpdateUserID, dance.UpdateUserID).
			Set(common.ColumnDateTimeUpdated, dance.DateTimeUpdated)

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

func (repo PostgresDanceRepository) DeleteDance(dance businesslogic.Dance) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Delete("").From(DAS_DANCE_TABLE).Where(
		squirrel.Eq{common.ColumnPrimaryKey: dance.ID},
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
