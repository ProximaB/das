package entrydal

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/DancesportSoftware/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
)

const (
	dasAthleteCompetitionEntryTable       = "DAS.COMPETITION_ENTRY_ATHLETE"
	dasPartnershipCompetitionEntryTable   = "DAS.COMPETITION_ENTRY_PARTNERSHIP"
	dasAdjudicatorCompetitionEntryTable   = "DAS.COMPETITION_ENTRY_ADJUDICATOR"
	dasScrutineerCompetitionEntryTable    = "DAS.COMPETITION_ENTRY_SCRUTINEER"
	dasCompetitionEntryColCheckinInd      = "CHECKIN_IND"
	dasCompetitionEntryColCheckinDateTime = "CHECKIN_DATETIME"
	dasCompetitionEntryColCompetitorTag   = "LEADTAG"
)

// PostgresAthleteCompetitionEntryRepository is a Postgres-based Athlete Competition Entry Repository
type PostgresAthleteCompetitionEntryRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

// CreateEntry creates an AthleteCompetitionEntry in a Postgres database
func (repo PostgresAthleteCompetitionEntryRepository) CreateEntry(entry *businesslogic.AthleteCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.Insert("").
		Into(dasAthleteCompetitionEntryTable).
		Columns(
			common.COL_COMPETITION_ID,
			"ATHLETE_ID",
			dasCompetitionEntryColCheckinInd,
			dasCompetitionEntryColCheckinDateTime,
			"PAYMENT_IND",
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated).
		Values(
			entry.CompetitionEntry.CompetitionID,
			entry.AthleteID,
			entry.CompetitionEntry.CheckInIndicator,
			entry.CompetitionEntry.DateTimeCheckIn,
			entry.PaymentReceivedIndicator,
			entry.CompetitionEntry.CreateUserID,
			entry.CompetitionEntry.DateTimeCreated,
			entry.CompetitionEntry.UpdateUserID,
			entry.CompetitionEntry.DateTimeUpdated).
		Suffix("RETURNING ID")

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		scanErr := row.Scan(&entry.ID)
		if scanErr != nil {
			return scanErr
		}
		if commitErr := tx.Commit(); commitErr != nil {
			return commitErr
		}
	}
	return err
}

// SearchEntry searches AthleteCompetitionEntry in a Postgres database
func (repo PostgresAthleteCompetitionEntryRepository) SearchEntry(criteria businesslogic.SearchAthleteCompetitionEntryCriteria) ([]businesslogic.AthleteCompetitionEntry, error) {
	entries := make([]businesslogic.AthleteCompetitionEntry, 0)

	if repo.Database == nil {
		return entries, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	clause := repo.SQLBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s",
		common.ColumnPrimaryKey,
		common.COL_COMPETITION_ID,
		common.ColumnAccountID,
		dasCompetitionEntryColCheckinInd,
		dasCompetitionEntryColCheckinDateTime,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated)).From(dasAthleteCompetitionEntryTable)

	if criteria.ID > 0 {
		clause = clause.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.ID})
	}
	if criteria.AthleteID > 0 {
		clause = clause.Where(squirrel.Eq{common.ColumnAccountID: criteria.AthleteID})
	}
	if criteria.CompetitionID > 0 {
		clause = clause.Where(squirrel.Eq{common.COL_COMPETITION_ID: criteria.CompetitionID})
	}

	rows, err := clause.RunWith(repo.Database).Query()
	if err != nil {
		return entries, err
	}

	for rows.Next() {
		each := businesslogic.AthleteCompetitionEntry{
			CompetitionEntry: businesslogic.BaseCompetitionEntry{},
		}
		scanErr := rows.Scan(
			&each.ID,
			&each.CompetitionEntry.CompetitionID,
			&each.AthleteID,
			&each.CompetitionEntry.CheckInIndicator,
			&each.CompetitionEntry.DateTimeCheckIn,
			&each.CompetitionEntry.CreateUserID,
			&each.CompetitionEntry.DateTimeCreated,
			&each.CompetitionEntry.UpdateUserID,
			&each.CompetitionEntry.DateTimeUpdated,
		)
		if scanErr != nil {
			return entries, scanErr
		}
		entries = append(entries, each)
	}
	return entries, err
}

// DeleteEntry deletes an AthleteCompetitionEntry from a Postgres database
func (repo PostgresAthleteCompetitionEntryRepository) DeleteEntry(entry businesslogic.AthleteCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.Delete("").From(dasAthleteCompetitionEntryTable)
	var err error
	if entry.ID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: entry.ID})
		if tx, txErr := repo.Database.Begin(); txErr != nil {
			return txErr
		} else {
			_, err = stmt.RunWith(repo.Database).Exec()
			if err != nil {
				log.Printf("[error] got error while deleting competition entry with ID")
				return err
			}
			return tx.Commit()
		}
	}
	return err
}

// UpdateEntry updates an AthleteCompetitionEntry from a Postgres database
func (repo PostgresAthleteCompetitionEntryRepository) UpdateEntry(entry businesslogic.AthleteCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.Update(dasAthleteCompetitionEntryTable)
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		err = tx.Commit()
	}
	return err
}

