package entrydal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ProximaB/das/dataaccess/accountdal"
	"github.com/ProximaB/das/dataaccess/competition"
	"github.com/ProximaB/das/dataaccess/partnershipdal"
	"log"

	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/dataaccess/common"
	"github.com/ProximaB/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
)

const (
	dasAthleteCompetitionEntryTable               = "DAS.COMPETITION_ENTRY_ATHLETE"
	dasPartnershipCompetitionEntryTable           = "DAS.COMPETITION_ENTRY_PARTNERSHIP"
	dasAdjudicatorCompetitionEntryTable           = "DAS.COMPETITION_ENTRY_ADJUDICATOR"
	dasScrutineerCompetitionEntryTable            = "DAS.COMPETITION_ENTRY_SCRUTINEER"
	dasCompetitionEntryColCheckinInd              = "CHECKIN_IND"
	dasCompetitionEntryColCheckinDateTime         = "CHECKIN_DATETIME"
	dasAthleteCompetitionEntryColumnLeadIndicator = "DAS.COMPETITION_ENTRY_ATHLETE.LEAD_INDICATOR"
	dasAthleteCompetitionEntryColumnLeadTag       = "DAS.COMPETITION_ENTRY_ATHLETE.LEAD_TAG"
	dasAthleteCompetitionEntryColumnOrganizerNote = "DAS.COMPETITION_ENTRY_ATHLETE.ORGANIZER_NOTE"
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
			common.COL_ATHLETE_ID,
			"LEAD_INDICATOR",
			"LEAD_TAG",
			"ORGANIZER_NOTE",
			dasCompetitionEntryColCheckinInd,
			dasCompetitionEntryColCheckinDateTime,
			"PAYMENT_IND",
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated).
		Values(
			entry.Competition.ID,
			entry.Athlete.ID,
			entry.IsLead,
			entry.LeadTag,
			entry.OrganizerNote,
			entry.CheckedIn,
			entry.DateTimeCheckedIn,
			entry.PaymentReceivedIndicator,
			entry.CreateUserID,
			entry.DateTimeCreated,
			entry.UpdateUserID,
			entry.DateTimeUpdated).
		Suffix(dalutil.SQLSuffixReturningID)

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
	clause := repo.SQLBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
		common.ColumnPrimaryKey,
		common.COL_COMPETITION_ID,
		common.COL_ATHLETE_ID,
		dasAthleteCompetitionEntryColumnLeadIndicator,
		dasAthleteCompetitionEntryColumnLeadTag,
		dasAthleteCompetitionEntryColumnOrganizerNote,
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
		clause = clause.Where(squirrel.Eq{common.COL_ATHLETE_ID: criteria.AthleteID})
	}
	if criteria.CompetitionID > 0 {
		clause = clause.Where(squirrel.Eq{common.COL_COMPETITION_ID: criteria.CompetitionID})
	}

	rows, err := clause.RunWith(repo.Database).Query()
	if err != nil {
		log.Printf("[error] generating search query for Athlete Competition Entry: %v", err)
		return entries, err
	}

	accountRepo := accountdal.PostgresAccountRepository{
		Database:   repo.Database,
		SQLBuilder: repo.SQLBuilder,
	}
	compRepo := competition.PostgresCompetitionRepository{
		Database:   repo.Database,
		SqlBuilder: repo.SQLBuilder,
	}

	for rows.Next() {
		each := businesslogic.AthleteCompetitionEntry{
			Competition: businesslogic.Competition{},
		}
		scanErr := rows.Scan(
			&each.ID,
			&each.Competition.ID,
			&each.Athlete.ID,
			&each.IsLead,
			&each.LeadTag,
			&each.OrganizerNote,
			&each.CheckedIn,
			&each.DateTimeCheckedIn,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		if scanErr != nil {
			return entries, scanErr
		}
		athletes, _ := accountRepo.SearchAccount(businesslogic.SearchAccountCriteria{ID: each.Athlete.ID})
		competitions, _ := compRepo.SearchCompetition(businesslogic.SearchCompetitionCriteria{ID: each.Competition.ID})
		each.Athlete = athletes[0]
		each.Competition = competitions[0]
		entries = append(entries, each)
	}
	rows.Close()
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
				log.Printf("[error] got error while deleting competition entry with ID = %v: %v", entry.ID, err)
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
	stmt := repo.SQLBuilder.Update("").Table(dasAthleteCompetitionEntryTable).
		Set(dasCompetitionEntryColCheckinInd, entry.CheckedIn).
		Set(dasCompetitionEntryColCheckinDateTime, entry.DateTimeCheckedIn).
		Set(dasAthleteCompetitionEntryColumnLeadIndicator, entry.IsLead).
		Set(dasAthleteCompetitionEntryColumnLeadTag, entry.LeadTag).
		Set(dasAthleteCompetitionEntryColumnOrganizerNote, entry.OrganizerNote)
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		err = tx.Commit()
	}
	return err
}

