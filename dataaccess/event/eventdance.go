package event

import (
	"database/sql"
	"errors"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/Masterminds/squirrel"
)

type PostgresEventDanceRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

const (
	DAS_EVENT_DANCES_TABLE = "DAS.EVENT_DANCES"
)

func (repo PostgresEventDanceRepository) SearchEventDance(criteria businesslogic.SearchEventDanceCriteria) ([]businesslogic.EventDance, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresEventDanceRepository is not specified")
	}
	return nil, errors.New("not implemented")
}
func (repo PostgresEventDanceRepository) CreateEventDance(eventDance *businesslogic.EventDance) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresEventDanceRepository is not specified")
	}
	stmt := repo.SqlBuilder.Insert("").
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
			eventDance.DateTimeUpdated,
		).Suffix("RETURNING ID")
	clause, args, err := stmt.ToSql()
	if err != nil {
		return err
	}
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		tx.QueryRow(clause, args...).Scan(&eventDance.ID)
		tx.Commit()
		return err
	}
}
func (repo PostgresEventDanceRepository) DeleteEventDance(eventDance businesslogic.EventDance) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresEventDanceRepository is not specified")
	}
	return errors.New("not implemented")
}
func (repo PostgresEventDanceRepository) UpdateEventDance(eventDance businesslogic.EventDance) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresEventDanceRepository is not specified")
	}
	return errors.New("not implemented")
}
