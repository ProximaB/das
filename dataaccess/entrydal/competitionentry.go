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

const (
	dasAthleteCompetitionEntryTable       = "DAS.COMPETITION_ENTRY_ATHLETE"
	dasPartnershipCompetitionEntryTable   = "DAS.COMPETITION_ENTRY_PARTNERSHIP"
	dasAdjudicatorCompetitionEntryTable   = "DAS.COMPETITION_ENTRY_ADJUDICATOR"
	dasScrutineerCompetitionEntryTable    = "DAS.COMPETITION_ENTRY_SCRUTINEER"
	dasCompetitionEntryColCheckinInd      = "CHECKIN_IND"
	dasCompetitionEntryColCheckinDateTime = "CHECKIN_DATETIME"
	dasCompetitionEntryColCompetitorTag   = "LEADTAG"
)

// PostgresAthleteCompetitionEntryRepository is a Postgres-based Athlete Competition Entry Repository
type PostgresAthleteCompetitionEntryRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

// CreateAthleteCompetitionEntry creates an AthleteCompetitionEntry in a Postgres database
func (repo PostgresAthleteCompetitionEntryRepository) CreateAthleteCompetitionEntry(entry *businesslogic.AthleteCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.Insert("").
		Into(dasAthleteCompetitionEntryTable).
		Columns(
			common.COL_COMPETITION_ID,
			common.COL_ACCOUNT_ID,
			dasCompetitionEntryColCheckinInd,
			dasCompetitionEntryColCheckinDateTime,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated).
		Values(
			entry.CompetitionEntry.CompetitionID,
			entry.AthleteID,
			entry.CompetitionEntry.CheckInIndicator,
			entry.CompetitionEntry.DateTimeCheckIn,
			entry.CompetitionEntry.CreateUserID,
			entry.CompetitionEntry.DateTimeCreated,
			entry.CompetitionEntry.UpdateUserID,
			entry.CompetitionEntry.DateTimeUpdated).
		Suffix("RETURNING ID")

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

// SearchAthleteCompetitionEntry searches AthleteCompetitionEntry in a Postgres database
func (repo PostgresAthleteCompetitionEntryRepository) SearchAthleteCompetitionEntry(criteria businesslogic.SearchAthleteCompetitionEntryCriteria) ([]businesslogic.AthleteCompetitionEntry, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresCompetitionEntryRepository is not specified")
	}
	clause := repo.SQLBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s",
		common.ColumnPrimaryKey,
		common.COL_COMPETITION_ID,
		common.COL_ACCOUNT_ID,
		dasCompetitionEntryColCheckinInd,
		dasCompetitionEntryColCheckinDateTime,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated)).From(dasAthleteCompetitionEntryTable)

	if criteria.ID > 0 {
		clause = clause.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.ID})
	}
	if criteria.AthleteID > 0 {
		clause = clause.Where(squirrel.Eq{common.COL_ACCOUNT_ID: criteria.AthleteID})
	}
	if criteria.CompetitionID > 0 {
		clause = clause.Where(squirrel.Eq{common.COL_COMPETITION_ID: criteria.CompetitionID})
	}

	rows, err := clause.RunWith(repo.Database).Query()
	entries := make([]businesslogic.AthleteCompetitionEntry, 0)
	if err != nil {
		return entries, err
	}

	for rows.Next() {
		each := businesslogic.AthleteCompetitionEntry{
			CompetitionEntry: businesslogic.CompetitionEntry{},
		}
		rows.Scan(
			&each.ID,
			&each.CompetitionEntry.CompetitionID,
			&each.AthleteID,
			&each.CompetitionEntry.CheckInIndicator,
			&each.CompetitionEntry.DateTimeCheckIn,
			&each.CompetitionEntry.CreateUserID,
			&each.CompetitionEntry.DateTimeCreated,
			&each.CompetitionEntry.UpdateUserID,
			&each.CompetitionEntry.DateTimeUpdated,
		)
		entries = append(entries, each)
	}
	return entries, err
}

