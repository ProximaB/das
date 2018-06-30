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

package event

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/Masterminds/squirrel"
	"time"
)

const (
	DAS_EVENT_TABLE                 = "DAS.EVENT"
	DAS_EVENT_COL_EVENT_CATEGORY_ID = "EVENT_CATEGORY_ID"
	DAS_EVENT_COL_EVENT_STATUS_ID   = "EVENT_STATUS_ID"
)

type PostgresEventRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresEventRepository) SearchEvent(criteria businesslogic.SearchEventCriteria) ([]businesslogic.Event, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresEventRepository is not specified")
	}
	stmt := repo.SqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
		common.PRIMARY_KEY,
		common.COL_COMPETITION_ID,
		DAS_EVENT_COL_EVENT_CATEGORY_ID,
		common.COL_FEDERATION_ID,
		common.COL_DIVISION_ID,
		common.COL_AGE_ID,
		common.COL_PROFICIENCY_ID,
		common.COL_STYLE_ID,
		common.COL_DESCRIPTION,
		DAS_EVENT_COL_EVENT_STATUS_ID,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED,
	)).From(DAS_EVENT_TABLE).OrderBy(common.PRIMARY_KEY)
	if criteria.CompetitionID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_COMPETITION_ID: criteria.CompetitionID})
	}
	if criteria.EventID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.PRIMARY_KEY: criteria.EventID})
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
		stmt = stmt.Where(squirrel.Eq{DAS_EVENT_COL_EVENT_STATUS_ID: criteria.StatusID})
	}
	rows, err := stmt.RunWith(repo.Database).Query()
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
		events = append(events, each)
	}
	rows.Close()
	return events, err
}

func (repo PostgresEventRepository) CreateEvent(event *businesslogic.Event) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresEventRepository is not specified")
	}
	stmt := repo.SqlBuilder.Insert("").
		Into(DAS_EVENT_TABLE).
		Columns(
			common.COL_COMPETITION_ID,
			DAS_EVENT_COL_EVENT_CATEGORY_ID,
			common.COL_FEDERATION_ID,
			common.COL_DIVISION_ID,
			common.COL_AGE_ID,
			common.COL_PROFICIENCY_ID,
			common.COL_STYLE_ID,
			common.COL_DESCRIPTION,
			DAS_EVENT_COL_EVENT_STATUS_ID,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED).
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

	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		tx.QueryRow(clause, args...).Scan(&event.ID)
		tx.Commit()
		return nil
	}
}

func (repo PostgresEventRepository) UpdateEvent(event businesslogic.Event) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresEventRepository is not specified")
	}
	stmt := repo.SqlBuilder.Update("").Table(DAS_EVENT_TABLE).
		Set(DAS_EVENT_COL_EVENT_STATUS_ID, event.StatusID).
		Where(squirrel.Eq{common.COL_COMPETITION_ID: event.CompetitionID})
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err := stmt.RunWith(repo.Database).Exec()
		err = tx.Commit()
		return err
	}
}

func (repo PostgresEventRepository) DeleteEvent(event businesslogic.Event) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresEventRepository is not specified")
	}
	return errors.New("not implemented")
}
