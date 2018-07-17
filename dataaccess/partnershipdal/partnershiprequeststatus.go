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

package partnershipdal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/Masterminds/squirrel"
)

const (
	DAS_PARTNERSHIP_REQUEST_STATUS_TABLE                 = "DAS.PARTNERSHIP_REQUEST_STATUS"
	DAS_PARTNERSHIP_REQUEST_STATUS_COL_REQUEST_STATUS_ID = "REQUEST_STATUS_ID"
	DAS_PARTNERSHIP_REQUEST_STATUS_COL_CODE              = "CODE"
)

type PostgresPartnershipRequestStatusRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresPartnershipRequestStatusRepository) GetPartnershipRequestStatus() ([]businesslogic.PartnershipRequestStatus, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresPartnershipRequestStatusRepository is not specified")
	}
	clause := repo.SqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s",
		common.ColumnPrimaryKey,
		DAS_PARTNERSHIP_REQUEST_STATUS_COL_CODE,
		common.COL_DESCRIPTION,
		common.ColumnDateTimeCreated,
		common.ColumnDateTimeUpdated)).From(DAS_PARTNERSHIP_REQUEST_STATUS_TABLE).OrderBy(common.ColumnPrimaryKey)
	rows, err := clause.RunWith(repo.Database).Query()
	output := make([]businesslogic.PartnershipRequestStatus, 0)
	if err != nil {
		return output, err
	}
	for rows.Next() {
		each := businesslogic.PartnershipRequestStatus{}
		rows.Scan(
			&each.ID,
			&each.Code,
			&each.Description,
			&each.DateTimeCreated,
			&each.DateTimeUpdated,
		)
		output = append(output, each)
	}
	rows.Close()
	return output, err

}
