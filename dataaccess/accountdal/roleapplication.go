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

const (
	DasAccountRoleApplicationTable                  = "DAS.ACCOUNT_ROLE_APPLICATION"
	DasAccountRoleApplicationColumnAppliedRoleID    = "APPLIED_ROLE_ID"
	DasAccountRoleApplicationColumnApprovalUserID   = "APPROVAL_USER_ID"
	DasAccountRoleApplicationColumnDateTimeApproved = "DATETIME_APPROVED"
)

type PostgresRoleApplicationRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

// CreateApplication creates a RoleApplication in a Postgres database
func (repo PostgresRoleApplicationRepository) CreateApplication(application *businesslogic.RoleApplication) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.
		Insert("").
		Into(DasAccountRoleApplicationTable).
		Columns(
			common.ColumnAccountID,
			DasAccountRoleApplicationColumnAppliedRoleID,
			common.COL_DESCRIPTION,
			common.ColumnStatusID,
			DasAccountRoleApplicationColumnApprovalUserID,
			DasAccountRoleApplicationColumnDateTimeApproved,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated).
		Values(
			application.AccountID,
			application.AppliedRoleID,
			application.Description,
			application.StatusID,
			application.ApprovalUserID,
			application.DateTimeApproved,
			application.CreateUserID,
			application.DateTimeCreated,
			application.UpdateUserID,
			application.DateTimeUpdated).
		Suffix(dalutil.SQLSuffixReturningID)

	clause, args, err := stmt.ToSql()
	tx, txErr := repo.Database.Begin()
	if txErr != nil {
		return txErr
	}

	row := repo.Database.QueryRow(clause, args...)
	row.Scan(&application.ID)
	tx.Commit()
	return err
}

// SearchApplication searches RoleApplications from a Postgres database with the given criteria. The returned applications
// can be nil or a concrete slice. The returned error can be either nil or concrete error
func (repo PostgresRoleApplicationRepository) SearchApplication(criteria businesslogic.SearchRoleApplicationCriteria) ([]businesslogic.RoleApplication, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.
		Select(
			fmt.Sprintf(
				"%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
				common.ColumnPrimaryKey,
				common.ColumnAccountID,
				DasAccountRoleApplicationColumnAppliedRoleID,
				common.COL_DESCRIPTION,
				common.ColumnStatusID,
				DasAccountRoleApplicationColumnApprovalUserID,
				DasAccountRoleApplicationColumnDateTimeApproved,
				common.ColumnCreateUserID,
				common.ColumnDateTimeCreated,
				common.ColumnUpdateUserID,
				common.ColumnDateTimeUpdated)).
		From(DasAccountRoleApplicationTable)
	if criteria.AccountID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnAccountID: criteria.AccountID})
	}
	if criteria.AppliedRoleID > 0 {
		stmt = stmt.Where(squirrel.Eq{DasAccountRoleApplicationColumnAppliedRoleID: criteria.AppliedRoleID})
	}
	if criteria.StatusID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnStatusID: criteria.StatusID})
	}
	if criteria.ApprovalUserID > 9 {
		stmt = stmt.Where(squirrel.Eq{DasAccountRoleApplicationColumnApprovalUserID: criteria.ApprovalUserID})
	}

	applications := make([]businesslogic.RoleApplication, 0)
	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		each := businesslogic.RoleApplication{}
		rows.Scan(
			&each.ID,
			&each.AccountID,
			&each.AppliedRoleID,
			&each.Description,
			&each.StatusID,
			&each.ApprovalUserID,
			&each.DateTimeApproved,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		applications = append(applications, each)
	}
	rows.Close()
	return applications, err
}

// UpdateApplication update the RoleApplication of the same ID to the given application. This mostly update the user
// who made the change as well as time stamps. The returned error could be either nil or a concrete error
func (repo PostgresRoleApplicationRepository) UpdateApplication(application businesslogic.RoleApplication) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.Update(DasAccountRoleApplicationTable)
	if application.ID > 0 {
		if application.StatusID > 0 && *application.ApprovalUserID > 0 {
			stmt = stmt.Set(common.ColumnStatusID, application.StatusID)
			stmt = stmt.Set(DasAccountRoleApplicationColumnApprovalUserID, *application.ApprovalUserID)
			stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: application.ID})
		}
		tx, txErr := repo.Database.Begin()
		if txErr != nil {
			return txErr
		}
		_, err := stmt.RunWith(repo.Database).Exec()
		tx.Commit()
		return err
	}
	return errors.New("the ID of role application is not specified")
}
