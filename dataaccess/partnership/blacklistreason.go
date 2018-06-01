package partnership

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/dataaccess/common"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
)

type PostgresPartnershipRequestBlacklistReasonRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresPartnershipRequestBlacklistReasonRepository) GetPartnershipRequestBlacklistReasons() ([]businesslogic.PartnershipRequestBlacklistReason, error) {
	stmt := repo.SqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s",
		common.PRIMARY_KEY,
		common.COL_NAME,
		common.COL_DESCRIPTION,
		common.COL_DATETIME_CREATED,
		common.COL_DATETIME_UPDATED,
	)).From(DAS_PARTNERSHIP_REQUEST_BLACKLIST_REASON_TABLE).
		OrderBy(common.PRIMARY_KEY)
	rows, err := stmt.RunWith(repo.Database).Query()
	output := make([]businesslogic.PartnershipRequestBlacklistReason, 0)
	if err != nil {
		return output, err
	}

	for rows.Next() {
		each := businesslogic.PartnershipRequestBlacklistReason{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Description,
			&each.DateTimeCreated,
			&each.DateTimeUpdated,
		)
		output = append(output, each)
	}
	rows.Close()
	return output, err
}
