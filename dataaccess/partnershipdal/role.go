package partnershipdal

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
)

type PostgresPartnershipRoleRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}
