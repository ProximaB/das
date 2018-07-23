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
	"errors"
	"fmt"
	"time"

	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/DancesportSoftware/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
	"log"
)

const (
	DasUserAccountTable                     = "DAS.ACCOUNT"
	DAS_USER_ACCOUNT_COL_USER_STATUS_ID     = "ACCOUNT_STATUS_ID"
	DAS_USER_ACCOUNT_COL_USER_GENDER_ID     = "USER_GENDER_ID"
	DAS_USER_ACCOUNT_COL_LAST_NAME          = "LAST_NAME"
	DAS_USER_ACCOUNT_COL_MIDDLE_NAMES       = "MIDDLE_NAMES"
	DAS_USER_ACCOUNT_COL_FIRST_NAME         = "FIRST_NAME"
	DAS_USER_ACCOUNT_COL_DATE_OF_BIRTH      = "DATE_OF_BIRTH"
	DAS_USER_ACCOUNT_COL_EMAIL              = "EMAIL"
	DAS_USER_ACCOUNT_COL_PHONE              = "PHONE"
	DAS_USER_ACCOUNT_COL_EMAIL_VERIFIED     = "EMAIL_VERIFIED"
	DAS_USER_ACCOUNT_COL_PHONE_VERIFIED     = "PHONE_VERIFIED"
	DAS_USER_ACCOUNT_HASH_ALGORITHM         = "HASH_ALGORITHM"
	DAS_USER_ACCOUNT_COL_PASSWORD_SALT      = "PASSWORD_SALT"
	DAS_USER_ACCOUNT_COL_PASSWORD_HASH      = "PASSWORD_HASH"
	DAS_USER_ACCOUNT_COL_DATETIME_CREATED   = "DATETIME_CREATED"
	DAS_USER_ACCOUNT_COL_DATETIME_UPDATED   = "DATETIME_UPDATED"
	DAS_USER_ACCOUNT_COL_TOS_ACCEPTED       = "TOS_ACCEPTED"
	DAS_USER_ACCOUNT_COL_PP_ACCEPTED        = "PP_ACCEPTED"
	DAS_USER_ACCOUNT_COL_BY_GUARDIAN        = "BY_GUARDIAN"
	DAS_USER_ACCOUNT_COL_GUARDIAN_SIGNATURE = "GUARDIAN_SIGNATURE"
)

type PostgresAccountRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

func (repo PostgresAccountRepository) CreateAccount(account *businesslogic.Account) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.
		Insert("").
		Into(DasUserAccountTable).
		Columns(
			common.ColumnUUID,
			DAS_USER_ACCOUNT_COL_USER_STATUS_ID,
			DAS_USER_ACCOUNT_COL_USER_GENDER_ID,
			DAS_USER_ACCOUNT_COL_LAST_NAME,
			DAS_USER_ACCOUNT_COL_MIDDLE_NAMES,
			DAS_USER_ACCOUNT_COL_FIRST_NAME,
			DAS_USER_ACCOUNT_COL_DATE_OF_BIRTH,
			DAS_USER_ACCOUNT_COL_EMAIL,
			DAS_USER_ACCOUNT_COL_PHONE,
			DAS_USER_ACCOUNT_COL_EMAIL_VERIFIED,
			DAS_USER_ACCOUNT_COL_PHONE_VERIFIED,
			DAS_USER_ACCOUNT_HASH_ALGORITHM,
			DAS_USER_ACCOUNT_COL_PASSWORD_SALT,
			DAS_USER_ACCOUNT_COL_PASSWORD_HASH,
			common.ColumnDateTimeCreated,
			common.ColumnDateTimeUpdated,
			DAS_USER_ACCOUNT_COL_TOS_ACCEPTED,
			DAS_USER_ACCOUNT_COL_PP_ACCEPTED,
			DAS_USER_ACCOUNT_COL_BY_GUARDIAN,
			DAS_USER_ACCOUNT_COL_GUARDIAN_SIGNATURE).
		Values(
			account.UUID,
			account.AccountStatusID,
			account.UserGenderID,
			account.LastName,
			account.MiddleNames,
			account.FirstName,
			account.DateOfBirth,
			account.Email,
			account.Phone,
			account.EmailVerified,
			account.PhoneVerified,
			account.HashAlgorithm,
			account.PasswordSalt,
			account.PasswordHash,
			time.Now(),
			time.Now(),
			account.ToSAccepted,
			account.PrivacyPolicyAccepted,
			account.ByGuardian,
			account.Signature,
		).Suffix(dalutil.SQLSuffixReturningID)

	// parsing arguments to ... parameters: https://golang.org/ref/spec#Passing_arguments_to_..._parameters
	// PostgreSQL does not return LastInsertID automatically: https://github.com/lib/pq/issues/24
	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&account.ID)
		err = tx.Commit()
	}
	return err
}

