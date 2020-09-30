package partnershipdal

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
	DAS_PARTNERSHIP_REQUEST_STATUS_TABLE                 = "DAS.PARTNERSHIP_REQUEST_STATUS"
	DAS_PARTNERSHIP_REQUEST_STATUS_COL_REQUEST_STATUS_ID = "REQUEST_STATUS_ID"
	DAS_PARTNERSHIP_REQUEST_STATUS_COL_CODE              = "CODE"
)

type PostgresPartnershipRequestStatusRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresPartnershipRequestStatusRepository) GetPartnershipRequestStatus() ([]businesslogic.PartnershipRequestStatus, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	clause := repo.SqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s",
		common.ColumnPrimaryKey,
		DAS_PARTNERSHIP_REQUEST_STATUS_COL_CODE,
		common.COL_DESCRIPTION,
		common.ColumnDateTimeCreated,
		common.ColumnDateTimeUpdated)).From(DAS_PARTNERSHIP_REQUEST_STATUS_TABLE).OrderBy(common.ColumnPrimaryKey)
	rows, err := clause.RunWith(repo.Database).Query()
	output := make([]businesslogic.PartnershipRequestStatus, 0)
	if err != nil {
		return output, err
	}
	for rows.Next() {
		each := businesslogic.PartnershipRequestStatus{}
		rows.Scan(
			&each.ID,
			&each.Code,
			&each.Description,
			&each.DateTimeCreated,
			&each.DateTimeUpdated,
		)
		output = append(output, each)
	}
	rows.Close()
	return output, err

}
