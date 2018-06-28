// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package account

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
				common.PRIMARY_KEY,
				common.COL_NAME,
				common.COL_DESCRIPTION,
				common.COL_DATETIME_CREATED,
				common.COL_DATETIME_UPDATED)).
		From(DAS_ACCOUNT_TYPE_TABLE).
		OrderBy(common.PRIMARY_KEY)
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
