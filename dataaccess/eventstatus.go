package dataaccess

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/dataaccess/common"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

const DAS_EVENT_STATUS_TABLE = "DAS.EVENT_STATUS"

type PostgresEventStatusRepository struct {
	database   *sql.DB
	sqlBuilder sq.StatementBuilderType
}

func (repo PostgresEventStatusRepository) GetEventStatus() ([]businesslogic.EventStatus, error) {
	stmt := repo.sqlBuilder.Select(
		fmt.Sprintf("%s, %s, %s, %s, %s, %s",
			common.PRIMARY_KEY, common.COL_NAME,
			common.COL_ABBREVIATION, common.COL_DESCRIPTION,
			common.COL_DATETIME_CREATED, common.COL_DATETIME_UPDATED)).From(DAS_EVENT_STATUS_TABLE)

	status := make([]businesslogic.EventStatus, 0)
	rows, err := stmt.RunWith(repo.database).Query()

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
