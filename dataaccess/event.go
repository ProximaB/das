package dataaccess

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/dataaccess/common"
	"database/sql"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"time"
)

const (
	DAS_EVENT_TABLE                 = "DAS.EVENT"
	DAS_EVENT_COL_EVENT_CATEGORY_ID = "EVENT_CATEGORY_ID"
	DAS_EVENT_COL_EVENT_STATUS_ID   = "EVENT_STATUS_ID"
)

const (
	DAS_EVENT_COMPETITIVE_BALLROOM_TABLE_PREFIX = "ECB"
	DAS_EVENT_COMPETITIVE_BALLROOM_TABLE        = "DAS.EVENT_COMPETITIVE_BALLROOM"
)

type PostgresEventRepository struct {
	database   *sql.DB
	sqlBuilder sq.StatementBuilderType
}

func (repo PostgresEventRepository) SearchEvent(criteria *businesslogic.SearchEventCriteria) ([]businesslogic.Event, error) {
	stmt := repo.sqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s",
		common.PRIMARY_KEY,
		common.COL_COMPETITION_ID,
		DAS_EVENT_COL_EVENT_CATEGORY_ID,
		common.COL_DESCRIPTION,
		DAS_EVENT_COL_EVENT_STATUS_ID,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED,
	)).From(DAS_EVENT_TABLE).OrderBy(common.PRIMARY_KEY)
	if criteria.CompetitionID > 0 {
		stmt = stmt.Where(sq.Eq{common.COL_COMPETITION_ID: criteria.CompetitionID})
	}
	if criteria.EventID > 0 {
		stmt = stmt.Where(sq.Eq{common.PRIMARY_KEY: criteria.EventID})
	}
	rows, err := stmt.RunWith(repo.database).Query()
	events := make([]businesslogic.Event, 0)
	if err != nil {
		return events, err
	}
	for rows.Next() {
		each := businesslogic.Event{}
		rows.Scan(
			&each.ID,
			&each.CompetitionID,
			&each.CategoryID,
			&each.Description,
			&each.StatusID,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		events = append(events, each)
	}
	rows.Close()
	return events, err
}

/*
func (repo PostgresEventRepository) SearchCompetitiveBallroomEvents(
	db *sql.DB,
	criteria businesslogic.SearchEventCriteria) (
	[]businesslogic.Event, error) {
	selCompetitiveBallroomEventsSQL := SQLBUILDER.Select(fmt.Sprintf(
		`%s.%s, %s.%s, %s.%s, %s.%s, %s.%s, %s.%s, %s.%s, %s.%s, %s.%s, %s.%s, %s.%s`,
		DAS_EVENT_COMPETITIVE_BALLROOM_TABLE_PREFIX, common.PRIMARY_KEY,
		DAS_EVENT_COMPETITIVE_BALLROOM_TABLE_PREFIX, common.COL_EVENT_ID,
		DAS_EVENT_COMPETITIVE_BALLROOM_TABLE_PREFIX, common.COL_FEDERATION_ID,
		DAS_EVENT_COMPETITIVE_BALLROOM_TABLE_PREFIX, common.COL_DIVISION_ID,
		DAS_EVENT_COMPETITIVE_BALLROOM_TABLE_PREFIX, common.COL_AGE_ID,
		DAS_EVENT_COMPETITIVE_BALLROOM_TABLE_PREFIX, common.COL_PROFICIENCY_ID,
		DAS_EVENT_COMPETITIVE_BALLROOM_TABLE_PREFIX, common.COL_STYLE_ID,
		DAS_EVENT_COMPETITIVE_BALLROOM_TABLE_PREFIX, common.COL_CREATE_USER_ID,
		DAS_EVENT_COMPETITIVE_BALLROOM_TABLE_PREFIX, common.COL_DATETIME_CREATED,
		DAS_EVENT_COMPETITIVE_BALLROOM_TABLE_PREFIX, common.COL_UPDATE_USER_ID,
		DAS_EVENT_COMPETITIVE_BALLROOM_TABLE_PREFIX, common.COL_DATETIME_UPDATED,
	)).From("DAS.EVENT_COMPETITIVE_BALLROOM ECB").
		Join("DAS.EVENT DE ON ECB.EVENT_ID = DE.ID").
		Join("DAS.FEDERATION DF ON ECB.FEDERATION_ID = DF.ID").
		Join("DAS.DIVISION DD ON ECB.DIVISION_ID = DD.ID").
		Join("DAS.AGE DA ON ECB.AGE_ID = DA.ID").
		Join("DAS.PROFICIENCY DP ON ECB.PROFICIENCY_ID = DP.ID").
		Join("DAS.STYLE DS ON ECB.STYLE_ID = DS.ID").
		Join("DAS.EVENT E ON ECB.EVENT_ID = E.ID")

	if criteria.EventID > 0 {
		selCompetitiveBallroomEventsSQL = selCompetitiveBallroomEventsSQL.Where(sq.Eq{"ECB.ID": criteria.EventID})
	}
	if criteria.CompetitionID > 0 {
		selCompetitiveBallroomEventsSQL = selCompetitiveBallroomEventsSQL.Where(sq.Eq{"DE.COMPETITION_ID": criteria.CompetitionID})
	}
	if criteria.FederationID > 0 {
		selCompetitiveBallroomEventsSQL = selCompetitiveBallroomEventsSQL.Where(sq.Eq{"ECB.FEDERATION_ID": criteria.FederationID})
	}
	if criteria.DivisionID > 0 {
		selCompetitiveBallroomEventsSQL = selCompetitiveBallroomEventsSQL.Where(sq.Eq{"ECB.DIVISION_ID": criteria.DivisionID})
	}
	if criteria.AgeID > 0 {
		selCompetitiveBallroomEventsSQL = selCompetitiveBallroomEventsSQL.Where(sq.Eq{"ECB.AGE_ID": criteria.AgeID})
	}
	if criteria.ProficiencyID > 0 {
		selCompetitiveBallroomEventsSQL = selCompetitiveBallroomEventsSQL.Where(sq.Eq{"ECB.PROFICIENCY_ID": criteria.ProficiencyID})
	}
	if criteria.ID > 0 {
		selCompetitiveBallroomEventsSQL = selCompetitiveBallroomEventsSQL.Where(sq.Eq{"ECB.STYLE_ID": criteria.ID})
	}
	if criteria.StatusID > 0 {
		selCompetitiveBallroomEventsSQL = selCompetitiveBallroomEventsSQL.Where(sq.Eq{"E.EVENT_STATUS_ID": criteria.StatusID})
	}
	rows, err := selCompetitiveBallroomEventsSQL.RunWith(db).Query()
	events := make([]businesslogic.Event, 0)
	if err != nil {
		return events, err
	}
	for rows.Next() {
		each := businesslogic.NewEvent()

		rows.Scan(
			&each.ID,
			&each.FederationID,
			&each.DivisionID,
			&each.AgeID,
			&each.ProficiencyID,
			&each.ID,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		events = append(events, each)
	}

	// select dances for each event
	for _, each := range events {
		selCompetitiveBallroomEventDancesSQL := SQLBUILDER.
			Select("DANCE_ID").
			From(DAS_EVENT_DANCES_TABLE).
			Where(sq.Eq{"COMPETITIVE_BALLROOM_EVENT_ID": each.ID})
		danceRows, _ := selCompetitiveBallroomEventDancesSQL.RunWith(db).Query()

		for danceRows.Next() {
			eachDance := 0
			danceRows.Scan(&eachDance)
			each.AddDance(eachDance)
		}
		danceRows.Close()
	}
	rows.Close()
	return events, err
}
*/
func (repo PostgresEventRepository) CreateEvent(event *businesslogic.Event) error {
	stmt := repo.sqlBuilder.Insert("").
		Into(DAS_EVENT_TABLE).
		Columns(
			common.COL_COMPETITION_ID,
			DAS_EVENT_COL_EVENT_CATEGORY_ID,
			common.COL_DESCRIPTION,
			DAS_EVENT_COL_EVENT_STATUS_ID,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED).
		Values(
			event.CompetitionID,
			event.CategoryID,
			event.Description,
			event.StatusID,
			event.CreateUserID,
			event.DateTimeCreated,
			event.CreateUserID,
			time.Now()).
		Suffix(fmt.Sprintf("RETURNING %s", common.PRIMARY_KEY))

	clause, args, err := stmt.ToSql()
	if err != nil {
		return err
	}

	if tx, txErr := repo.database.Begin(); txErr != nil {
		return txErr
	} else {
		tx.QueryRow(clause, args...).Scan(&event.ID)
		tx.Commit()
		return nil
	}
}

