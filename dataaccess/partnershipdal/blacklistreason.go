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

type PostgresPartnershipRequestBlacklistReasonRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresPartnershipRequestBlacklistReasonRepository) GetPartnershipRequestBlacklistReasons() ([]businesslogic.PartnershipRequestBlacklistReason, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s",
		common.ColumnPrimaryKey,
		common.COL_NAME,
		common.COL_DESCRIPTION,
		common.ColumnDateTimeCreated,
		common.ColumnDateTimeUpdated,
	)).From(DAS_PARTNERSHIP_REQUEST_BLACKLIST_REASON_TABLE).
		OrderBy(common.ColumnPrimaryKey)
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
