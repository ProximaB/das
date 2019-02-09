package eventdal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/DancesportSoftware/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
	"time"
)

const (
	dasEventTable                 = "DAS.EVENT"
	dasEventColumnEventCategoryID = "CATEGORY_ID"
	dasEventColumnEventStatusID   = "EVENT_STATUS_ID"
)

// PostgresEventRepository implements IEventRepository with a Postgres database
type PostgresEventRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

// SearchEvent searches Event in a Postgres database
func (repo PostgresEventRepository) SearchEvent(criteria businesslogic.SearchEventCriteria) ([]businesslogic.Event, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
		common.ColumnPrimaryKey,
		common.COL_COMPETITION_ID,
		dasEventColumnEventCategoryID,
		common.COL_FEDERATION_ID,
		common.COL_DIVISION_ID,
		common.COL_AGE_ID,
		common.COL_PROFICIENCY_ID,
		common.COL_STYLE_ID,
		common.COL_DESCRIPTION,
		dasEventColumnEventStatusID,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated,
	)).From(dasEventTable).OrderBy(common.ColumnPrimaryKey)
	if criteria.CompetitionID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_COMPETITION_ID: criteria.CompetitionID})
	}
	if criteria.EventID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.EventID})
	}
	if criteria.FederationID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_FEDERATION_ID: criteria.FederationID})
	}
	if criteria.DivisionID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_DIVISION_ID: criteria.DivisionID})
	}
	if criteria.AgeID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_AGE_ID: criteria.AgeID})
	}
	if criteria.ProficiencyID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_PROFICIENCY_ID: criteria.ProficiencyID})
	}
	if criteria.StyleID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_STYLE_ID: criteria.StyleID})
	}
	if criteria.StatusID > 0 {
		stmt = stmt.Where(squirrel.Eq{dasEventColumnEventStatusID: criteria.StatusID})
	}
	rows, err := stmt.RunWith(repo.Database).Query()
	events := make([]businesslogic.Event, 0)
	if err != nil {
		return events, err
	}
	for rows.Next() {
		each := businesslogic.NewEvent()
		rows.Scan(
			&each.ID,
			&each.CompetitionID,
			&each.CategoryID,
			&each.FederationID,
			&each.DivisionID,
			&each.AgeID,
			&each.ProficiencyID,
			&each.StyleID,
			&each.Description,
			&each.StatusID,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		events = append(events, *each)
	}
	closeErr := rows.Close()
	if closeErr != nil {
		return events, closeErr
	}

	// get event dances
	eventDanceRepo := PostgresEventDanceRepository{
		repo.Database,
		repo.SQLBuilder,
	}
	for i := 0; i < len(events); i++ {
		eventDances, searchDanceErr := eventDanceRepo.SearchEventDance(businesslogic.SearchEventDanceCriteria{
			EventID: events[i].ID,
		})
		if searchDanceErr != nil {
			return events, err
		}
		for j := 0; j < len(eventDances); j++ {
			(events[i]).AddEventDance(eventDances[j])
			events[i].AddDance(eventDances[j].DanceID)
		}
	}
	return events, err
}

// CreateEvent creates an Event in a Postgres database
func (repo PostgresEventRepository) CreateEvent(event *businesslogic.Event) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.Insert("").
		Into(dasEventTable).
		Columns(
			common.COL_COMPETITION_ID,
			dasEventColumnEventCategoryID,
			common.COL_FEDERATION_ID,
			common.COL_DIVISION_ID,
			common.COL_AGE_ID,
			common.COL_PROFICIENCY_ID,
			common.COL_STYLE_ID,
			common.COL_DESCRIPTION,
			dasEventColumnEventStatusID,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated).
		Values(
			event.CompetitionID,
			event.CategoryID,
			event.FederationID,
			event.DivisionID,
			event.AgeID,
			event.ProficiencyID,
			event.StyleID,
			event.Description,
			event.StatusID,
			event.CreateUserID,
			event.DateTimeCreated,
			event.CreateUserID,
			time.Now()).
		Suffix("RETURNING ID")

	clause, args, err := stmt.ToSql()
	if err != nil {
		return err
	}

	tx, txErr := repo.Database.Begin()
	if txErr != nil {
		return txErr
	}
	err = tx.QueryRow(clause, args...).Scan(&event.ID)
	return tx.Commit()
}

// UpdateEvent updates an Event in a Postgres database
func (repo PostgresEventRepository) UpdateEvent(event businesslogic.Event) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.Update("").Table(dasEventTable).
		Set(dasEventColumnEventStatusID, event.StatusID).
		Where(squirrel.Eq{common.COL_COMPETITION_ID: event.CompetitionID})
	tx, txErr := repo.Database.Begin()
	if txErr != nil {
		return txErr
	}
	_, err := stmt.RunWith(repo.Database).Exec()
	err = tx.Commit()
	return err
}

// DeleteEvent deletes an Event from a Postgres database
func (repo PostgresEventRepository) DeleteEvent(event businesslogic.Event) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}