// DeleteAthleteCompetitionEntry deletes an AthleteCompetitionEntry from a Postgres database
func (repo PostgresAthleteCompetitionEntryRepository) DeleteAthleteCompetitionEntry(entry businesslogic.AthleteCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// UpdateAthleteCompetitionEntry updates an AthleteCompetitionEntry from a Postgres database
func (repo PostgresAthleteCompetitionEntryRepository) UpdateAthleteCompetitionEntry(entry businesslogic.AthleteCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// PostgresPartnershipCompetitionEntryRepository implements a IPartnershipCompetitionEntryRepository with Postgres database
type PostgresPartnershipCompetitionEntryRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

// CreatePartnershipCompetitionEntry creates a PartnershipCompetitionEntry in a Postgres database
func (repo PostgresPartnershipCompetitionEntryRepository) CreatePartnershipCompetitionEntry(entry *businesslogic.PartnershipCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.
		Insert("").
		Into(dasPartnershipCompetitionEntryTable).
		Columns(
			common.COL_COMPETITION_ID,
			common.COL_PARTNERSHIP_ID,
			dasCompetitionEntryColCheckinInd,
			dasCompetitionEntryColCheckinDateTime,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated).
		Values(
			entry.CompetitionEntry.CompetitionID,
			entry.PartnershipID,
			entry.CompetitionEntry.CheckInIndicator,
			entry.CompetitionEntry.DateTimeCheckIn,
			entry.CompetitionEntry.CreateUserID,
			entry.CompetitionEntry.DateTimeCreated,
			entry.CompetitionEntry.UpdateUserID,
			entry.CompetitionEntry.DateTimeUpdated).
		Suffix(dalutil.SQLSuffixReturningID)

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

// DeletePartnershipCompetitionEntry deletes a PartnershipCompetitionEntry from a Postgres database
func (repo PostgresPartnershipCompetitionEntryRepository) DeletePartnershipCompetitionEntry(entry businesslogic.PartnershipCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// SearchPartnershipCompetitionEntry searches PartnershipCompetitionEntry in a Postgres database
func (repo PostgresPartnershipCompetitionEntryRepository) SearchPartnershipCompetitionEntry(criteria businesslogic.SearchPartnershipCompetitionEntryCriteria) ([]businesslogic.PartnershipCompetitionEntry, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return nil, errors.New("not implemented")
}

// UpdatePartnershipCompetitionEntry updates a PartnershipCompetitionEntry in a Postgres database
func (repo PostgresPartnershipCompetitionEntryRepository) UpdatePartnershipCompetitionEntry(entry businesslogic.PartnershipCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// PostgresAdjudicatorCompetitionEntryRepository implements the IAdjudicatorCompetitionEntryRepository with a Postgres database
type PostgresAdjudicatorCompetitionEntryRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

// CreateAdjudicatorCompetitionEntry creates an AdjudicatorCompetitionEntry in a Postgres database
func (repo PostgresAdjudicatorCompetitionEntryRepository) CreateAdjudicatorCompetitionEntry(entry *businesslogic.AdjudicatorCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// DeleteAdjudicatorCompetitionEntry deletes an AdjudicatorCompetitionEntry from a Postgres database
func (repo PostgresAdjudicatorCompetitionEntryRepository) DeleteAdjudicatorCompetitionEntry(entry businesslogic.AdjudicatorCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// SearchAdjudicatorCompetitionEntry searches AdjudicatorCompetitionEntry in a Postgres database
func (repo PostgresAdjudicatorCompetitionEntryRepository) SearchAdjudicatorCompetitionEntry(criteria businesslogic.SearchAdjudicatorCompetitionEntryCriteria) ([]businesslogic.AdjudicatorCompetitionEntry, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return nil, errors.New("not implemented")
}

// UpdateAdjudicatorCompetitionEntry updates an AdjudicatorCompetitionEntry in a Postgres database
func (repo PostgresAdjudicatorCompetitionEntryRepository) UpdateAdjudicatorCompetitionEntry(entry businesslogic.AdjudicatorCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}
