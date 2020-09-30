package accountdal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/dataaccess/common"
	dalutil "github.com/ProximaB/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
)

const (
	DAS_ACCOUNT_ROLE_APPLICATION_STATUS = "DAS.ACCOUNT_ROLE_APPLICATION_STATUS"
)

type PostgresRoleApplicationStatusRepository struct {
	Database  *sql.DB
	SqlBulder squirrel.StatementBuilderType
}

func (repo PostgresRoleApplicationStatusRepository) GetAllRoleApplicationStatus() ([]businesslogic.RoleApplicationStatus, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}

	applicationStatuses := make([]businesslogic.RoleApplicationStatus, 0)
	stmt := repo.SqlBulder.
		Select(
			fmt.Sprintf(
				"%s, %s, %s, %s",
				common.ColumnPrimaryKey,
				common.COL_NAME,
				common.ColumnDateTimeCreated,
				common.ColumnDateTimeUpdated)).
		From(DAS_ACCOUNT_ROLE_APPLICATION_STATUS).
		OrderBy(common.ColumnPrimaryKey)
	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return applicationStatuses, err
	}

	for rows.Next() {
		each := businesslogic.RoleApplicationStatus{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.DateTimeCreated,
			&each.DateTimeUpdated,
		)
		applicationStatuses = append(applicationStatuses, each)
	}
	rows.Close()
	return applicationStatuses, err
}
