package partnership

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/dataaccess/common"
	"database/sql"
	"fmt"
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
	clause := repo.SqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s",
		common.PRIMARY_KEY,
		DAS_PARTNERSHIP_REQUEST_STATUS_COL_CODE,
		common.COL_DESCRIPTION,
		common.COL_DATETIME_CREATED,
		common.COL_DATETIME_UPDATED)).From(DAS_PARTNERSHIP_REQUEST_STATUS_TABLE).OrderBy(common.PRIMARY_KEY)
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
