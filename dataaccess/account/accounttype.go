package account

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/dataaccess/common"
	"database/sql"
	"fmt"
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
