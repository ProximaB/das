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

package entry

import (
	"database/sql"
	"errors"

	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/Masterminds/squirrel"
)

type PostgresPartnershipRoundEntryRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
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
	SQLBuilder squirrel.StatementBuilderType
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
