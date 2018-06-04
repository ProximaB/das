package account

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/Masterminds/squirrel"
)

const DAS_ACCOUNT_STATUS_TABLE = "DAS.ACCOUNT_STATUS"

type PostgresAccountStatusRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresAccountStatusRepository) GetAccountStatus() ([]businesslogic.AccountStatus, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresAccountStatusRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s",
			common.PRIMARY_KEY,
			common.COL_NAME,
			common.COL_ABBREVIATION,
			common.COL_DESCRIPTION,
			common.COL_DATETIME_CREATED,
			common.COL_DATETIME_UPDATED)).
		From(DAS_ACCOUNT_STATUS_TABLE).
		OrderBy(common.PRIMARY_KEY)

	status := make([]businesslogic.AccountStatus, 0)
	rows, err := stmt.RunWith(repo.Database).Query()

	if err != nil {
		return status, err
	}

	for rows.Next() {
		each := businesslogic.AccountStatus{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Abbreviation,
			&each.Description,
			&each.DateTimeCreated,
			&each.DateTimeUpdated,
		)
		status = append(status, each)
	}
	rows.Close()

	return status, err
}
