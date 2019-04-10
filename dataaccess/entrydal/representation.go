package entrydal

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
)

// PostgresPartnershipCompetitionRepresentationRepository implements IPartnershipCompetitionRepresentationRepository with a Postgres database
type PostgresPartnershipCompetitionRepresentationRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

// CreateCompetitionRepresentation creates a PartnershipCompetitionRepresentation in a Postgres database
func (repo PostgresPartnershipCompetitionRepresentationRepository) CreateCompetitionRepresentation() {

}

// SearchCompetitionRepresentation searches PartnershipCompetitionRepresentation in a Postgres database
func (repo PostgresPartnershipCompetitionRepresentationRepository) SearchCompetitionRepresentation() {

}
