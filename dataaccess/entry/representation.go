// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package entry

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
)

type PostgresPartnershipCompetitionRepresentationRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresPartnershipCompetitionRepresentationRepository) CreateCompetitionRepresentation() {

}

func (repo PostgresPartnershipCompetitionRepresentationRepository) SearchCompetitionRepresentation() {

}
