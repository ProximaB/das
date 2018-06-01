package dataaccess

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/dataaccess/common"
	"database/sql"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

const (
	DAS_COMPETITION_ENTRY_TABLE                = "DAS.COMPETITION_ENTRY"
	DAS_COMPETITION_ENTRY_COL_CHECKIN_IND      = "CHECKIN_IND"
	DAS_COMPETITION_ENTRY_COL_CHECKIN_DATETIME = "CHECKIN_DATETIME"
	DAS_COMPETITION_ENTRY_COL_COMPETITOR_TAG   = "LEADTAG"
)

type PostgresCompetitionEntryRepository struct {
	database   *sql.DB
	sqlBuilder sq.StatementBuilderType
}

func (repo PostgresCompetitionEntryRepository) CreateCompetitionEntry(entry businesslogic.CompetitionEntry) error {
	clause := repo.sqlBuilder.Insert("").
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
		entry.DateTimeUpdated)

	_, err := clause.RunWith(repo.database).Exec() // it's okay if the error is duplicate entry, since db has unique constraint on it
	return err
}

func (repo PostgresCompetitionEntryRepository) SearchCompetitionEntry(criteria *businesslogic.SearchCompetitionEntryCriteria) ([]businesslogic.CompetitionEntry, error) {
	clause := repo.sqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s",
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
		clause = clause.Where(sq.Eq{common.PRIMARY_KEY: criteria.ID})
	}
	if criteria.AthleteID > 0 {
		clause = clause.Where(sq.Eq{common.COL_ACCOUNT_ID: criteria.AthleteID})
	}
	if criteria.CompetitionID > 0 {
		clause = clause.Where(sq.Eq{common.COL_COMPETITION_ID: criteria.CompetitionID})
	}

	rows, err := clause.RunWith(repo.database).Query()
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

func (repo PostgresCompetitionEntryRepository) DeleteCompetitionEntry(entry businesslogic.CompetitionEntry) error {
	return errors.New("not implemented")
}

func (repo PostgresCompetitionEntryRepository) UpdateCompetitionEntry(entry businesslogic.CompetitionEntry) error {
	return errors.New("not implemented")
}
