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

package partnership

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/Masterminds/squirrel"
)

type PostgresPartnershipRequestBlacklistReasonRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresPartnershipRequestBlacklistReasonRepository) GetPartnershipRequestBlacklistReasons() ([]businesslogic.PartnershipRequestBlacklistReason, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresPartnershipRequestBlacklistReasonRepository is not specified")
	}
	stmt := repo.SqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s",
		common.PRIMARY_KEY,
		common.COL_NAME,
		common.COL_DESCRIPTION,
		common.COL_DATETIME_CREATED,
		common.COL_DATETIME_UPDATED,
	)).From(DAS_PARTNERSHIP_REQUEST_BLACKLIST_REASON_TABLE).
		OrderBy(common.PRIMARY_KEY)
	rows, err := stmt.RunWith(repo.Database).Query()
	output := make([]businesslogic.PartnershipRequestBlacklistReason, 0)
	if err != nil {
		return output, err
	}

	for rows.Next() {
		each := businesslogic.PartnershipRequestBlacklistReason{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Description,
			&each.DateTimeCreated,
			&each.DateTimeUpdated,
		)
		output = append(output, each)
	}
	rows.Close()
	return output, err
}
