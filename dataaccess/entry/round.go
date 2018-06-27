// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package entry

import (
	"database/sql"
	"errors"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/Masterminds/squirrel"
)

type PostgresPartnershipRoundEntryRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresPartnershipRoundEntryRepository) CreatePartnershipRoundEntry(entry *businesslogic.PartnershipRoundEntry) error {
	return errors.New("not implemented")
}

func (repo PostgresPartnershipRoundEntryRepository) DeletePartnershipRoundEntry(entry businesslogic.PartnershipRoundEntry) error {
	return errors.New("not implemented")
}

func (repo PostgresPartnershipRoundEntryRepository) SearchPartnershipRoundEntry(criteria businesslogic.SearchPartnershipRoundEntryCriteria) ([]businesslogic.PartnershipRoundEntry, error) {
	return nil, errors.New("not implemented")
}

func (repo PostgresPartnershipRoundEntryRepository) UpdatePartnershipRoundEntry(entry businesslogic.PartnershipRoundEntry) error {
	return errors.New("not implemented")
}

type PostgresAdjudicatorRoundEntryRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresAdjudicatorRoundEntryRepository) CreateAdjudicatorRoundEntry(entry *businesslogic.AdjudicatorRoundEntry) error {
	return errors.New("not implemented")
}

func (repo PostgresAdjudicatorRoundEntryRepository) DeleteAdjudicatorRoundEntry(entry businesslogic.AdjudicatorRoundEntry) error {
	return errors.New("not implemented")
}

func (repo PostgresAdjudicatorRoundEntryRepository) SearchAdjudicatorRoundEntry(criteria businesslogic.SearchAdjudicatorRoundEntryCriteria) ([]businesslogic.AdjudicatorRoundEntry, error) {
	return nil, errors.New("not implemented")
}

func (repo PostgresAdjudicatorRoundEntryRepository) UpdateAdjudicatorRoundEntry(entry businesslogic.AdjudicatorRoundEntry) error {
	return errors.New("not implemented")
}