func (repo PostgresAccountRepository) SearchAccount(criteria businesslogic.SearchAccountCriteria) ([]businesslogic.Account, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.
		Select(
			fmt.Sprintf(
				"%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
				common.ColumnPrimaryKey,
				common.ColumnUUID,
				DAS_USER_ACCOUNT_COL_USER_STATUS_ID,
				DAS_USER_ACCOUNT_COL_USER_GENDER_ID,
				DAS_USER_ACCOUNT_COL_LAST_NAME,
				DAS_USER_ACCOUNT_COL_MIDDLE_NAMES,
				DAS_USER_ACCOUNT_COL_FIRST_NAME,
				DAS_USER_ACCOUNT_COL_DATE_OF_BIRTH,
				DAS_USER_ACCOUNT_COL_EMAIL,
				DAS_USER_ACCOUNT_COL_PHONE,
				DAS_USER_ACCOUNT_COL_EMAIL_VERIFIED,
				DAS_USER_ACCOUNT_COL_PHONE_VERIFIED,
				DAS_USER_ACCOUNT_HASH_ALGORITHM,
				DAS_USER_ACCOUNT_COL_PASSWORD_SALT,
				DAS_USER_ACCOUNT_COL_PASSWORD_HASH,
				DAS_USER_ACCOUNT_COL_DATETIME_CREATED,
				DAS_USER_ACCOUNT_COL_DATETIME_UPDATED,
				DAS_USER_ACCOUNT_COL_TOS_ACCEPTED,
				DAS_USER_ACCOUNT_COL_PP_ACCEPTED,
				DAS_USER_ACCOUNT_COL_BY_GUARDIAN,
				DAS_USER_ACCOUNT_COL_GUARDIAN_SIGNATURE,
			)).From(DasUserAccountTable)

	if len(criteria.UUID) != 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnUUID: criteria.UUID})
	}
	if criteria.ID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.ID})
	}
	if criteria.AccountStatus > 0 {
		stmt = stmt.Where(squirrel.Eq{DAS_USER_ACCOUNT_COL_USER_STATUS_ID: criteria.AccountStatus})
	}
	if criteria.Gender > 0 {
		stmt = stmt.Where(squirrel.Eq{DAS_USER_ACCOUNT_COL_USER_GENDER_ID: criteria.Gender})
	}
	if len(criteria.Email) > 0 {
		stmt = stmt.Where(squirrel.Eq{DAS_USER_ACCOUNT_COL_EMAIL: criteria.Email})
	}
	if len(criteria.Phone) > 0 {
		stmt = stmt.Where(squirrel.Eq{DAS_USER_ACCOUNT_COL_PHONE: criteria.Phone})
	}
	if len(criteria.LastName) > 0 {
		stmt = stmt.Where(squirrel.Eq{DAS_USER_ACCOUNT_COL_LAST_NAME: criteria.LastName})
	}
	if len(criteria.FirstName) > 0 {
		stmt = stmt.Where(squirrel.Eq{DAS_USER_ACCOUNT_COL_FIRST_NAME: criteria.FirstName})
	}

	accounts := make([]businesslogic.Account, 0)
	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return accounts, err
	}
	for rows.Next() {
		each := businesslogic.Account{}
		rows.Scan(
			&each.ID,
			&each.UUID,
			&each.AccountStatusID,
			&each.UserGenderID,
			&each.LastName,
			&each.MiddleNames,
			&each.FirstName,
			&each.DateOfBirth,
			&each.Email,
			&each.Phone,
			&each.EmailVerified,
			&each.PhoneVerified,
			&each.HashAlgorithm,
			&each.PasswordSalt,
			&each.PasswordHash,
			&each.DateTimeCreated,
			&each.DateTimeModified,
			&each.ToSAccepted,
			&each.PrivacyPolicyAccepted,
			&each.ByGuardian,
			&each.Signature,
		)
		accounts = append(accounts, each)
	}
	rows.Close()

	// now query roles for each account
	for _, each := range accounts {
		queryRoleStmt := repo.SQLBuilder.Select(fmt.Sprintf(
			"%s, %s, %s, %s, %s, %s, %s",
			common.ColumnPrimaryKey,
			common.ColumnAccountID,
			common.ColumnAccountTypeID,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated)).From(dasAccountRoleTable).
			Where(squirrel.Eq{common.ColumnAccountID: each.ID})
		roleRows, roleErr := queryRoleStmt.RunWith(repo.Database).Query()
		if roleErr != nil {
			return nil, roleErr
		}
		accountRoles := make([]businesslogic.AccountRole, 0)
		for roleRows.Next() {
			eachRole := businesslogic.AccountRole{}
			roleRows.Scan(
				&eachRole.ID,
				&eachRole.AccountID,
				&eachRole.AccountTypeID,
				&eachRole.CreateUserID,
				&eachRole.DateTimeCreated,
				&eachRole.UpdateUserID,
				&eachRole.DateTimeUpdated,
			)
			accountRoles = append(accountRoles, eachRole)
		}
		// TODO: High Priority[BUG] Roles are not set to accounts for some reason.
		each.SetRoles(accountRoles)
		log.Println(each.GetRoles())
		roleRows.Close()
	}

	for _, each := range accounts {
		log.Println(each.GetRoles())
	}
	return accounts, err
}

func (repo PostgresAccountRepository) DeleteAccount(account businesslogic.Account) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	if account.ID > 0 {
		stmt := repo.SQLBuilder.Delete("").From(DasUserAccountTable).Where(squirrel.Eq{common.ColumnPrimaryKey: account.ID})
		_, err := stmt.RunWith(repo.Database).Exec()
		return err
	}
	return errors.New("account ID was not specified")
}

func (repo PostgresAccountRepository) UpdateAccount(account businesslogic.Account) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.Update(DasUserAccountTable)
	if account.ID > 0 {
		if len(account.PasswordSalt) > 0 {
			stmt = stmt.Set(DAS_USER_ACCOUNT_COL_PASSWORD_SALT, account.PasswordSalt)
		}
	}
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		err = tx.Commit()
	}
	return err
}
