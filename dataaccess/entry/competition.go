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

package entry

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/Masterminds/squirrel"
)

const (
	dasCompetitionEntryTable              = "DAS.COMPETITION_ENTRY_ATHLETE"
	dasCompetitionEntryColCheckinInd      = "CHECKIN_IND"
	dasCompetitionEntryColCheckinDateTime = "CHECKIN_DATETIME"
	dasCompetitionEntryColCompetitorTag   = "LEADTAG"
)

// PostgresAthleteCompetitionEntryRepository is a Postgres-based Athlete Competition Entry Repository
type PostgresAthleteCompetitionEntryRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

// CreateAthleteCompetitionEntry creates a AthleteCompetitionEntry in the repository and update the ID of the AthleteCompetitionEntry
func (repo PostgresAthleteCompetitionEntryRepository) CreateAthleteCompetitionEntry(entry *businesslogic.AthleteCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresCompetitionEntryRepository is not specified")
	}
	clause := repo.SQLBuilder.Insert("").
		Into(dasCompetitionEntryTable).
		Columns(common.COL_COMPETITION_ID,
			common.COL_ACCOUNT_ID,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED).Values(
		entry.CompetitionEntry.CompetitionID,
		entry.AthleteID,
		entry.CompetitionEntry.CreateUserID,
		entry.CompetitionEntry.DateTimeCreated,
		entry.CompetitionEntry.UpdateUserID,
		entry.CompetitionEntry.DateTimeUpdated).Suffix("RETURNING ID")

	_, err := clause.RunWith(repo.Database).Exec() // it's okay if the error is duplicate entry, since db has unique constraint on it
	return err
}

func (repo PostgresAthleteCompetitionEntryRepository) SearchAthleteCompetitionEntry(criteria businesslogic.SearchAthleteCompetitionEntryCriteria) ([]businesslogic.AthleteCompetitionEntry, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresCompetitionEntryRepository is not specified")
	}
	clause := repo.SQLBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s",
		common.PRIMARY_KEY,
		common.COL_COMPETITION_ID,
		common.COL_ACCOUNT_ID,
		dasCompetitionEntryColCheckinInd,
		dasCompetitionEntryColCheckinDateTime,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED)).From(dasCompetitionEntryTable)

	if criteria.ID > 0 {
		clause = clause.Where(squirrel.Eq{common.PRIMARY_KEY: criteria.ID})
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

func (repo PostgresAthleteCompetitionEntryRepository) DeleteAthleteCompetitionEntry(entry businesslogic.AthleteCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresCompetitionEntryRepository is not specified")
	}
	return errors.New("not implemented")
}

func (repo PostgresAthleteCompetitionEntryRepository) UpdateAthleteCompetitionEntry(entry businesslogic.AthleteCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresCompetitionEntryRepository is not specified")
	}
	return errors.New("not implemented")
}

type PostgresPartnershipCompetitionEntryRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresPartnershipCompetitionEntryRepository) CreatePartnershipCompetitionEntry(entry *businesslogic.PartnershipCompetitionEntry) error {
	return errors.New("not implemented")
}

func (repo PostgresPartnershipCompetitionEntryRepository) DeletePartnershipCompetitionEntry(entry businesslogic.PartnershipCompetitionEntry) error {
	return errors.New("not implemented")
}

func (repo PostgresPartnershipCompetitionEntryRepository) SearchPartnershipCompetitionEntry(criteria businesslogic.SearchPartnershipCompetitionEntryCriteria) ([]businesslogic.PartnershipCompetitionEntry, error) {
	return nil, errors.New("not implemented")
}

func (repo PostgresPartnershipCompetitionEntryRepository) UpdatePartnershipCompetitionEntry(entry businesslogic.PartnershipCompetitionEntry) error {
	return errors.New("not implemented")
}

type PostgresAdjudicatorCompetitionEntryRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

func (repo PostgresAdjudicatorCompetitionEntryRepository) CreateAdjudicatorCompetitionEntry(entry *businesslogic.AdjudicatorCompetitionEntry) error {
	return errors.New("not implemented")
}

func (repo PostgresAdjudicatorCompetitionEntryRepository) DeleteAdjudicatorCompetitionEntry(entry businesslogic.AdjudicatorCompetitionEntry) error {
	return errors.New("not implemented")
}

func (repo PostgresAdjudicatorCompetitionEntryRepository) SearchAdjudicatorCompetitionEntry(criteria businesslogic.SearchAdjudicatorCompetitionEntryCriteria) ([]businesslogic.AdjudicatorCompetitionEntry, error) {
	return nil, errors.New("not implemented")
}

func (repo PostgresAdjudicatorCompetitionEntryRepository) UpdateAdjudicatorCompetitionEntry(entry businesslogic.AdjudicatorCompetitionEntry) error {
	return errors.New("not implemented")
}
