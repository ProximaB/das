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

package accountdal

import (
	"database/sql"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/Masterminds/squirrel"
)

const (
	DAS_ACCOUNT_TYPE_TABLE = "DAS.ACCOUNT_TYPE"
)

type PostgresAccountTypeRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresAccountTypeRepository) GetAccountTypes() ([]businesslogic.AccountType, error) {
	accountTypes := make([]businesslogic.AccountType, 0)
	stmt := repo.SqlBuilder.
		Select(
			fmt.Sprintf(
				"%s, %s, %s, %s, %s",
				common.ColumnPrimaryKey,
				common.COL_NAME,
				common.COL_DESCRIPTION,
				common.COL_DATETIME_CREATED,
				common.COL_DATETIME_UPDATED)).
		From(DAS_ACCOUNT_TYPE_TABLE).
		OrderBy(common.ColumnPrimaryKey)
	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return accountTypes, err
	}

	for rows.Next() {
		each := businesslogic.AccountType{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Description,
			&each.DateTimeCreated,
			&each.DateTimeUpdated,
		)
		accountTypes = append(accountTypes, each)
	}
	rows.Close()
	return accountTypes, err
}