// PostgresPartnershipCompetitionEntryRepository implements a IPartnershipCompetitionEntryRepository with Postgres database
type PostgresPartnershipCompetitionEntryRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

// CreateEntry creates a PartnershipCompetitionEntry in a Postgres database
func (repo PostgresPartnershipCompetitionEntryRepository) CreateEntry(entry *businesslogic.PartnershipCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.
		Insert("").
		Into(dasPartnershipCompetitionEntryTable).
		Columns(
			common.COL_COMPETITION_ID,
			common.COL_PARTNERSHIP_ID,
			dasCompetitionEntryColCheckinInd,
			dasCompetitionEntryColCheckinDateTime,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated).
		Values(
			entry.CompetitionEntry.CompetitionID,
			entry.PartnershipID,
			entry.CompetitionEntry.CheckInIndicator,
			entry.CompetitionEntry.DateTimeCheckIn,
			entry.CompetitionEntry.CreateUserID,
			entry.CompetitionEntry.DateTimeCreated,
			entry.CompetitionEntry.UpdateUserID,
			entry.CompetitionEntry.DateTimeUpdated).
		Suffix(dalutil.SQLSuffixReturningID)

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&entry.ID)
		tx.Commit()
	}
	return err
}

// DeleteEntry deletes a PartnershipCompetitionEntry from a Postgres database
func (repo PostgresPartnershipCompetitionEntryRepository) DeleteEntry(entry businesslogic.PartnershipCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.Delete("").From(dasPartnershipCompetitionEntryTable)
	if entry.ID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: entry.ID})
	} else {
		return errors.New("cannot find this partnership competition entry")
	}

	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return err
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		err = tx.Commit()
	}
	return err
}

// SearchEntry searches PartnershipCompetitionEntry in a Postgres database
func (repo PostgresPartnershipCompetitionEntryRepository) SearchEntry(criteria businesslogic.SearchPartnershipCompetitionEntryCriteria) ([]businesslogic.PartnershipCompetitionEntry, error) {
	entries := make([]businesslogic.PartnershipCompetitionEntry, 0)

	if repo.Database == nil {
		return entries, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}

	clause := repo.SQLBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s",
		common.ColumnPrimaryKey,
		common.COL_COMPETITION_ID,
		common.COL_PARTNERSHIP_ID,
		dasCompetitionEntryColCheckinInd,
		dasCompetitionEntryColCheckinDateTime,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated)).From(dasPartnershipCompetitionEntryTable)

	if criteria.ID > 0 {
		clause = clause.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.ID})
	}
	if criteria.PartnershipID > 0 {
		clause = clause.Where(squirrel.Eq{common.COL_PARTNERSHIP_ID: criteria.PartnershipID})
	}
	if criteria.CompetitionID > 0 {
		clause = clause.Where(squirrel.Eq{common.COL_COMPETITION_ID: criteria.CompetitionID})
	}

	rows, err := clause.RunWith(repo.Database).Query()
	if err != nil {
		return entries, err
	}

	for rows.Next() {
		each := businesslogic.PartnershipCompetitionEntry{
			CompetitionEntry: businesslogic.BaseCompetitionEntry{},
		}
		scanErr := rows.Scan(
			&each.ID,
			&each.CompetitionEntry.CompetitionID,
			&each.PartnershipID,
			&each.CompetitionEntry.CheckInIndicator,
			&each.CompetitionEntry.DateTimeCheckIn,
			&each.CompetitionEntry.CreateUserID,
			&each.CompetitionEntry.DateTimeCreated,
			&each.CompetitionEntry.UpdateUserID,
			&each.CompetitionEntry.DateTimeUpdated,
		)
		if scanErr != nil {
			return entries, scanErr
		}
		entries = append(entries, each)
	}

	return nil, err
}

// UpdateEntry updates a PartnershipCompetitionEntry in a Postgres database
func (repo PostgresPartnershipCompetitionEntryRepository) UpdateEntry(entry businesslogic.PartnershipCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// PostgresAdjudicatorCompetitionEntryRepository implements the IAdjudicatorCompetitionEntryRepository with a Postgres database
type PostgresAdjudicatorCompetitionEntryRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

// CreateEntry creates an AdjudicatorCompetitionEntry in a Postgres database
func (repo PostgresAdjudicatorCompetitionEntryRepository) CreateEntry(entry *businesslogic.AdjudicatorCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// DeleteEntry deletes an AdjudicatorCompetitionEntry from a Postgres database
func (repo PostgresAdjudicatorCompetitionEntryRepository) DeleteEntry(entry businesslogic.AdjudicatorCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// SearchEntry searches AdjudicatorCompetitionEntry in a Postgres database
func (repo PostgresAdjudicatorCompetitionEntryRepository) SearchEntry(criteria businesslogic.SearchAdjudicatorCompetitionEntryCriteria) ([]businesslogic.AdjudicatorCompetitionEntry, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return nil, errors.New("not implemented")
}

// UpdateEntry updates an AdjudicatorCompetitionEntry in a Postgres database
func (repo PostgresAdjudicatorCompetitionEntryRepository) UpdateEntry(entry businesslogic.AdjudicatorCompetitionEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}
