package partnershipdal

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
)

type PostgresPartnershipStatusRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}
