package dataaccess

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/dataaccess/common"
	"database/sql"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"log"
)

const (
	DAS_ORGANIZER_PROVISION                  = "DAS.ORGANIZER_PROVISION"
	DAS_ORGANIZER_PROVISION_COL_ORGANIZER_ID = "ORGANIZER_ID"
	DAS_ORGANIZER_PROVISION_COL_HOSTED       = "HOSTED"
	DAS_ORGANIZER_PROVISION_COL_AVAILABLE    = "AVAILABLE"
)

const (
	DAS_ORGANIZER_PROVISION_HISTORY                  = "DAS.ORGANIZER_PROVISION_HISTORY"
	DAS_ORGANIZER_PROVISION_HISTORY_COL_ORGANIZER_ID = "ORGANIZER_ID"
	DAS_ORGANIZER_PROVISION_HISTORY_COL_AMOUNT       = "AMOUNT"
	DAS_ORGANIZER_PROVISION_HISTORY_COL_NOTE         = "NOTE"
)

type PostgresOrganizerProvisionRepository struct {
	Database   *sql.DB
	SqlBuilder sq.StatementBuilderType
}

type PostgresOrganizerProvisionHistoryRepository struct {
	Database   *sql.DB
	SqlBuilder sq.StatementBuilderType
}

func (repo PostgresOrganizerProvisionRepository) CreateOrganizerProvision(provision businesslogic.OrganizerProvision) error {

	stmt := repo.SqlBuilder.Insert("").
		Into(DAS_ORGANIZER_PROVISION).
		Columns(
			DAS_ORGANIZER_PROVISION_COL_ORGANIZER_ID,
			DAS_ORGANIZER_PROVISION_COL_HOSTED,
			DAS_ORGANIZER_PROVISION_COL_AVAILABLE,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED,
		).Values(provision.OrganizerID, provision.Hosted, provision.Available, provision.CreateUserID, provision.DateTimeCreated, provision.UpdateUserID, provision.DateTimeUpdated)
	_, err := stmt.RunWith(repo.Database).Exec()
	if err != nil {
		log.Printf("[error] initializing organizer organizer: %s\n", err.Error())
		return err
	}

	//CreateOrganizerProvisionHistoryEntry(accountID, 0, "initial organizer", accountID)
	if err != nil {
		log.Printf("[error] initializing organizer organizer history: %s\n", err.Error())
		return err
	}
	return err
}

func (repo PostgresOrganizerProvisionRepository) UpdateOrganizerProvision(provision businesslogic.OrganizerProvision) error {
	stmt := repo.SqlBuilder.Update("").
		Table(DAS_ORGANIZER_PROVISION).
		Set(DAS_ORGANIZER_PROVISION_COL_AVAILABLE, provision.Available).
		Set(DAS_ORGANIZER_PROVISION_COL_HOSTED, provision.Hosted).
		Set(common.COL_DATETIME_UPDATED, provision.DateTimeUpdated).
		Where(sq.Eq{DAS_ORGANIZER_PROVISION_COL_ORGANIZER_ID: provision.OrganizerID})
	_, err := stmt.RunWith(repo.Database).Exec()
	return err
}

func (repo PostgresOrganizerProvisionRepository) SearchOrganizerProvision(
	criteria *businesslogic.SearchOrganizerProvisionCriteria) ([]businesslogic.OrganizerProvision, error) {
	stmt := repo.SqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s",
		common.PRIMARY_KEY,
		DAS_ORGANIZER_PROVISION_COL_ORGANIZER_ID,
		DAS_ORGANIZER_PROVISION_COL_HOSTED,
		DAS_ORGANIZER_PROVISION_COL_AVAILABLE,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED)).
		From(DAS_ORGANIZER_PROVISION).Where(sq.Eq{DAS_ORGANIZER_PROVISION_COL_ORGANIZER_ID: criteria.OrganizerID})

	rows, err := stmt.RunWith(repo.Database).Query()

	provisions := make([]businesslogic.OrganizerProvision, 0)
	for rows.Next() {
		each := businesslogic.OrganizerProvision{}
		rows.Scan(
			&each.ID,
			&each.OrganizerID,
			&each.Hosted,
			&each.Available,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		provisions = append(provisions, each)
	}

	return provisions, err
}

func (repo PostgresOrganizerProvisionRepository) DeleteOrganizerProvision(provision businesslogic.OrganizerProvision) error {
	return errors.New("not implemented")
}

func (repo PostgresOrganizerProvisionHistoryRepository) CreateOrganizerProvisionHistory(history businesslogic.OrganizerProvisionHistoryEntry) error {
	stmt := repo.SqlBuilder.Insert("").
		Into(DAS_ORGANIZER_PROVISION_HISTORY).
		Columns(
			DAS_ORGANIZER_PROVISION_HISTORY_COL_ORGANIZER_ID,
			DAS_ORGANIZER_PROVISION_HISTORY_COL_AMOUNT,
			DAS_ORGANIZER_PROVISION_HISTORY_COL_NOTE,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED,
		).Values(history.OrganizerID, history.Amount, history.Note, history.CreateUserID, history.DateTimeCreated, history.UpdateUserID, history.DateTimeUpdated)
	_, err := stmt.RunWith(repo.Database).Exec()
	return err
}

func (repo PostgresOrganizerProvisionHistoryRepository) SearchOrganizerProvisionHistory(criteria *businesslogic.SearchOrganizerProvisionHistoryCriteria) ([]businesslogic.OrganizerProvisionHistoryEntry, error) {
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
		Where(sq.Eq{"ORGANIZER_ID": criteria.OrganizerID})

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
	return errors.New("not implemented")
}

func (repo PostgresOrganizerProvisionHistoryRepository) UpdateOrganizerProvisionHistory(history businesslogic.OrganizerProvisionHistoryEntry) error {
	return errors.New("not implemented")
}
