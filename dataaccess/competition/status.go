package competition

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
	DAS_COMPETITION_STATUS_TABLE = "DAS.COMPETITION_STATUS"
)

type PostgresCompetitionStatusRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresCompetitionStatusRepository) GetCompetitionAllStatus() ([]businesslogic.CompetitionStatus, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	clause := repo.SqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s",
		common.ColumnPrimaryKey,
		common.COL_NAME,
		common.ColumnAbbreviation,
		common.COL_DESCRIPTION,
		common.ColumnDateTimeCreated,
		common.ColumnDateTimeUpdated)).
		From(DAS_COMPETITION_STATUS_TABLE).OrderBy(common.ColumnPrimaryKey)
	rows, err := clause.RunWith(repo.Database).Query()
	status := make([]businesslogic.CompetitionStatus, 0)
	if err != nil {
		return status, err
	}
	for rows.Next() {
		each := businesslogic.CompetitionStatus{}
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