func (repo PostgresEventRepository) UpdateEvent(event businesslogic.Event) error {
	return errors.New("not implemented")
}

func (repo PostgresEventRepository) DeleteEvent(event businesslogic.Event) error {
	return errors.New("not implemented")
}

/*
func CreateCompetitiveBallroomEvent(db *sql.DB, event businesslogic.Event, ballroom *businesslogic.Event) error {
	stmt := SQLBUILDER.Insert("").
		Into(DAS_EVENT_COMPETITIVE_BALLROOM_TABLE).
		Columns(
			common.COL_EVENT_ID,
			common.COL_FEDERATION_ID,
			common.COL_DIVISION_ID,
			common.COL_AGE_ID,
			common.COL_PROFICIENCY_ID,
			common.COL_STYLE_ID,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED).
		Values(
			event.ID,
			ballroom.FederationID,
			ballroom.DivisionID,
			ballroom.AgeID,
			ballroom.ProficiencyID,
			ballroom.ID,
			ballroom.CreateUserID,
			ballroom.DateTimeCreated,
			ballroom.CreateUserID,
			time.Now()).
		Suffix(fmt.Sprintf("RETURNING %s", common.PRIMARY_KEY))

	clause, args, err := stmt.ToSql()
	if err != nil {
		return err
	}

	if tx, txErr := db.Begin(); txErr != nil {
		return txErr
	} else {
		tx.QueryRow(clause, args...).Scan(&ballroom.ID)
		tx.Commit()
		return nil
	}
}

func CreateCompetitiveBallroomEventDances(db *sql.DB, event businesslogic.Event) error {
	for dance := range event.GetDances() {
		stmt := SQLBUILDER.Insert("").
			Into(DAS_EVENT_DANCES_TABLE).
			Columns(
				"COMPETITIVE_BALLROOM_EVENT_ID",
				common.COL_DANCE_ID,
				common.COL_CREATE_USER_ID,
				common.COL_DATETIME_CREATED,
				common.COL_UPDATE_USER_ID,
				common.COL_DATETIME_UPDATED).
			Values(
				event.ID,
				dance,
				event.CreateUserID,
				event.DateTimeCreated,
				event.CreateUserID,
				time.Now())
		_, err := stmt.RunWith(db).Exec()
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateCompetitionEventStatus(db *sql.DB, competitionID int, statusID int) error {
	stmt := SQLBUILDER.Update("").Table(DAS_EVENT_TABLE).Set(DAS_EVENT_COL_EVENT_STATUS_ID, statusID).Where(sq.Eq{common.COL_COMPETITION_ID: competitionID})
	_, err := stmt.RunWith(db).Exec()
	return err
}
*/
