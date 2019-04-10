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
	DAS_STYLE_TABLE = "DAS.STYLE"
)

// PostgresStyleRepository implements IStyleRepository and feeds data to
// businesslogic from a PostgreSQL database
type PostgresStyleRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresStyleRepository) CreateStyle(style *businesslogic.Style) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Insert("").Into(DAS_STYLE_TABLE).Columns(
		common.COL_NAME,
		common.COL_DESCRIPTION,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated,
	).Values(
		style.Name,
		style.Description,
		style.CreateUserID,
		style.DateTimeCreated,
		style.UpdateUserID,
		style.DateTimeUpdated,
	).Suffix(
		"RETURNING ID",
	)

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&style.ID)
		tx.Commit()
	}
	return err
}

func (repo PostgresStyleRepository) DeleteStyle(style businesslogic.Style) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.
		Delete("").
		From(DAS_STYLE_TABLE).
		Where(squirrel.Eq{common.ColumnPrimaryKey: style.ID})
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		tx.Commit()
	}
	return err
}

func (repo PostgresStyleRepository) SearchStyle(criteria businesslogic.SearchStyleCriteria) ([]businesslogic.Style, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Select(
		fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s",
			common.ColumnPrimaryKey,
			common.COL_NAME,
			common.COL_DESCRIPTION,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated)).
		From(DAS_STYLE_TABLE).
		OrderBy(common.ColumnPrimaryKey)
	if criteria.StyleID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.StyleID})
	}
	if len(criteria.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_NAME: criteria.Name})
	}
	rows, err := stmt.RunWith(repo.Database).Query()
	styles := make([]businesslogic.Style, 0)
	if err != nil {
		return styles, err
	}
	for rows.Next() {
		each := businesslogic.Style{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Description,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		styles = append(styles, each)
	}
	rows.Close()
	return styles, err
}

func (repo PostgresStyleRepository) UpdateStyle(style businesslogic.Style) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Update("").Table(DAS_STYLE_TABLE)
	if style.ID > 0 {
		stmt = stmt.Set(common.COL_NAME, style.Name).
			Set(common.COL_DESCRIPTION, style.Description).
			Set(common.ColumnUpdateUserID, style.UpdateUserID).
			Set(common.ColumnDateTimeUpdated, style.DateTimeUpdated)
		var err error
		if tx, txErr := repo.Database.Begin(); txErr != nil {
			return txErr
		} else {
			_, err = stmt.RunWith(repo.Database).Exec()
			if err != nil {
				tx.Rollback()
			} else {
				tx.Commit()
			}
		}
		return err
	} else {
		return errors.New("style is not specified")
	}
}
