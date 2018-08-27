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

package competition

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
	DAS_COMPETITION_TABLE              = "DAS.COMPETITION"
	DAS_COMPETITION_COL_STATUS_ID      = "STATUS_ID"
	DAS_COMPETITION_COL_WEBSITE        = "WEBSITE"
	DAS_COMPETITION_COL_ADDRESS        = "ADDRESS"
	DAS_COMPETITION_COL_DATETIME_START = "DATETIME_START"
	DAS_COMPETITION_COL_DATETIME_END   = "DATETIME_END"
	DAS_COMPETITION_COL_CONTACT_NAME   = "CONTACT_NAME"
	DAS_COMPETITION_COL_CONTACT_PHONE  = "CONTACT_PHONE"
	DAS_COMPETITION_COL_CONTACT_EMAIL  = "CONTACT_EMAIL"
	DAS_COMPETITION_COL_ATTENDANCE     = "ATTENDANCE"
)

/*
func UpdateCompetitionAttendance(db *sql.DB, compID int) error {
	var attendance = 0
	getAttendanceClause := SQLBUILDER.Select("COUNT(ID)").From(DAS_COMPETITION_ENTRY_TABLE).Where(sq.Eq{common.COL_COMPETITION_ID: compID})
	getAttendanceClause.RunWith(db).QueryRow().Scan(&attendance)
	updateAttendanceClause := SQLBUILDER.Update(DAS_COMPETITION_TABLE).Set(DAS_COMPETITION_COL_ATTENDANCE, attendance)
	_, err := updateAttendanceClause.RunWith(db).Exec()
	return err
}
*/

type PostgresCompetitionRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresCompetitionRepository) CreateCompetition(competition *businesslogic.Competition) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresCompetitionRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Insert("").
		Into(DAS_COMPETITION_TABLE).
		Columns(
			common.COL_FEDERATION_ID,
			common.COL_NAME,
			DAS_COMPETITION_COL_WEBSITE,
			DAS_COMPETITION_COL_STATUS_ID,
			common.COL_COUNTRY_ID,
			common.COL_STATE_ID,
			common.COL_CITY_ID,
			DAS_COMPETITION_COL_ADDRESS,
			DAS_COMPETITION_COL_DATETIME_START,
			DAS_COMPETITION_COL_DATETIME_END,
			DAS_COMPETITION_COL_CONTACT_NAME,
			DAS_COMPETITION_COL_CONTACT_PHONE,
			DAS_COMPETITION_COL_CONTACT_EMAIL,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated).
		Values(competition.FederationID,
			competition.Name,
			competition.Website,
			competition.GetStatus(),
			competition.Country.ID,
			competition.State.ID,
			competition.City.ID,
			competition.Street,
			competition.StartDateTime,
			competition.EndDateTime,
			competition.ContactName,
			competition.ContactPhone,
			competition.ContactEmail,
			competition.CreateUserID,
			competition.DateTimeCreated,
			competition.UpdateUserID,
			competition.DateTimeUpdated,
		).Suffix("RETURNING ID")
	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&competition.ID)
		err = tx.Commit()
		return err
	}
}

func (repo PostgresCompetitionRepository) UpdateCompetition(competition businesslogic.Competition) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresCompetitionRepository is not specified")
	}
	stmt := repo.SqlBuilder.Update("").Table(DAS_COMPETITION_TABLE)
	if competition.ID > 0 {
		stmt = stmt.Set(common.COL_NAME, competition.Name).
			Set(common.COL_WEBSITE, competition.Website).
			Set(DAS_COMPETITION_COL_STATUS_ID, competition.GetStatus()).
			Set(DAS_COMPETITION_COL_DATETIME_START, competition.StartDateTime).
			Set(DAS_COMPETITION_COL_DATETIME_END, competition.EndDateTime).
			Set(DAS_COMPETITION_COL_ADDRESS, competition.Street).
			Set(DAS_COMPETITION_COL_CONTACT_NAME, competition.ContactName).
			Set(DAS_COMPETITION_COL_CONTACT_EMAIL, competition.ContactEmail).
			Set(DAS_COMPETITION_COL_CONTACT_PHONE, competition.ContactPhone).
			Set(common.ColumnDateTimeUpdated, time.Now())
	}
	stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: competition.ID})

	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err := stmt.RunWith(repo.Database).Exec()
		tx.Commit()
		return err
	}
}

func (repo PostgresCompetitionRepository) DeleteCompetition(competition businesslogic.Competition) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresCompetitionRepository is not specified")
	}
	stmt := repo.SqlBuilder.Delete("").From(DAS_COMPETITION_TABLE)
	if competition.ID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: competition.ID})
	}

	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		tx.Commit()
	}
	return err
}

func (repo PostgresCompetitionRepository) SearchCompetition(criteria businesslogic.SearchCompetitionCriteria) ([]businesslogic.Competition, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresCompetitionRepository is not specified")
	}
	stmt := repo.SqlBuilder.Select(fmt.Sprintf(
		"%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
		common.ColumnPrimaryKey,
		common.COL_FEDERATION_ID,
		common.COL_NAME,
		common.COL_ADDRESS,
		common.COL_CITY_ID,
		common.COL_STATE_ID,
		common.COL_COUNTRY_ID,
		DAS_COMPETITION_COL_DATETIME_START,
		DAS_COMPETITION_COL_DATETIME_END,
		DAS_COMPETITION_COL_CONTACT_NAME,
		DAS_COMPETITION_COL_CONTACT_PHONE,
		DAS_COMPETITION_COL_CONTACT_EMAIL,
		DAS_COMPETITION_COL_WEBSITE,
		DAS_COMPETITION_COL_STATUS_ID,
		DAS_COMPETITION_COL_ATTENDANCE,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated),
	).From(DAS_COMPETITION_TABLE).
		OrderBy(DAS_COMPETITION_COL_DATETIME_START)

	if criteria.ID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.ID})
	}
	if len(criteria.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_NAME: criteria.Name})
	}
	if criteria.FederationID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_FEDERATION_ID: criteria.FederationID})
	}
	if criteria.StateID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_STATE_ID: criteria.StateID})
	}

	if criteria.CountryID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_COUNTRY_ID: criteria.CountryID})
	}
	if criteria.StartDateTime.After(time.Now()) {
		stmt = stmt.Where(squirrel.Eq{DAS_COMPETITION_COL_DATETIME_START: criteria.StartDateTime})
	}
	if criteria.OrganizerID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnCreateUserID: criteria.OrganizerID})
	}
	if criteria.StatusID > 0 {
		stmt = stmt.Where(squirrel.Eq{DAS_COMPETITION_COL_STATUS_ID: criteria.StatusID})
	}

	rows, err := stmt.RunWith(repo.Database).Query()
	comps := make([]businesslogic.Competition, 0)

	for rows.Next() {
		each := businesslogic.Competition{}
		status := 0
		rows.Scan(
			&each.ID,
			&each.FederationID,
			&each.Name,
			&each.Street,
			&each.City.ID,
			&each.State.ID,
			&each.Country.ID,
			&each.StartDateTime,
			&each.EndDateTime,
			&each.ContactName,
			&each.ContactPhone,
			&each.ContactEmail,
			&each.Website,
			&status,
			&each.Attendance,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		each.UpdateStatus(status)
		comps = append(comps, each)
	}
	return comps, err
}
