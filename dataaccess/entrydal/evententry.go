package entrydal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/dataaccess/accountdal"
	"github.com/DancesportSoftware/das/dataaccess/competition"
	"github.com/DancesportSoftware/das/dataaccess/eventdal"
	"github.com/DancesportSoftware/das/dataaccess/partnershipdal"
	"log"

	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/DancesportSoftware/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
)

const (
	dasAthleteEventEntryTable     = "DAS.EVENT_ENTRY_ATHLETE"
	dasPartnershipEventEntryTable = "DAS.EVENT_ENTRY_PARTNERSHIP"
	columnCheckinIndicator        = "CHECKIN_IND"
	columnCheckinDateTime         = "CHECKIN_DATETIME"
)

type PostgresAthleteEventEntryRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

func (repo PostgresAthleteEventEntryRepository) CreateAthleteEventEntry(entry *businesslogic.AthleteEventEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.Insert("").
		Into(dasAthleteEventEntryTable).
		Columns(
			common.COL_ATHLETE_ID,
			common.COL_COMPETITION_ID,
			common.COL_EVENT_ID,
			columnCheckinIndicator,
			columnCheckinDateTime,
			common.COL_PLACEMENT,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated,
		).
		Values(
			entry.Athlete.ID,
			entry.Competition.ID,
			entry.Event.ID,
			entry.CheckedIn,
			entry.DateTimeCheckedIn,
			entry.Placement,
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
		txErr := tx.Commit()
		if txErr != nil {
			return txErr
		}
	}
	return err
}

func (repo PostgresAthleteEventEntryRepository) DeleteAthleteEventEntry(entry businesslogic.AthleteEventEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.Delete("").From(dasAthleteEventEntryTable)
	if entry.ID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: entry.ID})
	} else {
		return errors.New(fmt.Sprintf("cannot find this Athlete Event Entry with ID: %v", entry.ID))
	}

	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		err = tx.Commit()
	}
	return err
}

func (repo PostgresAthleteEventEntryRepository) SearchAthleteEventEntry(criteria businesslogic.SearchAthleteEventEntryCriteria) ([]businesslogic.AthleteEventEntry, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
		common.ColumnPrimaryKey,
		common.COL_ATHLETE_ID,
		common.COL_COMPETITION_ID,
		common.COL_EVENT_ID,
		columnCheckinIndicator,
		columnCheckinDateTime,
		common.COL_PLACEMENT,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated,
	)).From(dasAthleteEventEntryTable)

	if criteria.ID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.ID})
	}
	if criteria.CompetitionID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_COMPETITION_ID: criteria.CompetitionID})
	}
	if criteria.EventID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_EVENT_ID: criteria.EventID})
	}
	if criteria.AthleteID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_ATHLETE_ID: criteria.AthleteID})
	}

	entries := make([]businesslogic.AthleteEventEntry, 0)
	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		each := businesslogic.AthleteEventEntry{}
		scanErr := rows.Scan(
			&each.ID,
			&each.Athlete.ID,
			&each.Competition.ID,
			&each.Event.ID,
			&each.CheckedIn,
			&each.DateTimeCheckedIn,
			&each.Placement,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated)
		if scanErr != nil {
			log.Printf("[error] scanning Athlete Event Entry: %v", scanErr)
			return entries, scanErr
		}
		entries = append(entries, each)
	}

	accountRepo := accountdal.PostgresAccountRepository{
		Database:   repo.Database,
		SQLBuilder: repo.SQLBuilder,
	}
	competitionRepo := competition.PostgresCompetitionRepository{
		Database:   repo.Database,
		SqlBuilder: repo.SQLBuilder,
	}
	eventRepo := eventdal.PostgresEventRepository{
		Database:   repo.Database,
		SQLBuilder: repo.SQLBuilder,
	}

	for _, each := range entries {
		athletes, _ := accountRepo.SearchAccount(businesslogic.SearchAccountCriteria{ID: each.Athlete.ID})
		each.Athlete = athletes[0]

		competitions, _ := competitionRepo.SearchCompetition(businesslogic.SearchCompetitionCriteria{ID: each.Competition.ID})
		each.Competition = competitions[0]

		events, _ := eventRepo.SearchEvent(businesslogic.SearchEventCriteria{EventID: each.Event.ID})
		each.Event = events[0]
	}
	return entries, rows.Close()
}

