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

package partnershipdal

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
	DasPartnershipRequestBlacklistTable                   = "DAS.PARTNERSHIP_REQUEST_BLACKLIST"
	DasPartnershipRequestBlacklistColumnReporterID        = "REPORTER_ID"
	DasPartnershipRequestBlacklistColumnBlockedUserID     = "BLOCKED_USER_ID"
	DasPartnershipRequestBlacklistColumnBlacklistReasonID = "BLACKLIST_REASON_ID"
	DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_DETAIL          = "DETAIL"
	DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_WHITELISTED_IND = "WHITELISTED_IND"
)

type PostgresPartnershipRequestBlacklistRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresPartnershipRequestBlacklistRepository) SearchPartnershipRequestBlacklist(criteria businesslogic.SearchPartnershipRequestBlacklistCriteria) ([]businesslogic.PartnershipRequestBlacklistEntry, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
			common.ColumnPrimaryKey,
			DasPartnershipRequestBlacklistColumnReporterID,
			DasPartnershipRequestBlacklistColumnBlockedUserID,
			DasPartnershipRequestBlacklistColumnBlacklistReasonID,
			DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_DETAIL,
			DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_WHITELISTED_IND,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated)).
		From(DasPartnershipRequestBlacklistTable).
		OrderBy(common.ColumnDateTimeCreated)
	if criteria.ReporterID > 0 {
		stmt = stmt.Where(squirrel.Eq{DasPartnershipRequestBlacklistColumnReporterID: criteria.ReporterID})
	}
	if criteria.BlockedUserID > 0 {
		stmt = stmt.Where(squirrel.Eq{DasPartnershipRequestBlacklistColumnBlockedUserID: criteria.BlockedUserID})
	}
	if criteria.ReasonID > 0 {
		stmt = stmt.Where(squirrel.Eq{DasPartnershipRequestBlacklistColumnBlacklistReasonID: criteria.ReasonID})
	}
	stmt = stmt.Where(squirrel.Eq{DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_WHITELISTED_IND: criteria.Whitelisted})

	rows, err := stmt.RunWith(repo.Database).Query()
	blacklist := make([]businesslogic.PartnershipRequestBlacklistEntry, 0)
	if err != nil {
		return blacklist, err
	}
	for rows.Next() {
		entry := businesslogic.PartnershipRequestBlacklistEntry{}
		rows.Scan(
			&entry.ID,
			&entry.Reporter.ID,
			&entry.BlockedUser.ID,
			&entry.BlockedReason.ID,
			&entry.Detail,
			&entry.Whitelisted,
			&entry.CreateUserID,
			&entry.DateTimeCreated,
			&entry.UpdateUserID,
			&entry.DateTimeUpdated,
		)
		blacklist = append(blacklist, entry)
	}
	rows.Close()
	return blacklist, err
}
func (repo PostgresPartnershipRequestBlacklistRepository) CreatePartnershipRequestBlacklist(blacklist *businesslogic.PartnershipRequestBlacklistEntry) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresPartnershipRequestBlacklistRepository is not specified")
	}
	stmt := repo.SqlBuilder.Insert("").Into(DasPartnershipRequestBlacklistTable).Columns(
		DasPartnershipRequestBlacklistColumnReporterID,
		DasPartnershipRequestBlacklistColumnBlockedUserID,
		DasPartnershipRequestBlacklistColumnBlacklistReasonID,
		DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_DETAIL,
		DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_WHITELISTED_IND,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated,
	).Values(
		blacklist.Reporter.ID,
		blacklist.BlockedUser.ID,
		blacklist.BlockedReason.ID,
		blacklist.Detail,
		blacklist.Whitelisted,
		blacklist.CreateUserID,
		blacklist.DateTimeCreated,
		blacklist.UpdateUserID,
		blacklist.DateTimeUpdated,
	).Suffix(
		"RETURNING ID",
	)

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&blacklist.ID)
		err = tx.Commit()
	}
	return err
}
func (repo PostgresPartnershipRequestBlacklistRepository) DeletePartnershipRequestBlacklist(blacklist businesslogic.PartnershipRequestBlacklistEntry) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresPartnershipRequestBlacklistRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Delete("").
		From(DasPartnershipRequestBlacklistTable).
		Where(squirrel.Eq{common.ColumnPrimaryKey: blacklist.ID})
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		tx.Commit()
	}
	return err
}

// UpdatePartnershipRequestBlacklist will only update the whitelist status. Updating the detail or report reason is
// not allowed.
func (repo PostgresPartnershipRequestBlacklistRepository) UpdatePartnershipRequestBlacklist(blacklist businesslogic.PartnershipRequestBlacklistEntry) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresPartnershipRequestBlacklistRepository is not specified")
	}
	stmt := repo.SqlBuilder.Update("").Table(DasPartnershipRequestBlacklistTable)
	if blacklist.ID > 0 {
		stmt = stmt.Set(DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_WHITELISTED_IND, blacklist.Whitelisted)

		var err error
		if tx, txErr := repo.Database.Begin(); txErr != nil {
			return txErr
		} else {
			_, err = stmt.RunWith(repo.Database).Exec()
			tx.Commit()
		}
		return err
	} else {
		return errors.New("blacklist is not specified")
	}
}
