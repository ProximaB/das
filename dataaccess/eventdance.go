package dataaccess

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/dataaccess/common"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
)

type PostgresEventDanceRepository struct {
	database   *sql.DB
	sqlBuilder sq.StatementBuilderType
}

const (
	DAS_EVENT_DANCES_TABLE = "DAS.EVENT_DANCES"
)

func (repo PostgresEventDanceRepository) SearchEventDance(criteria *businesslogic.SearchEventDanceCriteria) ([]businesslogic.EventDance, error) {
	return nil, errors.New("not implemented")
}
func (repo PostgresEventDanceRepository) CreateEventDance(eventDance *businesslogic.EventDance) error {
	stmt := repo.sqlBuilder.Insert("").
		Into(DAS_EVENT_DANCES_TABLE).
		Columns(
			common.COL_EVENT_ID,
			common.COL_DANCE_ID,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED).
		Values(
			eventDance.EventID,
			eventDance.DanceID,
			eventDance.CreateUserID,
			eventDance.DateTimeCreated,
			eventDance.CreateUserID,
			eventDance.DateTimeUpdated)
	_, err := stmt.RunWith(repo.database).Exec()
	return err
}
func (repo PostgresEventDanceRepository) DeleteEventDance(eventDance businesslogic.EventDance) error {
	return errors.New("not implemented")
}
func (repo PostgresEventDanceRepository) UpdateEventDance(eventDance businesslogic.EventDance) error {
	return errors.New("not implemented")
}
