// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package entry

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/Masterminds/squirrel"
)

// PostgresPartnershipEventEntryRepository is a Postgres-based implementation of IPartnershipEventEntryRepository
type PostgresPartnershipEventEntryRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

const (
	DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_TABLE                             = "DAS.EVENT_ENTRY_PARTNERSHIP"
	DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_COL_COMPETITIVE_BALLROOM_EVENT_ID = "EVENT_ID"
	DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_COL_LEADTAG                       = "LEADTAG"
)

func (repo PostgresPartnershipEventEntryRepository) CreatePartnershipEventEntry(entry *businesslogic.PartnershipEventEntry) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresEventEntryRepository is not specified")
	}
	stmt := repo.SQLBuilder.Insert("").Into(DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_TABLE).Columns(
		DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_COL_COMPETITIVE_BALLROOM_EVENT_ID,
		common.COL_PARTNERSHIP_ID,
		DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_COL_LEADTAG,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED,
	).Values(
		entry.EventEntry.EventID,
		entry.PartnershipID,
		entry.EventEntry.Mask,
		entry.EventEntry.CreateUserID,
		entry.EventEntry.DateTimeCreated,
		entry.EventEntry.UpdateUserID,
		entry.EventEntry.DateTimeUpdated,
	).Suffix("RETURNING ID")
	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&entry.ID)
		tx.Commit()
	}
	return err
}

func (repo PostgresPartnershipEventEntryRepository) DeletePartnershipEventEntry(entry businesslogic.PartnershipEventEntry) error {
	clause := repo.SQLBuilder.Delete("").
		From(DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_TABLE).
		Where(squirrel.Eq{DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_COL_COMPETITIVE_BALLROOM_EVENT_ID: entry.EventEntry.EventID}).
		Where(squirrel.Eq{common.COL_PARTNERSHIP_ID: entry.PartnershipID})
	_, err := clause.RunWith(repo.Database).Exec()
	return err
}

func (repo PostgresPartnershipEventEntryRepository) UpdatePartnershipEventEntry(entry businesslogic.PartnershipEventEntry) error {
	return errors.New("not implemented")
}

// Returns CompetitiveBallroomEventEntry, which is supposed to be used by competitor only
func (repo PostgresPartnershipEventEntryRepository) SearchPartnershipEventEntry(criteria businesslogic.SearchPartnershipEventEntryCriteria) ([]businesslogic.PartnershipEventEntry, error) {
	clause := repo.SQLBuilder.Select(
		fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s",
			common.PRIMARY_KEY,
			DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_COL_COMPETITIVE_BALLROOM_EVENT_ID,
			common.COL_PARTNERSHIP_ID,
			dasCompetitionEntryColCompetitorTag,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED)).
		From(DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_TABLE)

	if criteria.PartnershipID > 0 {
		clause = clause.Where(squirrel.Eq{common.COL_PARTNERSHIP_ID: criteria.PartnershipID})
	}
	if criteria.EventID > 0 {
		clause = clause.Where(squirrel.Eq{DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_COL_COMPETITIVE_BALLROOM_EVENT_ID: criteria.EventID})
	}

	entries := make([]businesslogic.PartnershipEventEntry, 0)
	rows, err := clause.RunWith(repo.Database).Query()

	if err != nil {
		rows.Close()
		return entries, err
	}

	for rows.Next() {
		each := businesslogic.PartnershipEventEntry{}
		rows.Scan(
			&each.ID,
			&each.EventEntry.EventID,
			&each.PartnershipID,
			&each.EventEntry.Mask,
			&each.EventEntry.CreateUserID,
			&each.EventEntry.DateTimeCreated,
			&each.EventEntry.UpdateUserID,
			&each.EventEntry.DateTimeUpdated,
		)
		entries = append(entries, each)
	}
	rows.Close()
	return entries, err
}

type PostgresAdjudicatorEventEntryRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

func (repo PostgresAdjudicatorEventEntryRepository) CreateAdjudicatorEventEntry(entry *businesslogic.AdjudicatorEventEntry) error {
	return errors.New("not implemented")
}

func (repo PostgresAdjudicatorEventEntryRepository) DeleteAdjudicatorEventEntry(entry businesslogic.AdjudicatorEventEntry) error {
	return errors.New("not implemented")
}

func (repo PostgresAdjudicatorEventEntryRepository) SearchAdjudicatorEventEntry(criteria businesslogic.SearchAdjudicatorEventEntryCriteria) ([]businesslogic.AdjudicatorEventEntry, error) {
	return nil, errors.New("not implemented")
}

func (repo PostgresAdjudicatorEventEntryRepository) UpdateAdjudicatorEventEntry(entry businesslogic.AdjudicatorEventEntry) error {
	return errors.New("not implemented")
}

// Returns CompetitiveBallroomEventEntryPublicView, which contains minimal information of the entry and is used by
// public only
/*
func GetCompetitiveBallroomEventEntrylist(criteria *businesslogic.SearchEventEntryCriteria) ([]businesslogic.EventEntryPublicView, error) {
	entries := make([]businesslogic.EventEntryPublicView, 0)

	clause := repo.SqlBuilder.Select(`ECBE.ID, ECB.ID, E.ID, C.ID, P.ID, P.LEAD, P.FOLLOW,
					AL.FIRST_NAME, AL.LAST_NAME,
					AF.FIRST_NAME, AF.LAST_NAME,
					RC.NAME, RST.NAME, RSC.NAME, RSO.NAME
			`).
		From(DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_TABLE).
		Where(sq.Eq{"E.COMPETITION_ID": criteria.ID})

	if criteria.Federation > 0 {
		clause = clause.Where(sq.Eq{"ECB.FEDERATION_ID": criteria.Federation})
	}
	if criteria.Division > 0 {
		clause = clause.Where(sq.Eq{"ECB.DIVISION_ID": criteria.Division})
	}
	if criteria.Age > 0 {
		clause = clause.Where(sq.Eq{"ECB.AGE_ID": criteria.Age})
	}
	if criteria.Proficiency > 0 {
		clause = clause.Where(sq.Eq{"ECB.PROFICIENCY_ID": criteria.Proficiency})
	}

	rows, err := clause.RunWith(repo.Database).Query()

	if err != nil {
		rows.Close()
		return entries, err
	}
	for rows.Next() {

	}
	rows.Close()
	return entries, err
}
*/
