// Dancesport Application System (DAS)
// Copyright (C) 2018 Yubing Hou
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
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/DancesportSoftware/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
)

const dasAccountRoleTable = "DAS.ACCOUNT_ROLE"

type PostgresAccountRoleRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

func (repo PostgresAccountRoleRepository) CreateAccountRole(role *businesslogic.AccountRole) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.Insert("").
		Into(dasAccountRoleTable).
		Columns(
			common.ColumnAccountID,
			common.ColumnAccountTypeID,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated,
		).Values(
		role.AccountID,
		role.AccountTypeID,
		role.CreateUserID,
		role.DateTimeCreated,
		role.UpdateUserID,
		role.DateTimeUpdated,
	).Suffix(dalutil.SQLSuffixReturningID)

	clause, args, err := stmt.ToSql()
	tx, txErr := repo.Database.Begin()
	if txErr != nil {
		return txErr
	}
	row := repo.Database.QueryRow(clause, args...)
	row.Scan(&role.ID)
	tx.Commit()
	return err
}

func (repo PostgresAccountRoleRepository) SearchAccountRole(criteria businesslogic.SearchAccountRoleCriteria) ([]businesslogic.AccountRole, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s",
		common.ColumnPrimaryKey,
		common.ColumnAccountID,
		common.ColumnAccountTypeID,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated,
	)).From(dasAccountRoleTable)

	if criteria.AccountID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnAccountID: criteria.AccountID})
	}
	if criteria.AccountTypeID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnAccountTypeID: criteria.AccountTypeID})
	}
	roles := make([]businesslogic.AccountRole, 0)
	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return roles, err
	}
	for rows.Next() {
		each := businesslogic.AccountRole{}
		rows.Scan(
			&each.ID,
			&each.AccountID,
			&each.AccountTypeID,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		roles = append(roles, each)
	}
	rows.Close()
	return roles, err
}
