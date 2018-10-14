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

package partnershipdal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/accountdal"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/Masterminds/squirrel"
	"log"
	"time"
)

const (
	DasPartnershipTable               = "DAS.PARTNERSHIP"
	partnershipColumnLeadID           = "LEAD_ID"
	partnershipColumnFollowID         = "FOLLOW_ID"
	partnershipColumnSameSexIndicator = "SAMESEX_IND"
	columnFavoriteByLead              = "FAVORITE_BY_LEAD"
	columnFavoriteByFollow            = "FAVORITE_BY_FOLLOW"
	columnCompetitionsAttended        = "COMPETITIONS_ATTENDED"
	columnEventsAttended              = "EVENTS_ATTENDED"
)

const (
	DAS_PARTNERSHIP_REQUEST_BLACKLIST_REASON_TABLE = "DAS.PARTNERSHIP_REQUEST_BLACKLIST_REASON"
)

type PostgresPartnershipRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

// CreatePartnership creates the specified partnership in Postgres database and updates the ID
func (repo PostgresPartnershipRepository) CreatePartnership(partnership *businesslogic.Partnership) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresPartnershipRepository is not specified")
	}
	clause := repo.SqlBuilder.Insert("").
		Into(DasPartnershipTable).
		Columns(
			partnershipColumnLeadID,
			partnershipColumnFollowID,
			partnershipColumnSameSexIndicator,
			columnFavoriteByLead,
			columnFavoriteByFollow,
			common.ColumnDateTimeCreated,
			common.ColumnDateTimeUpdated).
		Values(
			partnership.LeadID,
			partnership.FollowID,
			partnership.SameSex,
			partnership.FavoriteByLead,
			partnership.FavoriteByFollow,
			partnership.DateTimeCreated,
			time.Now())

	_, err := clause.RunWith(repo.Database).Exec()
	return err
}

// SearchPartnership searches partnerships in a Postgres database based on the criteria of search
func (repo PostgresPartnershipRepository) SearchPartnership(criteria businesslogic.SearchPartnershipCriteria) ([]businesslogic.Partnership, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresPartnershipRepository is not specified")
	}
	stmt := repo.SqlBuilder.Select(fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s",
		common.ColumnPrimaryKey,
		partnershipColumnLeadID,
		partnershipColumnFollowID,
		partnershipColumnSameSexIndicator,
		columnFavoriteByLead,
		columnFavoriteByFollow,
		common.ColumnDateTimeCreated,
		common.ColumnDateTimeUpdated)).From(DasPartnershipTable)
	if criteria.PartnershipID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.PartnershipID})
	}
	if criteria.LeadID > 0 {
		stmt = stmt.Where(squirrel.Eq{partnershipColumnLeadID: criteria.LeadID})
	}
	if criteria.FollowID > 0 {
		stmt = stmt.Where(squirrel.Eq{partnershipColumnFollowID: criteria.FollowID})
	}

	// get account
	accountRepo := accountdal.PostgresAccountRepository{
		Database:   repo.Database,
		SQLBuilder: repo.SqlBuilder,
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
			&each.FavoriteByFollow,
			&each.DateTimeCreated,
			&each.DateTimeUpdated,
		)
		leads, searchLeadErr := accountRepo.SearchAccount(businesslogic.SearchAccountCriteria{ID: each.LeadID})
		follows, searchFollowErr := accountRepo.SearchAccount(businesslogic.SearchAccountCriteria{ID: each.FollowID})

		if searchLeadErr != nil {
			log.Printf("[error] %v", searchLeadErr)
		} else if len(leads) != 1 {
			log.Printf("[warning] cannot find the lead with account ID: %d", each.LeadID)
		} else {
			each.Lead = leads[0]
		}
		if searchFollowErr != nil {
			log.Printf("[error] %v", searchFollowErr)
		} else if len(follows) != 1 {
			log.Printf("[warning] cannot find the follow with account ID: %d", each.FollowID)
		} else {
			each.Follow = follows[0]
		}
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
