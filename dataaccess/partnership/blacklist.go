package partnership

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/dataaccess/common"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
)

const (
	DAS_PARTNERSHIP_REQUEST_BLACKLIST_TABLE                   = "DAS.PARTNERSHIP_REQUEST_BLACKLIST"
	DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_REPORTER_ID         = "REPORTER_ID"
	DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_BLOCKED_USER_ID     = "BLOCKED_USER_ID"
	DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_BLACKLIST_REASON_ID = "BLACKLIST_REASON_ID"
	DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_DETAIL              = "DETAIL"
	DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_WHITELISTED_IND     = "WHITELISTED_IND"
)

type PostgresPartnershipRequestBlacklistRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresPartnershipRequestBlacklistRepository) SearchPartnershipRequestBlacklist(criteria businesslogic.SearchPartnershipRequestBlacklistCriteria) ([]businesslogic.PartnershipRequestBlacklistEntry, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresPartnershipRequestBlacklistRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
			common.PRIMARY_KEY,
			DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_REPORTER_ID,
			DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_BLOCKED_USER_ID,
			DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_BLACKLIST_REASON_ID,
			DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_DETAIL,
			DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_WHITELISTED_IND,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED)).
		From(DAS_PARTNERSHIP_REQUEST_BLACKLIST_TABLE).
		OrderBy(common.COL_DATETIME_CREATED)
	if criteria.ReporterID > 0 {
		stmt = stmt.Where(squirrel.Eq{DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_REPORTER_ID: criteria.ReporterID})
	}
	if criteria.BlockedUserID > 0 {
		stmt = stmt.Where(squirrel.Eq{DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_BLOCKED_USER_ID: criteria.BlockedUserID})
	}
	if criteria.ReasonID > 0 {
		stmt = stmt.Where(squirrel.Eq{DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_BLACKLIST_REASON_ID: criteria.ReasonID})
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
			&entry.ReporterID,
			&entry.BlockedUserID,
			&entry.BlackListReasonID,
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
	stmt := repo.SqlBuilder.Insert("").Into(DAS_PARTNERSHIP_REQUEST_BLACKLIST_TABLE).Columns(
		DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_REPORTER_ID,
		DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_BLOCKED_USER_ID,
		DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_BLACKLIST_REASON_ID,
		DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_DETAIL,
		DAS_PARTNERSHIP_REQUEST_BLACKLIST_COL_WHITELISTED_IND,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED,
	).Values(
		blacklist.ReporterID,
		blacklist.BlockedUserID,
		blacklist.BlackListReasonID,
		blacklist.Detail,
		blacklist.Whitelisted,
		blacklist.CreateUserID,
		blacklist.DateTimeCreated,
		blacklist.UpdateUserID,
		blacklist.DateTimeUpdated,
	).Suffix(
		fmt.Sprintf("RETURNING %s", common.PRIMARY_KEY),
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
		From(DAS_PARTNERSHIP_REQUEST_BLACKLIST_TABLE).
		Where(squirrel.Eq{common.PRIMARY_KEY: blacklist.ID})
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
	stmt := repo.SqlBuilder.Update("").Table(DAS_PARTNERSHIP_REQUEST_BLACKLIST_TABLE)
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
