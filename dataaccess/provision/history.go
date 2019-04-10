package provision

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
	DAS_ORGANIZER_PROVISION_HISTORY                  = "DAS.ORGANIZER_PROVISION_HISTORY"
	DAS_ORGANIZER_PROVISION_HISTORY_COL_ORGANIZER_ID = "ORGANIZER_ID"
	DAS_ORGANIZER_PROVISION_HISTORY_COL_AMOUNT       = "AMOUNT"
	DAS_ORGANIZER_PROVISION_HISTORY_COL_NOTE         = "NOTE"
)

type PostgresOrganizerProvisionHistoryRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresOrganizerProvisionHistoryRepository) CreateOrganizerProvisionHistory(history *businesslogic.OrganizerProvisionHistoryEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Insert("").Into(DAS_ORGANIZER_PROVISION_HISTORY).Columns(
		DAS_ORGANIZER_PROVISION_HISTORY_COL_ORGANIZER_ID,
		DAS_ORGANIZER_PROVISION_HISTORY_COL_AMOUNT,
		DAS_ORGANIZER_PROVISION_HISTORY_COL_NOTE,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated,
	).Values(
		history.OrganizerRoleID,
		history.Amount,
		history.Note,
		history.CreateUserID,
		history.DateTimeCreated,
		history.UpdateUserID,
		history.DateTimeUpdated,
	).Suffix("RETURNING ID")
	clause, args, err := stmt.ToSql()

	tx, txErr := repo.Database.Begin()
	if txErr != nil {
		return txErr
	}
	row := repo.Database.QueryRow(clause, args...)
	err = row.Scan(&history.ID)
	if err != nil {
		return err
	}
	err = tx.Commit()
	return err

}

func (repo PostgresOrganizerProvisionHistoryRepository) SearchOrganizerProvisionHistory(criteria businesslogic.SearchOrganizerProvisionHistoryCriteria) ([]businesslogic.OrganizerProvisionHistoryEntry, error) {
	history := make([]businesslogic.OrganizerProvisionHistoryEntry, 0)
	if repo.Database == nil {
		return history, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	clause := repo.SqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s",
		common.ColumnPrimaryKey,
		DAS_ORGANIZER_PROVISION_COL_ORGANIZER_ID,
		DAS_ORGANIZER_PROVISION_HISTORY_COL_AMOUNT,
		common.COL_NOTE,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated)).
		From(DAS_ORGANIZER_PROVISION_HISTORY).
		Where(squirrel.Eq{"ORGANIZER_ID": criteria.OrganizerID})

	rows, err := clause.RunWith(repo.Database).Query()
	if err != nil {
		return history, err
	}

	for rows.Next() {
		each := businesslogic.OrganizerProvisionHistoryEntry{}
		err = rows.Scan(
			&each.ID,
			&each.OrganizerRoleID,
			&each.Amount,
			&each.Note,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		history = append(history, each)
		if err != nil {
			return history, err
		}
	}
	err = rows.Close()
	return history, err
}
