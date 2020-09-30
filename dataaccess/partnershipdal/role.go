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

const dasPartnershipRoleTable = "DAS.PARTNERSHIP_ROLE"

// PostgresPartnershipRoleRepository implements the IPartnershipRoleRepository with a Postgres database
type PostgresPartnershipRoleRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresPartnershipRoleRepository) GetAllPartnershipRoles() ([]businesslogic.PartnershipRole, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s", common.ColumnPrimaryKey, common.COL_NAME, common.ColumnDateTimeCreated, common.ColumnDateTimeUpdated)).From(dasPartnershipRoleTable)
	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return nil, err
	}
	roles := make([]businesslogic.PartnershipRole, 0)
	for rows.Next() {
		each := businesslogic.PartnershipRole{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.DateTimeCreated,
			&each.DateTimeUpdated,
		)
		roles = append(roles, each)
	}
	rows.Close()
	return roles, err

}
