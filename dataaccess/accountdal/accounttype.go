package accountdal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/dataaccess/common"
	"github.com/ProximaB/das/dataaccess/util"
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
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	accountTypes := make([]businesslogic.AccountType, 0)
	stmt := repo.SqlBuilder.
		Select(
			fmt.Sprintf(
				"%s, %s, %s, %s, %s",
				common.ColumnPrimaryKey,
				common.COL_NAME,
				common.COL_DESCRIPTION,
				common.ColumnDateTimeCreated,
				common.ColumnDateTimeUpdated)).
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
