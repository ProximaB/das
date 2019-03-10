package partnershipdal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/accountdal"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/DancesportSoftware/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
	"log"
)

const (
	DasPartnershipTable                            = "DAS.PARTNERSHIP"
	partnershipColumnLeadID                        = "LEAD_ID"
	partnershipColumnFollowID                      = "FOLLOW_ID"
	partnershipColumnSameSexIndicator              = "SAMESEX_IND"
	columnFavoriteByLead                           = "FAVORITE_BY_LEAD"
	columnFavoriteByFollow                         = "FAVORITE_BY_FOLLOW"
	columnCompetitionsAttended                     = "COMPETITIONS_ATTENDED"
	columnEventsAttended                           = "EVENTS_ATTENDED"
	DAS_PARTNERSHIP_REQUEST_BLACKLIST_REASON_TABLE = "DAS.PARTNERSHIP_REQUEST_BLACKLIST_REASON"
)

// PostgresPartnershipRepository implements IPartnershipRepository
type PostgresPartnershipRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

// CreatePartnership creates the specified partnership in Postgres database and updates the ID
func (repo PostgresPartnershipRepository) CreatePartnership(partnership *businesslogic.Partnership) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	clause := repo.SqlBuilder.Insert("").
		Into(DasPartnershipTable).
		Columns(
			partnershipColumnLeadID,
			partnershipColumnFollowID,
			partnershipColumnSameSexIndicator,
			columnFavoriteByLead,
			columnFavoriteByFollow,
			columnCompetitionsAttended,
			columnEventsAttended,
			common.ColumnDateTimeCreated,
			common.ColumnDateTimeUpdated).
		Values(
			partnership.Lead.ID,
			partnership.Follow.ID,
			partnership.SameSex,
			partnership.FavoriteByLead,
			partnership.FavoriteByFollow,
			partnership.CompetitionsAttended,
			partnership.EventsAttended,
			partnership.DateTimeCreated,
			partnership.DateTimeUpdated)

	_, err := clause.RunWith(repo.Database).Exec()
	return err
}

// SearchPartnership searches partnerships in a Postgres database based on the criteria of search
func (repo PostgresPartnershipRepository) SearchPartnership(criteria businesslogic.SearchPartnershipCriteria) ([]businesslogic.Partnership, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Select(fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s",
		common.ColumnPrimaryKey,
		partnershipColumnLeadID,
		partnershipColumnFollowID,
		partnershipColumnSameSexIndicator,
		columnFavoriteByLead,
		columnFavoriteByFollow,
		columnCompetitionsAttended,
		columnEventsAttended,
		common.ColumnDateTimeCreated,
		common.ColumnDateTimeUpdated)).From(DasPartnershipTable)
	if criteria.PartnershipID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.PartnershipID})
	}

	if criteria.AccountID > 0 {
		// this will overrides the search by lead ID or by follow ID
		stmt = stmt.Where(squirrel.Or{
			squirrel.Eq{partnershipColumnLeadID: criteria.AccountID},
			squirrel.Eq{partnershipColumnFollowID: criteria.AccountID},
		})
	} else {
		if criteria.LeadID > 0 {
			stmt = stmt.Where(squirrel.Eq{partnershipColumnLeadID: criteria.LeadID})
		}
		if criteria.FollowID > 0 {
			stmt = stmt.Where(squirrel.Eq{partnershipColumnFollowID: criteria.FollowID})
		}
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
		scanErr := rows.Scan(
			&each.ID,
			&each.Lead.ID,
			&each.Follow.ID,
			&each.SameSex,
			&each.FavoriteByLead,
			&each.FavoriteByFollow,
			&each.CompetitionsAttended,
			&each.EventsAttended,
			&each.DateTimeCreated,
			&each.DateTimeUpdated,
		)
		if scanErr != nil {
			return partnerships, scanErr
		}
		leads, searchLeadErr := accountRepo.SearchAccount(businesslogic.SearchAccountCriteria{ID: each.Lead.ID})
		follows, searchFollowErr := accountRepo.SearchAccount(businesslogic.SearchAccountCriteria{ID: each.Follow.ID})

		if searchLeadErr != nil {
			log.Printf("[error] %v", searchLeadErr)
		} else if len(leads) != 1 {
			log.Printf("[warning] cannot find the lead with account ID: %d", each.Lead.ID)
		} else {
			each.Lead = leads[0]
		}
		if searchFollowErr != nil {
			log.Printf("[error] %v", searchFollowErr)
		} else if len(follows) != 1 {
			log.Printf("[warning] cannot find the follow with account ID: %d", each.Follow.ID)
		} else {
			each.Follow = follows[0]
		}

		partnerships = append(partnerships, each)
	}
	closeErr := rows.Close()
	return partnerships, closeErr
}

func (repo PostgresPartnershipRepository) DeletePartnership(partnership businesslogic.Partnership) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

func (repo PostgresPartnershipRepository) UpdatePartnership(partnership businesslogic.Partnership) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Update("").Table(DasPartnershipTable).
		Set(columnFavoriteByLead, partnership.FavoriteByLead).
		Set(columnFavoriteByFollow, partnership.FavoriteByFollow).
		Where(squirrel.Eq{common.ColumnPrimaryKey: partnership.ID})
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		if _, exeErr := stmt.RunWith(repo.Database).Exec(); exeErr != nil {
			return exeErr
		}
		if commitErr := tx.Commit(); commitErr != nil {
			return commitErr
		}
	}
	return nil
}
