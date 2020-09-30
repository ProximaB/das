package eventdal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/dataaccess/common"
	"github.com/ProximaB/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
	"log"
)

// PostgresEventDanceRepository implements IEventDanceRepository
type PostgresEventDanceRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

const (
	DAS_EVENT_DANCES_TABLE = "DAS.EVENT_DANCES"
)

func (repo PostgresEventDanceRepository) SearchEventDance(criteria businesslogic.SearchEventDanceCriteria) ([]businesslogic.EventDance, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Select(
		fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s",
			common.ColumnPrimaryKey,
			common.COL_EVENT_ID,
			common.COL_DANCE_ID,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated),
	).From(DAS_EVENT_DANCES_TABLE).
		OrderBy(common.ColumnPrimaryKey)
	if criteria.EventID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_EVENT_ID: criteria.EventID})
	}
	rows, err := stmt.RunWith(repo.Database).Query()
	output := make([]businesslogic.EventDance, 0)
	if err != nil {
		log.Printf("[error] querying EventDances with criteria %#v: %v", criteria, err)
		return output, err
	}
	for rows.Next() {
		each := businesslogic.EventDance{}
		scanErr := rows.Scan(
			&each.ID,
			&each.EventID,
			&each.DanceID,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		if scanErr != nil {
			log.Printf("[error] scanning EventDance with %#v: %v", criteria, scanErr)
			return output, scanErr
		}
		output = append(output, each)
	}
	closeErr := rows.Close()
	return output, closeErr
}
func (repo PostgresEventDanceRepository) CreateEventDance(eventDance *businesslogic.EventDance) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Insert("").
		Into(DAS_EVENT_DANCES_TABLE).
		Columns(
			common.COL_EVENT_ID,
			common.COL_DANCE_ID,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated).
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
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}
func (repo PostgresEventDanceRepository) UpdateEventDance(eventDance businesslogic.EventDance) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}
