package organizer

import (
	"database/sql"
	"errors"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
)

type PostgresCompetitionOfficialInvitationRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresCompetitionOfficialInvitationRepository) CreateCompetitionOfficialInvitationRepository(invitation *businesslogic.CompetitionOfficialInvitation) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}
func (repo PostgresCompetitionOfficialInvitationRepository) DeleteCompetitionOfficialInvitationRepository(invitation businesslogic.CompetitionOfficialInvitation) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}
func (repo PostgresCompetitionOfficialInvitationRepository) SearchCompetitionOfficialInvitationRepository(criteria businesslogic.SearchCompetitionOfficialInvitationCriteria) ([]businesslogic.CompetitionOfficialInvitation, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return nil, errors.New("not implemented")
}
func (repo PostgresCompetitionOfficialInvitationRepository) UpdateCompetitionOfficialInvitationRepository(invitation businesslogic.CompetitionOfficialInvitation) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}
