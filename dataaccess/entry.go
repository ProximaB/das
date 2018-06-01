package dataaccess

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/dataaccess/common"
	"database/sql"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

type PostgresEventEntryRepository struct {
	database   *sql.DB
	sqlBuilder sq.StatementBuilderType
}

const (
	DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_TABLE                             = "DAS.EVENT_COMPETITIVE_BALLROOM_ENTRY"
	DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_COL_COMPETITIVE_BALLROOM_EVENT_ID = "COMPETITIVE_BALLROOM_EVENT_ID"
	DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_COL_LEADTAG                       = "LEADTAG"
)

func (repo PostgresEventEntryRepository) CreateEventEntry(entry businesslogic.EventEntry) error {
	clause := repo.sqlBuilder.Insert("").
		Into(DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_TABLE).
		Columns(DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_COL_COMPETITIVE_BALLROOM_EVENT_ID,
			common.COL_PARTNERSHIP_ID,
			DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_COL_LEADTAG,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED).Values(
		entry.EventID,
		entry.PartnershipID,
		entry.CompetitorTag,
		entry.CreateUserID,
		entry.DateTimeCreated,
		entry.UpdateUserID,
		entry.DateTimeUpdated,
	)
	_, err := clause.RunWith(repo.database).Exec()
	return err
}

func (repo PostgresEventEntryRepository) DeleteEventEntry(entry businesslogic.EventEntry) error {
	clause := repo.sqlBuilder.Delete("").
		From(DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_TABLE).
		Where(sq.Eq{DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_COL_COMPETITIVE_BALLROOM_EVENT_ID: entry.EventID}).
		Where(sq.Eq{common.COL_PARTNERSHIP_ID: entry.PartnershipID})
	_, err := clause.RunWith(repo.database).Exec()
	return err
}

func (repo PostgresEventEntryRepository) UpdateEventEntry(entry businesslogic.EventEntry) error {
	return errors.New("not implemented")
}

// Returns CompetitiveBallroomEventEntry, which is supposed to be used by competitor only
func (repo PostgresEventEntryRepository) SearchEventEntry(criteria *businesslogic.SearchEventEntryCriteria) ([]businesslogic.EventEntry, error) {
	clause := repo.sqlBuilder.Select(
		fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s",
			common.PRIMARY_KEY,
			DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_COL_COMPETITIVE_BALLROOM_EVENT_ID,
			common.COL_PARTNERSHIP_ID,
			DAS_COMPETITION_ENTRY_COL_COMPETITOR_TAG,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED)).
		From(DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_TABLE)

	if criteria.PartnershipID > 0 {
		clause = clause.Where(sq.Eq{common.COL_PARTNERSHIP_ID: criteria.PartnershipID})
	}
	if criteria.EventID > 0 {
		clause = clause.Where(sq.Eq{DAS_EVENT_COMPETITIVE_BALLROOM_ENTRY_COL_COMPETITIVE_BALLROOM_EVENT_ID: criteria.EventID})
	}

	entries := make([]businesslogic.EventEntry, 0)
	rows, err := clause.RunWith(repo.database).Query()

	if err != nil {
		rows.Close()
		return entries, err
	}

	for rows.Next() {
		each := businesslogic.EventEntry{}
		rows.Scan(
			&each.ID,
			&each.EventID,
			&each.PartnershipID,
			&each.CompetitorTag,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		entries = append(entries, each)
	}
	rows.Close()
	return entries, err
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
		Where(sq.Eq{"E.COMPETITION_ID": criteria.CompetitionID})

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