func (repo PostgresAthleteCompetitionEntryRepository) NextAvailableLeadTag(competition businesslogic.Competition) (int, error) {
	currentMaxTag := 0
	stmt := repo.SQLBuilder.Select("MAX(LEAD_TAG)").From(dasAthleteCompetitionEntryTable).Where(squirrel.Eq{common.COL_COMPETITION_ID: competition.ID})
	scanErr := stmt.RunWith(repo.Database).QueryRow().Scan(&currentMaxTag)
	if scanErr != nil {
		return currentMaxTag, scanErr
	}

	if currentMaxTag == 0 {
		currentMaxTag = 101
	} else {
		currentMaxTag += 1
	}

	return currentMaxTag, nil
}

func (repo PostgresAthleteCompetitionEntryRepository) GetEntriesByCompetition(competitionId int) ([]businesslogic.AthleteCompetitionEntry, error) {
	entries := make([]businesslogic.AthleteCompetitionEntry, 0)
	var err error
	stmt := `SELECT * FROM GET_ATHLETE_COMPETITION_ENTRIES_BY_COMPETITION($1)`
	rows, err := repo.Database.Query(stmt, competitionId)
	if err != nil {
		return entries, err
	}
	for rows.Next() {
		each := businesslogic.AthleteCompetitionEntry{}
		competitionStatus := 0
		err = rows.Scan(
			&each.ID,
			&each.Competition.ID,
			&each.Competition.Name,
			&competitionStatus,
			&each.Athlete.ID,
			&each.Athlete.FirstName,
			&each.Athlete.LastName,
			&each.Athlete.DateTimeCreated,
			&each.Athlete.UserGenderID,
			&each.IsLead,
			&each.LeadTag,
			&each.OrganizerNote,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		if err != nil {
			return entries, err
		}
		each.Competition.UpdateStatus(competitionStatus)
		entries = append(entries, each)
	}
	return entries, err
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
			entry.Competition.ID,
			entry.Couple.ID,
			entry.CheckedIn,
			entry.DateTimeCheckedIn,
			entry.CreateUserID,
			entry.DateTimeCreated,
			entry.UpdateUserID,
			entry.DateTimeUpdated).
		Suffix(dalutil.SQLSuffixReturningID)

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		scanErr := row.Scan(&entry.ID)
		if scanErr != nil {
			return scanErr
		}
		err = tx.Commit()
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

	partnershipRepo := partnershipdal.PostgresPartnershipRepository{
		Database:   repo.Database,
		SqlBuilder: repo.SQLBuilder,
	}

	for rows.Next() {
		each := businesslogic.PartnershipCompetitionEntry{}
		scanErr := rows.Scan(
			&each.ID,
			&each.Competition.ID,
			&each.Couple.ID,
			&each.CheckedIn,
			&each.DateTimeCheckedIn,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		if scanErr != nil {
			return entries, scanErr
		}

		couples, _ := partnershipRepo.SearchPartnership(businesslogic.SearchPartnershipCriteria{PartnershipID: each.Couple.ID})
		each.Couple = couples[0]
		entries = append(entries, each)
	}

	return entries, err
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
