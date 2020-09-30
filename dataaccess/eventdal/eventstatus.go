package eventdal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/dataaccess/common"
	"github.com/ProximaB/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
)

const DAS_EVENT_STATUS_TABLE = "DAS.EVENT_STATUS"

type PostgresEventStatusRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresEventStatusRepository) GetEventStatus() ([]businesslogic.EventStatus, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Select(
		fmt.Sprintf("%s, %s, %s, %s, %s, %s",
			common.ColumnPrimaryKey, common.COL_NAME,
			common.ColumnAbbreviation, common.COL_DESCRIPTION,
			common.ColumnDateTimeCreated, common.ColumnDateTimeUpdated)).From(DAS_EVENT_STATUS_TABLE)

	status := make([]businesslogic.EventStatus, 0)
	rows, err := stmt.RunWith(repo.Database).Query()

	if err != nil {
		return status, err
	}

	for rows.Next() {
		each := businesslogic.EventStatus{}
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
