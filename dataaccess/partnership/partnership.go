package partnership

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/dataaccess/common"
	"database/sql"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"time"
)

const (
	DAS_PARTNERSHIP_TABLE              = "DAS.PARTNERSHIP"
	DAS_PARTNERSHIP_COL_PARTNERSHIP_ID = "PARTNERSHIP_ID"
	DAS_PARTNERSHIP_COL_LEAD_ID        = "LEAD_ID"
	DAS_PARTNERSHIP_COL_FOLLOW_ID      = "FOLLOW_ID"
	DAS_PARTNERSHIP_COL_SAMESEX_IND    = "SAMESEX_IND"
	DAS_PARTNERSHIP_COL_FAVORITE       = "FAVORITE"
)

const (
	DAS_PARTNERSHIP_REQUEST_BLACKLIST_REASON_TABLE = "DAS.PARTNERSHIP_REQUEST_BLACKLIST_REASON"
)

type PostgresPartnershipRepository struct {
	Database   *sql.DB
	SqlBuilder sq.StatementBuilderType
}

func (repo PostgresPartnershipRepository) CreatePartnership(partnership businesslogic.Partnership) error {
	clause := repo.SqlBuilder.Insert("").
		Into(DAS_PARTNERSHIP_TABLE).
		Columns(
			DAS_PARTNERSHIP_COL_LEAD_ID,
			DAS_PARTNERSHIP_COL_FOLLOW_ID,
			DAS_PARTNERSHIP_COL_SAMESEX_IND,
			DAS_PARTNERSHIP_COL_FAVORITE,
			common.COL_DATETIME_CREATED,
			common.COL_DATETIME_UPDATED).Values(partnership.LeadID, partnership.FollowID, partnership.SameSex, partnership.FavoriteLead, partnership.DateTimeCreated, time.Now())

	_, err := clause.RunWith(repo.Database).Exec()
	return err
}

func (repo PostgresPartnershipRepository) SearchPartnership(criteria *businesslogic.SearchPartnershipCriteria) ([]businesslogic.Partnership, error) {
	stmt := repo.SqlBuilder.Select(fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s",
		common.PRIMARY_KEY,
		DAS_PARTNERSHIP_COL_LEAD_ID,
		DAS_PARTNERSHIP_COL_FOLLOW_ID,
		DAS_PARTNERSHIP_COL_SAMESEX_IND,
		DAS_PARTNERSHIP_COL_FAVORITE,
		common.COL_DATETIME_CREATED,
		common.COL_DATETIME_UPDATED)).From(DAS_PARTNERSHIP_TABLE)
	if criteria.PartnershipID > 0 {
		stmt = stmt.Where(sq.Eq{common.PRIMARY_KEY: criteria.PartnershipID})
	}
	if criteria.LeadID > 0 {
		stmt = stmt.Where(sq.Eq{DAS_PARTNERSHIP_COL_LEAD_ID: criteria.LeadID})
	}
	if criteria.FollowID > 0 {
		stmt = stmt.Where(sq.Eq{DAS_PARTNERSHIP_COL_FOLLOW_ID: criteria.FollowID})
	}

	partnerships := make([]businesslogic.Partnership, 0)
	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return partnerships, err
	}

	for rows.Next() {
		each := businesslogic.Partnership{}
		rows.Scan(
			&each.PartnershipID,
			&each.LeadID,
			&each.FollowID,
			&each.SameSex,
			&each.FavoriteLead,
			&each.DateTimeCreated,
			&each.DateTimeUpdated,
		)
		partnerships = append(partnerships, each)
	}
	rows.Close()
	return partnerships, err
}

func (repo PostgresPartnershipRepository) DeletePartnership(partnership businesslogic.Partnership) error {
	return errors.New("not implemented")
}

func (repo PostgresPartnershipRepository) UpdatePartnership(partnership businesslogic.Partnership) error {
	return errors.New("not implemented")
}
