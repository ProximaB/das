package entrydal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/DancesportSoftware/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
	"log"
)

const (
	dasCompetitionLeadTagTable                = "DAS.COMPETITION_LEAD_TAG"
	dasCompetitionLeadTagTableColumnLeadID    = "DAS.COMPETITION_LEAD_TAG.LEAD_ID"
	dasCompetitionLeadTagTableColumnTagNumber = "DAS.COMPETITION_LEAD_TAG.TAG_NUMBER"
)

type PostgresCompetitionLeadTagRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

// CreateCompetitionLeadTag creates the provided tag in a Postgres database.
func (repo PostgresCompetitionLeadTagRepository) CreateCompetitionLeadTag(tag *businesslogic.CompetitionLeadTag) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.Insert("").
		Into(dasCompetitionLeadTagTable).
		Columns(
			common.COL_COMPETITION_ID,
			dasCompetitionLeadTagTableColumnLeadID,
			dasCompetitionLeadTagTableColumnTagNumber,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated).
		Values(
			tag.CompetitionID,
			tag.LeadID,
			tag.Tag,
			tag.CreateUserID,
			tag.DateTimeCreated,
			tag.UpdateUserID,
			tag.DateTimeUpdated).Suffix(dalutil.SQLSuffixReturningID)
	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		scanErr := row.Scan(&tag.ID)
		if scanErr != nil {
			return scanErr
		}
		if commitErr := tx.Commit(); commitErr != nil {
			return commitErr
		}
	}
	return err
}

// DeleteCompetitionLeadTag deletes the tag from a Postgres database by the ID of the tag
func (repo PostgresCompetitionLeadTagRepository) DeleteCompetitionLeadTag(tag businesslogic.CompetitionLeadTag) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.Delete("").From(dasCompetitionLeadTagTable)
	var err error
	if tag.ID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: tag.ID})
		if tx, txErr := repo.Database.Begin(); txErr != nil {
			return txErr
		} else {
			_, err = stmt.RunWith(repo.Database).Exec()
			if err != nil {
				log.Printf("[error] go error while deleting Competition Lead Tag with ID = %v: %v", tag.ID, err)
				return err
			}
			return tx.Commit()
		}
	}
	return err
}

// SearchCompetitionLeadTag searches lead tags that meet the requirement of the specified critieria in a Postgres database
func (repo PostgresCompetitionLeadTagRepository) SearchCompetitionLeadTag(criteria businesslogic.SearchCompetitionLeadTagCriteria) (businesslogic.CompetitionLeadTagCollection, error) {
	tags := make([]businesslogic.CompetitionLeadTag, 0)
	collection := businesslogic.CompetitionLeadTagCollection{}
	collection.SetTags(tags)
	if repo.Database == nil {
		return collection, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}

	clause := repo.SQLBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s",
		common.ColumnPrimaryKey,
		common.COL_COMPETITION_ID,
		dasCompetitionLeadTagTableColumnLeadID,
		dasCompetitionLeadTagTableColumnTagNumber,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated)).From(dasCompetitionLeadTagTable)

	if criteria.ID > 0 {
		clause = clause.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.ID})
	}
	if criteria.CompetitionID > 0 {
		clause = clause.Where(squirrel.Eq{dasCompetitionLeadTagTableColumnTagNumber: criteria.CompetitionID})
	}
	if criteria.LeadID > 0 {
		clause = clause.Where(squirrel.Eq{dasCompetitionLeadTagTableColumnLeadID: criteria.LeadID})
	}
	if criteria.Tag > 0 {
		clause = clause.Where(squirrel.Eq{dasCompetitionLeadTagTableColumnTagNumber: criteria.Tag})
	}
	if criteria.CreateUserID > 0 {
		clause = clause.Where(squirrel.Eq{common.ColumnCreateUserID: criteria.CreateUserID})
	}

	rows, err := clause.RunWith(repo.Database).Query()
	if err != nil {
		return collection, err
	}

	for rows.Next() {
		each := businesslogic.CompetitionLeadTag{}
		scanErr := rows.Scan(
			&each.ID,
			&each.CompetitionID,
			&each.LeadID,
			&each.Tag,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated)
		if scanErr != nil {
			return collection, err
		}
		tags = append(tags, each)
	}
	err = rows.Close()

	return collection, err
}

// UpdateCompetitionLeadTag updates the tag of the provided ID to the new tag property in Postgres database
func (repo PostgresCompetitionLeadTagRepository) UpdateCompetitionLeadTag(tag businesslogic.CompetitionLeadTag) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.Update("").Table(dasCompetitionLeadTagTable)
	if tag.ID > 0 {
		stmt = stmt.Set(common.COL_COMPETITION_ID, tag.CompetitionID).
			Set(dasCompetitionLeadTagTableColumnLeadID, tag.LeadID).
			Set(dasCompetitionLeadTagTableColumnTagNumber, tag.Tag).
			Set(common.ColumnUpdateUserID, tag.DateTimeCreated).
			Set(common.ColumnDateTimeUpdated, tag.DateTimeUpdated)
	} else {
		return errors.New("ID of CompetitionLeadTag must be specified")
	}
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		if commitErr := tx.Commit(); commitErr != nil {
			return commitErr
		}
	}
	return err
}