func (repo PostgresAthleteEventEntryRepository) UpdateAthleteEventEntry(entry businesslogic.AthleteEventEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// PostgresPartnershipEventEntryRepository is a Postgres-based implementation of IPartnershipEventEntryRepository
type PostgresPartnershipEventEntryRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

const (
	leadTag = "LEADTAG"
)

// CreatePartnershipEventEntry creates a Partnership Event Entry in a Postgres database
func (repo PostgresPartnershipEventEntryRepository) CreatePartnershipEventEntry(entry *businesslogic.PartnershipEventEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SQLBuilder.Insert("").Into(dasPartnershipEventEntryTable).Columns(
		common.COL_EVENT_ID,
		common.COL_PARTNERSHIP_ID,
		leadTag,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated,
	).Values(
		entry.Event.ID,
		entry.Couple.ID,
		entry.CompetitorTag, // TODO: this needs to be fixed later
		entry.CreateUserID,
		entry.DateTimeCreated,
		entry.UpdateUserID,
		entry.DateTimeUpdated,
	).Suffix(dalutil.SQLSuffixReturningID)
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

// DeletePartnershipEventEntry deletes a Partnership Event Entry from a Postgres database
func (repo PostgresPartnershipEventEntryRepository) DeletePartnershipEventEntry(entry businesslogic.PartnershipEventEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	if entry.ID == 0 {
		return errors.New("ID of Partnership Event Entry is required")
	}
	clause := repo.SQLBuilder.Delete("").
		From(dasPartnershipEventEntryTable).
		Where(squirrel.Eq{common.COL_EVENT_ID: entry.Event.ID}).
		Where(squirrel.Eq{common.COL_PARTNERSHIP_ID: entry.Couple.ID})
	_, err := clause.RunWith(repo.Database).Exec()
	return err
}

// UpdatePartnershipEventEntry makes changes to a Partnership Event Entry in a Postgres database
func (repo PostgresPartnershipEventEntryRepository) UpdatePartnershipEventEntry(entry businesslogic.PartnershipEventEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// SearchPartnershipEventEntry returns CompetitiveBallroomEventEntry, which is supposed to be used by competitor only
func (repo PostgresPartnershipEventEntryRepository) SearchPartnershipEventEntry(criteria businesslogic.SearchPartnershipEventEntryCriteria) ([]businesslogic.PartnershipEventEntry, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	clause := repo.SQLBuilder.Select(
		fmt.Sprintf("%s.%s, %s.%s, %s.%s, %s.%s, %s.%s, %s.%s, %s.%s, %s.%s, %s.%s",
			dasPartnershipEventEntryTable, common.ColumnPrimaryKey,
			dasPartnershipEventEntryTable, common.COL_EVENT_ID,
			dasPartnershipEventEntryTable, common.COL_PARTNERSHIP_ID,
			dasPartnershipEventEntryTable, dasAthleteCompetitionEntryColumnLeadTag,
			dasPartnershipEventEntryTable, dasCompetitionEntryColCheckinInd,
			dasPartnershipEventEntryTable, common.ColumnCreateUserID,
			dasPartnershipEventEntryTable, common.ColumnDateTimeCreated,
			dasPartnershipEventEntryTable, common.ColumnUpdateUserID,
			dasPartnershipEventEntryTable, common.ColumnDateTimeUpdated)).
		From(dasPartnershipEventEntryTable)

	if criteria.PartnershipID > 0 {
		clause = clause.Where(squirrel.Eq{common.COL_PARTNERSHIP_ID: criteria.PartnershipID})
	}
	if criteria.EventID > 0 {
		clause = clause.Where(squirrel.Eq{"DAS.EVENT_ENTRY_PARTNERSHIP.EVENT_ID": criteria.EventID})
	}

	entries := make([]businesslogic.PartnershipEventEntry, 0)
	rows, err := clause.RunWith(repo.Database).Query()

	if err != nil {
		return entries, err
	}

	for rows.Next() {
		each := businesslogic.PartnershipEventEntry{
			Competition: businesslogic.Competition{},
			Couple:      businesslogic.Partnership{},
		}
		scanErr := rows.Scan(
			&each.ID,
			&each.Event.ID,
			&each.Couple.ID,
			&each.CompetitorTag,
			&each.CheckedIn,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		if scanErr != nil {
			log.Printf("[error] scanning Partnership Event Entry: %v", scanErr)
			return entries, scanErr
		}
		entries = append(entries, each)
	}
	closeRowErr := rows.Close()

	partnershipRepo := partnershipdal.PostgresPartnershipRepository{
		repo.Database,
		repo.SQLBuilder,
	}
	eventRepo := eventdal.PostgresEventRepository{
		repo.Database,
		repo.SQLBuilder,
	}
	competitionRepo := competition.PostgresCompetitionRepository{
		repo.Database,
		repo.SQLBuilder,
	}
	for i := 0; i < len(entries); i++ {
		partnerships, _ := partnershipRepo.SearchPartnership(businesslogic.SearchPartnershipCriteria{PartnershipID: entries[i].Couple.ID})
		entries[i].Couple = partnerships[0]
		events, _ := eventRepo.SearchEvent(businesslogic.SearchEventCriteria{EventID: entries[i].Event.ID})
		entries[i].Event = events[0]
		competitions, _ := competitionRepo.SearchCompetition(businesslogic.SearchCompetitionCriteria{ID: entries[i].Event.CompetitionID})
		entries[i].Competition = competitions[0]
	}
	return entries, closeRowErr
}

// PostgresAdjudicatorEventEntryRepository implements IAdjudicatorEventEntryRepository with a Postgres database
type PostgresAdjudicatorEventEntryRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

// CreateAdjudicatorEventEntry creates an Adjudicator Event Entry in a Postgres database
func (repo PostgresAdjudicatorEventEntryRepository) CreateAdjudicatorEventEntry(entry *businesslogic.AdjudicatorEventEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// DeleteAdjudicatorEventEntry deletes an Adjudicator Event Entry from a Postgres database
func (repo PostgresAdjudicatorEventEntryRepository) DeleteAdjudicatorEventEntry(entry businesslogic.AdjudicatorEventEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// SearchAdjudicatorEventEntry searches Adjudicator Event Entries in a Postgres database
func (repo PostgresAdjudicatorEventEntryRepository) SearchAdjudicatorEventEntry(criteria businesslogic.SearchAdjudicatorEventEntryCriteria) ([]businesslogic.AdjudicatorEventEntry, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return nil, errors.New("not implemented")
}

// UpdateAdjudicatorEventEntry updates an Adjudicator Event Entry in a Postgres database
func (repo PostgresAdjudicatorEventEntryRepository) UpdateAdjudicatorEventEntry(entry businesslogic.AdjudicatorEventEntry) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	return errors.New("not implemented")
}

// Returns CompetitiveBallroomEventEntryPublicView, which contains minimal information of the entry and is used by
// public only
/*
func GetCompetitiveBallroomEventEntrylist(criteria *businesslogic.SearchEventEntryCriteria) ([]businesslogic.EventEntryPublicView, error) {
	entries := make([]businesslogic.EventEntryPublicView, 0)

	clause := repo.SQLBuilder.Select(`ECBE.ID, ECB.ID, E.ID, C.ID, P.ID, P.LEAD, P.FOLLOW,
					AL.FIRST_NAME, AL.LAST_NAME,
					AF.FIRST_NAME, AF.LAST_NAME,
					RC.NAME, RST.NAME, RSC.NAME, RSO.NAME
			`).
		From(dasEventCompetitiveBallroomEntryTable).
		Where(sq.Eq{"E.COMPETITION_ID": criteria.ID})

	if criteria.Federation > 0 {
		clause = clause.Where(sq.Eq{"ECB.FEDERATION_ID": criteria.Federation})
	}
	if criteria.Division > 0 {
		clause = clause.Where(sq.Eq{"ECB.DIVISION_ID": criteria.Division})
	}
	if criteria.Age > 0 {
		clause = clause.Where(sq.Eq{"ECB.AGE_ID": criteria.Age})
	}
	if criteria.Proficiency > 0 {
		clause = clause.Where(sq.Eq{"ECB.PROFICIENCY_ID": criteria.Proficiency})
	}

	rows, err := clause.RunWith(repo.Database).Query()

	if err != nil {
		rows.Close()
		return entries, err
	}
	for rows.Next() {

	}
	rows.Close()
	return entries, err
}
*/
