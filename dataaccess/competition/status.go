// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package competition

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/Masterminds/squirrel"
)

const (
	DAS_COMPETITION_STATUS_TABLE = "DAS.COMPETITION_STATUS"
)

type PostgresCompetitionStatusRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresCompetitionStatusRepository) GetCompetitionAllStatus() ([]businesslogic.CompetitionStatus, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresCompetitionStatusRepository is not specified")
	}
	clause := repo.SqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s",
		common.PRIMARY_KEY,
		common.COL_NAME,
		common.COL_ABBREVIATION,
		common.COL_DESCRIPTION,
		common.COL_DATETIME_CREATED,
		common.COL_DATETIME_UPDATED)).
		From(DAS_COMPETITION_STATUS_TABLE).OrderBy(common.PRIMARY_KEY)
	rows, err := clause.RunWith(repo.Database).Query()
	status := make([]businesslogic.CompetitionStatus, 0)
	if err != nil {
		return status, err
	}
	for rows.Next() {
		each := businesslogic.CompetitionStatus{}
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
