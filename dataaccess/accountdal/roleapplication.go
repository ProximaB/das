package accountdal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/dataaccess/common"
	"github.com/ProximaB/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
	"log"
)

const (
	dasAccountRoleApplicationTable                  = "DAS.ACCOUNT_ROLE_APPLICATION"
	dasAccountRoleApplicationColumnAppliedRoleID    = "APPLIED_ROLE_ID"
	dasAccountRoleApplicationColumnApprovalUserID   = "APPROVAL_USER_ID"
	dasAccountRoleApplicationColumnDateTimeApproved = "DATETIME_APPROVED"
)

// PostgresRoleApplicationRepository implements IRoleApplicationRepository
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
		Into(dasAccountRoleApplicationTable).
		Columns(
			common.ColumnAccountID,
			dasAccountRoleApplicationColumnAppliedRoleID,
			common.COL_DESCRIPTION,
			common.ColumnStatusID,
			dasAccountRoleApplicationColumnApprovalUserID,
			dasAccountRoleApplicationColumnDateTimeApproved,
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
	// create an account repository here as it will be needed to query accounts later.
	accountRepo := PostgresAccountRepository{repo.Database, repo.SQLBuilder}
	stmt := repo.SQLBuilder.
		Select(
			fmt.Sprintf(
				"%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
				common.ColumnPrimaryKey,
				common.ColumnAccountID,
				dasAccountRoleApplicationColumnAppliedRoleID,
				common.COL_DESCRIPTION,
				common.ColumnStatusID,
				dasAccountRoleApplicationColumnApprovalUserID,
				dasAccountRoleApplicationColumnDateTimeApproved,
				common.ColumnCreateUserID,
				common.ColumnDateTimeCreated,
				common.ColumnUpdateUserID,
				common.ColumnDateTimeUpdated)).
		From(dasAccountRoleApplicationTable)
	if criteria.ID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.ID})
	}
	if criteria.AccountID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnAccountID: criteria.AccountID})
	}
	if criteria.AppliedRoleID > 0 {
		stmt = stmt.Where(squirrel.Eq{dasAccountRoleApplicationColumnAppliedRoleID: criteria.AppliedRoleID})
	}
	if criteria.StatusID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnStatusID: criteria.StatusID})
	}
	if criteria.ApprovalUserID > 9 {
		stmt = stmt.Where(squirrel.Eq{dasAccountRoleApplicationColumnApprovalUserID: criteria.ApprovalUserID})
	}
	if criteria.Responded && criteria.StatusID == 0 {
		stmt = stmt.Where(squirrel.NotEq{common.ColumnStatusID: businesslogic.RoleApplicationStatusPending})
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
		accountSearchResults, searchErr := accountRepo.SearchAccount(businesslogic.SearchAccountCriteria{ID: each.AccountID})
		if searchErr != nil {
			log.Printf("[error] searching account in role application throw and error: %v", searchErr)
			return nil, errors.New("error in searching for accounts")
		}
		if len(accountSearchResults) < 1 {
			log.Printf("[error] cannot find account when it should be in database: AccountID = %v", criteria.AccountID)
			return nil, errors.New("cannot find the specified account")
		}
		each.Account = accountSearchResults[0]
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
	stmt := repo.SQLBuilder.Update(dasAccountRoleApplicationTable)
	if application.ID > 0 {
		if application.StatusID > 0 && *application.ApprovalUserID > 0 {
			stmt = stmt.Set(common.ColumnStatusID, application.StatusID)
			stmt = stmt.Set(dasAccountRoleApplicationColumnApprovalUserID, *application.ApprovalUserID)
			stmt = stmt.Set(dasAccountRoleApplicationColumnDateTimeApproved, application.DateTimeApproved)
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
