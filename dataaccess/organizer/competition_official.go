package organizer

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/dataaccess/common"
	"github.com/ProximaB/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
	"log"
)

const (
	dasCompetitionOfficialTable = "DAS.COMPETITION_OFFICIAL"
	columnOfficialAccountID     = "OFFICIAL_ACCOUNT_ID"
	columnOfficialRoleID        = "OFFICIAL_ROLE_ID"
	columnEffectiveFrom         = "EFFECTIVE_FROM"
	columnEffectiveUntil        = "EFFECTIVE_UNTIL"
	columnAssignedBy            = "ASSIGNED_BY"
)

type PostgresCompetitionOfficialRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresCompetitionOfficialRepository) CreateCompetitionOfficial(official *businesslogic.CompetitionOfficial) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Insert("").
		Into(dasCompetitionOfficialTable).
		Columns(
			common.COL_COMPETITION_ID,
			columnOfficialAccountID,
			columnOfficialRoleID,
			columnEffectiveFrom,
			columnEffectiveUntil,
			columnAssignedBy,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated).
		Values(
			official.Competition.ID,
			official.Official.ID,
			official.OfficialRoleID,
			official.EffectiveFrom,
			official.EffectiveUntil,
			official.AssignedBy,
			official.CreateUserID,
			official.DateTimeCreated,
			official.UpdateUserID,
			official.DateTimeUpdated).
		Suffix(dalutil.SQLSuffixReturningID)
	hasError := false
	clause, args, sqlErr := stmt.ToSql()
	if sqlErr != nil {
		log.Printf("[error] generating SQL clause: %v", sqlErr)
		hasError = true
	}
	tx, txErr := repo.Database.Begin()
	if txErr != nil {
		log.Printf("[error] starting transaction: %v", txErr)
		hasError = true
	}
	row := repo.Database.QueryRow(clause, args...)
	if scanErr := row.Scan(&official.ID); scanErr != nil {
		log.Printf("[error] scanning ID of newly created Competition Official: %v", scanErr)
		hasError = true
	}
	if commitErr := tx.Commit(); commitErr != nil {
		log.Printf("[error] commiting transaction: %v", commitErr)
		hasError = true
	}
	if hasError {
		return errors.New("An error occurred while creating competition official record")
	}
	return nil

}

func (repo PostgresCompetitionOfficialRepository) DeleteCompetitionOfficial(official businesslogic.CompetitionOfficial) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

func (repo PostgresCompetitionOfficialRepository) SearchCompetitionOfficial(criteria businesslogic.SearchCompetitionOfficialCriteria) ([]businesslogic.CompetitionOfficial, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Select(
		fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
			common.ColumnPrimaryKey,
			common.COL_COMPETITION_ID,
			columnOfficialAccountID,
			columnOfficialRoleID,
			columnEffectiveFrom,
			columnEffectiveUntil,
			columnAssignedBy,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated,
		)).From(dasCompetitionOfficialTable)
	if criteria.ID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.ID})
	}
	if criteria.OfficialID != "" {
		// TODO: implement this logic: joining account table on UUID and ID
	}
	if criteria.CompetitionID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_COMPETITION_ID: criteria.CompetitionID})
	}
	if criteria.OfficialRoleID > 0 {
		stmt = stmt.Where(squirrel.Eq{columnOfficialRoleID: criteria.OfficialRoleID})
	}

	officials := make([]businesslogic.CompetitionOfficial, 0)

	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		each := businesslogic.CompetitionOfficial{
			Official:    businesslogic.Account{},
			Competition: businesslogic.Competition{},
		}
		scanerr := rows.Scan(
			&each.ID,
			&each.Competition.ID,
			&each.Official.ID,
			&each.OfficialRoleID,
			&each.EffectiveFrom,
			&each.EffectiveUntil,
			&each.AssignedBy,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		if scanerr != nil {
			log.Printf("[error] scanning Competition Official: %v", scanerr)
			return officials, errors.New("An error occurred in reading data")
		}
		officials = append(officials, each)
	}
	rows.Close()
	return officials, err
}

func (repo PostgresCompetitionOfficialRepository) UpdateCompetitionOfficial(official businesslogic.CompetitionOfficial) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}
