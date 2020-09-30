package accountdal

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
	dasUserPreferenceTable             = "DAS.USER_PREFERENCE"
	dasUserPreferenceColumnDefaultRole = "DEFAULT_ROLE"
)

// PostgresUserPreferenceRepository implements the IUserPreferenceRepository with a Postgres database
type PostgresUserPreferenceRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

// CreatePreference inserts a User Preference object into a Postgres database
func (repo PostgresUserPreferenceRepository) CreatePreference(preference *businesslogic.UserPreference) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.
		Insert("").
		Into(dasUserPreferenceTable).
		Columns(
			common.ColumnAccountID,
			dasUserPreferenceColumnDefaultRole,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated).
		Values(
			preference.AccountID,
			preference.DefaultRole,
			preference.CreateUserID,
			preference.DateTimeCreated,
			preference.UpdateUserID,
			preference.DateTimeUpdated).
		Suffix(dalutil.SQLSuffixReturningID)
	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&preference.ID)
		err = tx.Commit()
	}
	return err
}

// SearchPreference searches User Preference object from a Postgres database using the provided criteria
func (repo PostgresUserPreferenceRepository) SearchPreference(criteria businesslogic.SearchUserPreferenceCriteria) ([]businesslogic.UserPreference, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.
		Select(
			fmt.Sprintf(
				"%s, %s, %s, %s, %s, %s, %s",
				common.ColumnPrimaryKey,
				common.ColumnAccountID,
				dasUserPreferenceColumnDefaultRole,
				common.ColumnCreateUserID,
				common.ColumnDateTimeCreated,
				common.ColumnUpdateUserID,
				common.ColumnDateTimeUpdated,
			)).From(dasUserPreferenceTable)
	if criteria.AccountID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnAccountID: criteria.AccountID})
	} else {
		return nil, errors.New("account must be specified")
	}

	preferences := make([]businesslogic.UserPreference, 0)
	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		each := businesslogic.UserPreference{}
		rows.Scan(
			&each.ID,
			&each.AccountID,
			&each.DefaultRole,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		preferences = append(preferences, each)
	}
	return preferences, err
}

// UpdatePreference updates a User Preference object in a Postgres database to the provided preference object
func (repo PostgresUserPreferenceRepository) UpdatePreference(preference businesslogic.UserPreference) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	if preference.ID == 0 {
		return errors.New("object UserPreference's ID must be specified")
	}
	if preference.UpdateUserID == 0 {
		return errors.New("update user must be specified")
	}
	stmt := repo.SQLBuilder.Update("").Table(dasUserPreferenceTable)
	if preference.DefaultRole > 0 {
		stmt = stmt.
			Set(dasUserPreferenceColumnDefaultRole, preference.DefaultRole).
			Set(common.ColumnUpdateUserID, preference.UpdateUserID).
			Set(common.ColumnDateTimeUpdated, preference.DateTimeUpdated)

		var err error
		if tx, txErr := repo.Database.Begin(); txErr != nil {
			return txErr
		} else {
			_, err = stmt.RunWith(repo.Database).Exec()
			tx.Commit()
		}
		return err
	}
	return nil
}
