package entrydal

import (
	"database/sql"
	"errors"

	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
)

// PostgresPartnershipRoundEntryRepository implements the IPartnershipRoundEntryRepository with a Postgres database
type PostgresPartnershipRoundEntryRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

// CreatePartnershipRoundEntry creates a PartnershipRoundEntry in a Postgres database
func (repo PostgresPartnershipRoundEntryRepository) CreatePartnershipRoundEntry(entry *businesslogic.PartnershipRoundEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// DeletePartnershipRoundEntry deletes a PartnershipRoundEntry from a Postgres database
func (repo PostgresPartnershipRoundEntryRepository) DeletePartnershipRoundEntry(entry businesslogic.PartnershipRoundEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// SearchPartnershipRoundEntry searches PartnershipRoundEntry in a Postgres database
func (repo PostgresPartnershipRoundEntryRepository) SearchPartnershipRoundEntry(criteria businesslogic.SearchPartnershipRoundEntryCriteria) ([]businesslogic.PartnershipRoundEntry, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return nil, errors.New("not implemented")
}

// UpdatePartnershipRoundEntry updates a PartnershipRoundEntry in a Postgres database
func (repo PostgresPartnershipRoundEntryRepository) UpdatePartnershipRoundEntry(entry businesslogic.PartnershipRoundEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// PostgresAdjudicatorRoundEntryRepository implements IAdjudicatorRoundEntryRepository with a Postgres database
type PostgresAdjudicatorRoundEntryRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

// CreateAdjudicatorRoundEntry creates an AdjudicatorRoundEntry in a Postgres database
func (repo PostgresAdjudicatorRoundEntryRepository) CreateAdjudicatorRoundEntry(entry *businesslogic.AdjudicatorRoundEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// DeleteAdjudicatorRoundEntry deletes an AdjudicatorRoundEntry from a Postgres database
func (repo PostgresAdjudicatorRoundEntryRepository) DeleteAdjudicatorRoundEntry(entry businesslogic.AdjudicatorRoundEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// SearchAdjudicatorRoundEntry searches AdjudicatorRoundEntry from a Postgres database
func (repo PostgresAdjudicatorRoundEntryRepository) SearchAdjudicatorRoundEntry(criteria businesslogic.SearchAdjudicatorRoundEntryCriteria) ([]businesslogic.AdjudicatorRoundEntry, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return nil, errors.New("not implemented")
}

// UpdateAdjudicatorRoundEntry updates an AdjudicatorRoundEntry from a Postgres database
func (repo PostgresAdjudicatorRoundEntryRepository) UpdateAdjudicatorRoundEntry(entry businesslogic.AdjudicatorRoundEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}
