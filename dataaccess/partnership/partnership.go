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

package partnership

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/Masterminds/squirrel"
	"time"
)

const (
	DasPartnershipTable                  = "DAS.PARTNERSHIP"
	DasPartnershipColumnLeadID           = "LEAD_ID"
	DasPartnershipColumnFollowID         = "FOLLOW_ID"
	DasPartnershipColumnSameSexIndicator = "SAMESEX_IND"
	DasPartnershipColumnFavorite         = "FAVORITE"
)

const (
	DAS_PARTNERSHIP_REQUEST_BLACKLIST_REASON_TABLE = "DAS.PARTNERSHIP_REQUEST_BLACKLIST_REASON"
)

type PostgresPartnershipRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresPartnershipRepository) CreatePartnership(partnership *businesslogic.Partnership) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresPartnershipRepository is not specified")
	}
	clause := repo.SqlBuilder.Insert("").
		Into(DasPartnershipTable).
		Columns(
			DasPartnershipColumnLeadID,
			DasPartnershipColumnFollowID,
			DasPartnershipColumnSameSexIndicator,
			DasPartnershipColumnFavorite,
			common.COL_DATETIME_CREATED,
			common.COL_DATETIME_UPDATED).Values(partnership.LeadID, partnership.FollowID, partnership.SameSex, partnership.FavoriteByLead, partnership.DateTimeCreated, time.Now())

	_, err := clause.RunWith(repo.Database).Exec()
	return err
}

func (repo PostgresPartnershipRepository) SearchPartnership(criteria businesslogic.SearchPartnershipCriteria) ([]businesslogic.Partnership, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresPartnershipRepository is not specified")
	}
	stmt := repo.SqlBuilder.Select(fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s",
		common.ColumnPrimaryKey,
		DasPartnershipColumnLeadID,
		DasPartnershipColumnFollowID,
		DasPartnershipColumnSameSexIndicator,
		DasPartnershipColumnFavorite,
		common.COL_DATETIME_CREATED,
		common.COL_DATETIME_UPDATED)).From(DasPartnershipTable)
	if criteria.PartnershipID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.PartnershipID})
	}
	if criteria.LeadID > 0 {
		stmt = stmt.Where(squirrel.Eq{DasPartnershipColumnLeadID: criteria.LeadID})
	}
	if criteria.FollowID > 0 {
		stmt = stmt.Where(squirrel.Eq{DasPartnershipColumnFollowID: criteria.FollowID})
	}

	partnerships := make([]businesslogic.Partnership, 0)
	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return partnerships, err
	}

	for rows.Next() {
		each := businesslogic.Partnership{}
		rows.Scan(
			&each.ID,
			&each.LeadID,
			&each.FollowID,
			&each.SameSex,
			&each.FavoriteByLead,
			&each.DateTimeCreated,
			&each.DateTimeUpdated,
		)
		partnerships = append(partnerships, each)
	}
	rows.Close()
	return partnerships, err
}

func (repo PostgresPartnershipRepository) DeletePartnership(partnership businesslogic.Partnership) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresPartnershipRepository is not specified")
	}
	return errors.New("not implemented")
}

func (repo PostgresPartnershipRepository) UpdatePartnership(partnership businesslogic.Partnership) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresPartnershipRepository is not specified")
	}
	return errors.New("not implemented")
}
