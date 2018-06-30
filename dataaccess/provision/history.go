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

package provision

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
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
		return errors.New("data source of PostgresOrganizerProvisionHistoryRepository is not specified")
	}
	stmt := repo.SqlBuilder.Insert("").Into(DAS_ORGANIZER_PROVISION_HISTORY).Columns(
		DAS_ORGANIZER_PROVISION_HISTORY_COL_ORGANIZER_ID,
		DAS_ORGANIZER_PROVISION_HISTORY_COL_AMOUNT,
		DAS_ORGANIZER_PROVISION_HISTORY_COL_NOTE,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED,
	).Values(
		history.OrganizerID,
		history.Amount,
		history.Note,
		history.CreateUserID,
		history.DateTimeCreated,
		history.UpdateUserID,
		history.DateTimeUpdated,
	).Suffix("RETURNING ID")
	clause, args, err := stmt.ToSql()

	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&history.ID)
		err = tx.Commit()
		return err
	}
}

func (repo PostgresOrganizerProvisionHistoryRepository) SearchOrganizerProvisionHistory(criteria businesslogic.SearchOrganizerProvisionHistoryCriteria) ([]businesslogic.OrganizerProvisionHistoryEntry, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresOrganizerProvisionHistoryRepository is not specified")
	}
	clause := repo.SqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s",
		common.PRIMARY_KEY,
		DAS_ORGANIZER_PROVISION_COL_ORGANIZER_ID,
		DAS_ORGANIZER_PROVISION_HISTORY_COL_AMOUNT,
		common.COL_NOTE,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED)).
		From(DAS_ORGANIZER_PROVISION_HISTORY).
		Where(squirrel.Eq{"ORGANIZER_ID": criteria.OrganizerID})

	history := make([]businesslogic.OrganizerProvisionHistoryEntry, 0)
	rows, err := clause.RunWith(repo.Database).Query()

	if err != nil {
		rows.Close()
		return history, err
	}

	for rows.Next() {
		each := businesslogic.OrganizerProvisionHistoryEntry{}
		rows.Scan(
			&each.ID,
			&each.OrganizerID,
			&each.Amount,
			&each.Note,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
	}
	rows.Close()
	return history, err
}

func (repo PostgresOrganizerProvisionHistoryRepository) DeleteOrganizerProvisionHistory(history businesslogic.OrganizerProvisionHistoryEntry) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresOrganizerProvisionHistoryRepository is not specified")
	}
	return errors.New("not implemented")
}

func (repo PostgresOrganizerProvisionHistoryRepository) UpdateOrganizerProvisionHistory(history businesslogic.OrganizerProvisionHistoryEntry) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresOrganizerProvisionHistoryRepository is not specified")
	}
	return errors.New("not implemented")
}
