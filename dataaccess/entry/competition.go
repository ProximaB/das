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

const (
	DAS_COMPETITION_ENTRY_TABLE                = "DAS.COMPETITION_ENTRY_ATHLETE"
	DAS_COMPETITION_ENTRY_COL_CHECKIN_IND      = "CHECKIN_IND"
	DAS_COMPETITION_ENTRY_COL_CHECKIN_DATETIME = "CHECKIN_DATETIME"
	DAS_COMPETITION_ENTRY_COL_COMPETITOR_TAG   = "LEADTAG"
)

type PostgresAthleteCompetitionEntryRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresAthleteCompetitionEntryRepository) CreateCompetitionEntry(entry * businesslogic.AthleteCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresCompetitionEntryRepository is not specified")
	}
	clause := repo.SqlBuilder.Insert("").
		Into(DAS_COMPETITION_ENTRY_TABLE).
		Columns(common.COL_COMPETITION_ID,
			common.COL_ACCOUNT_ID,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED).Values(
		entry.CompetitionID,
		entry.AthleteID,
		entry.CreateUserID,
		entry.DateTimeCreated,
		entry.UpdateUserID,
		entry.DateTimeUpdated).Suffix("RETURNING ID")

	_, err := clause.RunWith(repo.Database).Exec() // it's okay if the error is duplicate entry, since db has unique constraint on it
	return err
}

func (repo PostgresAthleteCompetitionEntryRepository) SearchCompetitionEntry(criteria businesslogic.SearchCompetitionEntryCriteria) ([]businesslogic.CompetitionEntry, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresCompetitionEntryRepository is not specified")
	}
	clause := repo.SqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s",
		common.PRIMARY_KEY,
		common.COL_COMPETITION_ID,
		common.COL_ACCOUNT_ID,
		DAS_COMPETITION_ENTRY_COL_CHECKIN_IND,
		DAS_COMPETITION_ENTRY_COL_CHECKIN_DATETIME,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED)).From(DAS_COMPETITION_ENTRY_TABLE)

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
	entries := make([]businesslogic.CompetitionEntry, 0)
	if err != nil {
		return entries, err
	}

	for rows.Next() {
		each := businesslogic.CompetitionEntry{}
		rows.Scan(
			&each.ID,
			&each.CompetitionID,
			&each.AthleteID,
			&each.CheckedIn,
			&each.CheckInDateTime,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		entries = append(entries, each)
	}
	return entries, err
}

func (repo PostgresAthleteCompetitionEntryRepository) DeleteCompetitionEntry(entry businesslogic.CompetitionEntry) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresCompetitionEntryRepository is not specified")
	}
	return errors.New("not implemented")
}

func (repo PostgresAthleteCompetitionEntryRepository) UpdateCompetitionEntry(entry businesslogic.CompetitionEntry) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresCompetitionEntryRepository is not specified")
	}
	return errors.New("not implemented")
}
