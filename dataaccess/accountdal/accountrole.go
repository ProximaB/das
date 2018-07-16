package accountdal

import (
	"database/sql"
	"errors"
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
	stmt := repo.SQLBuilder.Insert("").Into(dasAccountRoleTable).Columns(
		common.COL_ACCOUNT_ID,
		common.ColumnAccountTypeID,
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

func (repo PostgresAccountRoleRepository) DeleteAccountRole(role businesslogic.AccountRole) error {
	return errors.New("not implemented")
}

func (repo PostgresAccountRoleRepository) SearchAccountRole(criteria businesslogic.SearchAccountRoleCriteria) ([]businesslogic.AccountRole, error) {
	return nil, errors.New("not implemented")
}

func (repo PostgresAccountRoleRepository) UpdateAccountRole(role businesslogic.AccountRole) error {
	return errors.New("not implemented")
}
