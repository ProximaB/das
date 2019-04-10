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
	DAS_STATE_TABLE = "DAS.STATE"
)

type PostgresStateRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresStateRepository) SearchState(criteria businesslogic.SearchStateCriteria) ([]businesslogic.State, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s",
			common.ColumnPrimaryKey,
			common.COL_NAME,
			common.ColumnAbbreviation,
			common.COL_COUNTRY_ID,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated)).
		From(DAS_STATE_TABLE)
	if criteria.CountryID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_COUNTRY_ID: criteria.CountryID})
	}
	if len(criteria.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_NAME: criteria.Name})
	}
	if criteria.StateID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.StateID})
	}
	stmt = stmt.OrderBy(common.ColumnPrimaryKey,
		common.COL_NAME)

	states := make([]businesslogic.State, 0)
	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return states, err
	}

	for rows.Next() {
		each := businesslogic.State{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Abbreviation,
			&each.CountryID,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		states = append(states, each)
	}
	if err != nil {
		return nil, err
	}
	return states, nil
}

func (repo PostgresStateRepository) CreateState(state *businesslogic.State) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Insert("").Into(DAS_STATE_TABLE).Columns(
		common.COL_NAME,
		common.ColumnAbbreviation,
		common.COL_COUNTRY_ID,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated,
	).Values(
		state.Name,
		state.Abbreviation,
		state.CountryID,
		state.CreateUserID,
		state.DateTimeCreated,
		state.UpdateUserID,
		state.DateTimeUpdated,
	).Suffix(
		"RETURNING ID",
	)

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&state.ID)
		tx.Commit()
	}
	return err
}

func (repo PostgresStateRepository) UpdateState(state businesslogic.State) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Update("").Table(DAS_STATE_TABLE)
	if state.ID > 0 {
		stmt = stmt.Set(common.COL_NAME, state.Name).
			Set(common.ColumnAbbreviation, state.Abbreviation).
			Set(common.COL_COUNTRY_ID, state.CountryID).
			Set(common.ColumnUpdateUserID, state.UpdateUserID).
			Set(common.ColumnDateTimeUpdated, state.DateTimeUpdated)

		var err error
		if tx, txErr := repo.Database.Begin(); txErr != nil {
			return txErr
		} else {
			_, err = stmt.RunWith(repo.Database).Exec()
			if err = tx.Commit(); err != nil {
				tx.Rollback()
			}
		}
		return err
	} else {
		return errors.New("state is not specified")
	}
}

func (repo PostgresStateRepository) DeleteState(state businesslogic.State) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Delete("").From(DAS_STATE_TABLE).Where(squirrel.Eq{common.ColumnPrimaryKey: state.ID})
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		if err = tx.Commit(); err != nil {
			tx.Rollback()
		}
	}
	return err
}
