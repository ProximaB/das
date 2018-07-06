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

package eventdal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/Masterminds/squirrel"
)

const DAS_EVENT_STATUS_TABLE = "DAS.EVENT_STATUS"

type PostgresEventStatusRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresEventStatusRepository) GetEventStatus() ([]businesslogic.EventStatus, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresEventStatusRepository is not specified")
	}
	stmt := repo.SqlBuilder.Select(
		fmt.Sprintf("%s, %s, %s, %s, %s, %s",
			common.PRIMARY_KEY, common.COL_NAME,
			common.COL_ABBREVIATION, common.COL_DESCRIPTION,
			common.COL_DATETIME_CREATED, common.COL_DATETIME_UPDATED)).From(DAS_EVENT_STATUS_TABLE)

	status := make([]businesslogic.EventStatus, 0)
	rows, err := stmt.RunWith(repo.Database).Query()

	if err != nil {
		return status, err
	}

	for rows.Next() {
		each := businesslogic.EventStatus{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Abbreviation,
			&each.Description,
			&each.DateTimeCreated,
			&each.DateTimeUpdated,
		)
		status = append(status, each)
	}
	rows.Close()

	return status, err
}
