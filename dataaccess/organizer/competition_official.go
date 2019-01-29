package organizer

import (
	"database/sql"
	"errors"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/Masterminds/squirrel"
)

type PostgresCompetitionOfficialRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresCompetitionOfficialRepository) CreateCompetitionOfficial(official *businesslogic.CompetitionOfficial) error {
	return errors.New("not implemented")
}

func (repo PostgresCompetitionOfficialRepository) DeleteCompetitionOfficial(official businesslogic.CompetitionOfficial) error {
	return errors.New("not implemented")
}

func (repo PostgresCompetitionOfficialRepository) SearchCompetitionOfficial(criteria businesslogic.SearchCompetitionOfficialCriteria) ([]businesslogic.CompetitionOfficial, error) {
	return nil, errors.New("not implemented")
}

func (repo PostgresCompetitionOfficialRepository) UpdateCompetitionOfficial(official businesslogic.CompetitionOfficial) error {
	return errors.New("not implemented")
}
