// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
