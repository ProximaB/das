// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package event

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
