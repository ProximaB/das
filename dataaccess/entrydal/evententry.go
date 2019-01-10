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

package entrydal

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/DancesportSoftware/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
)

type PostgresAthleteEventEntryRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

func (repo PostgresAthleteEventEntryRepository) CreateAthleteEventEntry(entry *businesslogic.AthleteEventEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.Insert("").
		Columns(
			common.ColumnPrimaryKey,

			"").
		Values(
			entry.AthleteID,
			entry.CompetitionID,
			entry.EventID,
			entry.CheckedIn,
			entry.Placement,
			entry.CreateUserID,
			entry.DateTimeCreated,
			entry.UpdateUserID,
			entry.DateTimeUpdated)
}

func (repo PostgresAthleteEventEntryRepository) DeleteAthleteEventEntry(entry businesslogic.AthleteEventEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
}

func (repo PostgresAthleteEventEntryRepository) SearchAthleteEventEntry(criteria businesslogic.SearchAthleteEventEntryCriteria) ([]businesslogic.AthleteEventEntry, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
}

func (repo PostgresAthleteEventEntryRepository) UpdateAthleteEventEntry(entry businesslogic.AthleteEventEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
}

// PostgresPartnershipEventEntryRepository is a Postgres-based implementation of IPartnershipEventEntryRepository
type PostgresPartnershipEventEntryRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

const (
	dasEventCompetitiveBallroomEntryTable = "DAS.EVENT_ENTRY_PARTNERSHIP"
	leadTag                               = "LEADTAG"
)

// CreatePartnershipEventEntry creates a Partnership Event Entry in a Postgres database
func (repo PostgresPartnershipEventEntryRepository) CreatePartnershipEventEntry(entry *businesslogic.PartnershipEventEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.Insert("").Into(dasEventCompetitiveBallroomEntryTable).Columns(
		common.COL_EVENT_ID,
		common.COL_PARTNERSHIP_ID,
		leadTag,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated,
	).Values(
		entry.EventEntry.EventID,
		entry.PartnershipID,
		entry.EventEntry.Mask,
		entry.EventEntry.CreateUserID,
		entry.EventEntry.DateTimeCreated,
		entry.EventEntry.UpdateUserID,
		entry.EventEntry.DateTimeUpdated,
	).Suffix(dalutil.SQLSuffixReturningID)
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

// DeletePartnershipEventEntry deletes a Partnership Event Entry from a Postgres database
func (repo PostgresPartnershipEventEntryRepository) DeletePartnershipEventEntry(entry businesslogic.PartnershipEventEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	clause := repo.SQLBuilder.Delete("").
		From(dasEventCompetitiveBallroomEntryTable).
		Where(squirrel.Eq{common.COL_EVENT_ID: entry.EventEntry.EventID}).
		Where(squirrel.Eq{common.COL_PARTNERSHIP_ID: entry.PartnershipID})
	_, err := clause.RunWith(repo.Database).Exec()
	return err
}

// UpdatePartnershipEventEntry makes changes to a Partnership Event Entry in a Postgres database
func (repo PostgresPartnershipEventEntryRepository) UpdatePartnershipEventEntry(entry businesslogic.PartnershipEventEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// SearchPartnershipEventEntry returns CompetitiveBallroomEventEntry, which is supposed to be used by competitor only
func (repo PostgresPartnershipEventEntryRepository) SearchPartnershipEventEntry(criteria businesslogic.SearchPartnershipEventEntryCriteria) ([]businesslogic.PartnershipEventEntry, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	clause := repo.SQLBuilder.Select(
		fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s",
			common.ColumnPrimaryKey,
			common.COL_EVENT_ID,
			common.COL_PARTNERSHIP_ID,
			dasCompetitionEntryColCompetitorTag,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated)).
		From(dasEventCompetitiveBallroomEntryTable)

	if criteria.PartnershipID > 0 {
		clause = clause.Where(squirrel.Eq{common.COL_PARTNERSHIP_ID: criteria.PartnershipID})
	}
	if criteria.EventID > 0 {
		clause = clause.Where(squirrel.Eq{common.COL_EVENT_ID: criteria.EventID})
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

// PostgresAdjudicatorEventEntryRepository implements IAdjudicatorEventEntryRepository with a Postgres database
type PostgresAdjudicatorEventEntryRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

// CreateAdjudicatorEventEntry creates an Adjudicator Event Entry in a Postgres database
func (repo PostgresAdjudicatorEventEntryRepository) CreateAdjudicatorEventEntry(entry *businesslogic.AdjudicatorEventEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// DeleteAdjudicatorEventEntry deletes an Adjudicator Event Entry from a Postgres database
func (repo PostgresAdjudicatorEventEntryRepository) DeleteAdjudicatorEventEntry(entry businesslogic.AdjudicatorEventEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// SearchAdjudicatorEventEntry searches Adjudicator Event Entries in a Postgres database
func (repo PostgresAdjudicatorEventEntryRepository) SearchAdjudicatorEventEntry(criteria businesslogic.SearchAdjudicatorEventEntryCriteria) ([]businesslogic.AdjudicatorEventEntry, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return nil, errors.New("not implemented")
}

// UpdateAdjudicatorEventEntry updates an Adjudicator Event Entry in a Postgres database
func (repo PostgresAdjudicatorEventEntryRepository) UpdateAdjudicatorEventEntry(entry businesslogic.AdjudicatorEventEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// Returns CompetitiveBallroomEventEntryPublicView, which contains minimal information of the entry and is used by
// public only
/*
func GetCompetitiveBallroomEventEntrylist(criteria *businesslogic.SearchEventEntryCriteria) ([]businesslogic.EventEntryPublicView, error) {
	entries := make([]businesslogic.EventEntryPublicView, 0)

	clause := repo.SQLBuilder.Select(`ECBE.ID, ECB.ID, E.ID, C.ID, P.ID, P.LEAD, P.FOLLOW,
					AL.FIRST_NAME, AL.LAST_NAME,
					AF.FIRST_NAME, AF.LAST_NAME,
					RC.NAME, RST.NAME, RSC.NAME, RSO.NAME
			`).
		From(dasEventCompetitiveBallroomEntryTable).
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
