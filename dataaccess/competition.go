package dataaccess

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
	DAS_COMPETITION_TABLE              = "DAS.COMPETITION"
	DAS_COMPETITION_COL_STATUS_ID      = "STATUS_ID"
	DAS_COMPETITION_COL_WEBSITE        = "WEBSITE"
	DAS_COMPETITION_COL_ADDRESS        = "ADDRESS"
	DAS_COMPETITION_COL_DATETIME_START = "DATETIME_START"
	DAS_COMPETITION_COL_DATETIME_END   = "DATETIME_END"
	DAS_COMPETITION_COL_CONTACT_NAME   = "CONTACT_NAME"
	DAS_COMPETITION_COL_CONTACT_PHONE  = "CONTACT_PHONE"
	DAS_COMPETITION_COL_CONTACT_EMAIL  = "CONTACT_EMAIL"
	DAS_COMPETITION_COL_ATTENDANCE     = "ATTENDANCE"
)

/*
func UpdateCompetitionAttendance(db *sql.DB, compID int) error {
	var attendance = 0
	getAttendanceClause := SQLBUILDER.Select("COUNT(ID)").From(DAS_COMPETITION_ENTRY_TABLE).Where(sq.Eq{common.COL_COMPETITION_ID: compID})
	getAttendanceClause.RunWith(db).QueryRow().Scan(&attendance)
	updateAttendanceClause := SQLBUILDER.Update(DAS_COMPETITION_TABLE).Set(DAS_COMPETITION_COL_ATTENDANCE, attendance)
	_, err := updateAttendanceClause.RunWith(db).Exec()
	return err
}
*/
type PostgresCompetitionStatusRepository struct {
	Database   *sql.DB
	SqlBuilder sq.StatementBuilderType
}

func (repo PostgresCompetitionStatusRepository) GetCompetitionStatus() ([]businesslogic.CompetitionStatus, error) {
	clause := repo.SqlBuilder.Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s",
		common.PRIMARY_KEY,
		common.COL_NAME,
		common.COL_ABBREVIATION,
		common.COL_DESCRIPTION,
		common.COL_DATETIME_CREATED,
		common.COL_DATETIME_UPDATED)).
		From("DAS.COMPETITION_STATUS").OrderBy(common.PRIMARY_KEY)
	rows, err := clause.RunWith(repo.Database).Query()
	status := make([]businesslogic.CompetitionStatus, 0)
	if err != nil {
		return status, err
	}
	for rows.Next() {
		each := businesslogic.CompetitionStatus{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Abbreviation,
			&each.Description,
			&each.DateTimeCreated,
			&each.DateTimeUpdated,
		)
		status = append(status, each)

	}
	rows.Close()
	return status, err
}

type PostgresCompetitionRepository struct {
	Database   *sql.DB
	SqlBuilder sq.StatementBuilderType
}

func (repo PostgresCompetitionRepository) CreateCompetition(competition businesslogic.Competition) error {
	clause := repo.SqlBuilder.
		Insert("").
		Into(DAS_COMPETITION_TABLE).
		Columns(
			common.COL_FEDERATION_ID,
			common.COL_NAME,
			DAS_COMPETITION_COL_WEBSITE,
			DAS_COMPETITION_COL_STATUS_ID,
			common.COL_COUNTRY_ID,
			common.COL_STATE_ID,
			common.COL_CITY_ID,
			DAS_COMPETITION_COL_ADDRESS,
			DAS_COMPETITION_COL_DATETIME_START,
			DAS_COMPETITION_COL_DATETIME_END,
			DAS_COMPETITION_COL_CONTACT_NAME,
			DAS_COMPETITION_COL_CONTACT_PHONE,
			DAS_COMPETITION_COL_CONTACT_EMAIL,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED).
		Values(competition.FederationID,
			competition.Name,
			competition.Website,
			competition.GetStatus(),
			competition.Country.ID,
			competition.State.ID,
			competition.City.CityID,
			competition.Street,
			competition.StartDateTime,
			competition.EndDateTime,
			competition.ContactName,
			competition.ContactPhone,
			competition.ContactEmail,
			competition.CreateUserID,
			competition.DateTimeCreated,
			competition.UpdateUserID,
			competition.DateTimeUpdated)
	_, err := clause.RunWith(repo.Database).Exec()
	return err
}

func (repo PostgresCompetitionRepository) UpdateCompetition(competition businesslogic.Competition) error {
	stmt := repo.SqlBuilder.Update("").Table(DAS_COMPETITION_TABLE)
	if competition.CompetitionID > 0 {
		stmt = stmt.Set(common.COL_NAME, competition.Name).
			Set(common.COL_WEBSITE, competition.Website).
			Set(DAS_COMPETITION_COL_STATUS_ID, competition.GetStatus()).
			Set(DAS_COMPETITION_COL_DATETIME_START, competition.StartDateTime).
			Set(DAS_COMPETITION_COL_DATETIME_END, competition.EndDateTime).
			Set(DAS_COMPETITION_COL_ADDRESS, competition.Street).
			Set(DAS_COMPETITION_COL_CONTACT_NAME, competition.ContactName).
			Set(DAS_COMPETITION_COL_CONTACT_EMAIL, competition.ContactEmail).
			Set(DAS_COMPETITION_COL_CONTACT_PHONE, competition.ContactPhone).
			Set(common.COL_DATETIME_UPDATED, time.Now())
	}
	stmt = stmt.Where(sq.Eq{common.PRIMARY_KEY: competition.CompetitionID})

	_, err := stmt.RunWith(repo.Database).Exec()
	return err
}

func (repo PostgresCompetitionRepository) DeleteCompetition(competition businesslogic.Competition) error {
	return errors.New("not implemented")
}

func (repo PostgresCompetitionRepository) SearchCompetition(criteria *businesslogic.SearchCompetitionCriteria) ([]businesslogic.Competition, error) {
	stmt := repo.SqlBuilder.Select(fmt.Sprintf(
		"%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
		common.PRIMARY_KEY,
		common.COL_FEDERATION_ID,
		common.COL_NAME,
		common.COL_ADDRESS,
		common.COL_CITY_ID,
		common.COL_STATE_ID,
		common.COL_COUNTRY_ID,
		DAS_COMPETITION_COL_DATETIME_START,
		DAS_COMPETITION_COL_DATETIME_END,
		DAS_COMPETITION_COL_CONTACT_NAME,
		DAS_COMPETITION_COL_CONTACT_PHONE,
		DAS_COMPETITION_COL_CONTACT_EMAIL,
		DAS_COMPETITION_COL_WEBSITE,
		DAS_COMPETITION_COL_STATUS_ID,
		DAS_COMPETITION_COL_ATTENDANCE,
		common.COL_CREATE_USER_ID,
		common.COL_DATETIME_CREATED,
		common.COL_UPDATE_USER_ID,
		common.COL_DATETIME_UPDATED)).From(DAS_COMPETITION_TABLE).OrderBy(DAS_COMPETITION_COL_DATETIME_START)

	if criteria.ID > 0 {
		stmt = stmt.Where(sq.Eq{common.PRIMARY_KEY: criteria.ID})
	}
	if len(criteria.Name) > 0 {
		stmt = stmt.Where(sq.Eq{common.COL_NAME: criteria.Name})
	}
	if criteria.FederationID > 0 {
		stmt = stmt.Where(sq.Eq{common.COL_FEDERATION_ID: criteria.FederationID})
	}
	if criteria.StateID > 0 {
		stmt = stmt.Where(sq.Eq{common.COL_STATE_ID: criteria.StateID})
	}

	if criteria.CountryID > 0 {
		stmt = stmt.Where(sq.Eq{common.COL_COUNTRY_ID: criteria.CountryID})
	}
	if criteria.StartDateTime.After(time.Now()) {
		stmt = stmt.Where(sq.Eq{DAS_COMPETITION_COL_DATETIME_START: criteria.StartDateTime})
	}
	if criteria.OrganizerID > 0 {
		stmt = stmt.Where(sq.Eq{common.COL_CREATE_USER_ID: criteria.OrganizerID})
	}
	if criteria.StatusID > 0 {
		stmt = stmt.Where(sq.Eq{DAS_COMPETITION_COL_STATUS_ID: criteria.StatusID})
	}

	rows, err := stmt.RunWith(repo.Database).Query()
	comps := make([]businesslogic.Competition, 0)

	for rows.Next() {
		each := businesslogic.Competition{}
		status := 0
		rows.Scan(
			&each.CompetitionID,
			&each.FederationID,
			&each.Name,
			&each.Street,
			&each.City.CityID,
			&each.State.ID,
			&each.Country.ID,
			&each.StartDateTime,
			&each.EndDateTime,
			&each.ContactName,
			&each.ContactPhone,
			&each.ContactEmail,
			&each.Website,
			&status,
			&each.Attendance,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		each.UpdateStatus(status)
		comps = append(comps, each)
	}
	return comps, err
}

/*
func GetEventUniqueFederations(db *sql.DB, compID int) ([]reference.Federation, error) {
	clause := `SELECT DISTINCT ECB.FEDERATION_ID, F.NAME, F.ABBREVIATION FROM DAS.EVENT_COMPETITIVE_BALLROOM ECB JOIN DAS.FEDERATION F
				ON ECB.FEDERATION_ID = F.ID JOIN DAS.EVENT E ON E.ID = ECB.EVENT_ID WHERE E.COMPETITION_ID = $1`

	rows, err := db.Query(clause, compID)
	federations := make([]reference.Federation, 0)
	if err != nil {
		rows.Close()
		return federations, err
	}

	for rows.Next() {
		each := reference.Federation{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Abbreviation,
		)
		federations = append(federations, each)
	}
	rows.Close()
	return federations, err
}

func GetEventUniqueDivisions(db *sql.DB, compID int) ([]reference.Division, error) {
	clause := `SELECT DISTINCT ECB.DIVISION_ID, D.NAME, D.FEDERATION_ID FROM DAS.EVENT_COMPETITIVE_BALLROOM ECB JOIN DAS.DIVISION D
				ON ECB.DIVISION_ID = D.ID JOIN DAS.EVENT E ON E.ID = ECB.EVENT_ID WHERE E.COMPETITION_ID = $1`
	rows, err := db.Query(clause, compID)
	divisions := make([]reference.Division, 0)
	if err != nil {
		rows.Close()
		return divisions, err
	}

	for rows.Next() {
		each := reference.Division{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.FederationID,
		)
		divisions = append(divisions, each)
	}
	rows.Close()
	return divisions, err
}

func GetEventUniqueAges(db *sql.DB, compID int) ([]reference.Age, error) {
	clause := `SELECT DISTINCT ECB.AGE_ID, A.NAME, A.DIVISION_ID, A.MINIMUM_AGE, A.MAXIMUM_AGE FROM DAS.EVENT_COMPETITIVE_BALLROOM ECB
				JOIN DAS.AGE A ON ECB.AGE_ID = A.ID JOIN DAS.EVENT E ON E.ID = ECB.EVENT_ID WHERE E.COMPETITION_ID = $1`
	rows, err := db.Query(clause, compID)
	ages := make([]reference.Age, 0)
	if err != nil {
		rows.Close()
		return ages, err
	}

	for rows.Next() {
		each := reference.Age{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.DivisionID,
			&each.AgeMinimum,
			&each.AgeMaximum,
		)
		ages = append(ages, each)
	}
	rows.Close()
	return ages, err
}

func GetEventUniqueProficiencies(db *sql.DB, compID int) ([]reference.Proficiency, error) {
	clause := `SELECT DISTINCT ECB.PROFICIENCY_ID, P.NAME, P.DIVISION_ID FROM DAS.EVENT_COMPETITIVE_BALLROOM ECB
				JOIN DAS.PROFICIENCY P ON ECB.PROFICIENCY_ID = P.ID JOIN DAS.EVENT E ON E.ID = ECB.EVENT_ID WHERE E.COMPETITION_ID = $1`
	rows, err := db.Query(clause, compID)
	proficiencies := make([]reference.Proficiency, 0)
	if err != nil {
		rows.Close()
		return proficiencies, err
	}

	for rows.Next() {
		each := reference.Proficiency{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.DivisionID,
		)
		proficiencies = append(proficiencies, each)
	}
	rows.Close()
	return proficiencies, err
}

func GetEventUniqueStyles(db *sql.DB, compID int) ([]reference.Style, error) {
	clause := `SELECT DISTINCT ECB.STYLE_ID, S.NAME FROM DAS.EVENT_COMPETITIVE_BALLROOM ECB JOIN DAS.STYLE S ON ECB.STYLE_ID = S.ID
				JOIN DAS.EVENT E ON E.ID = ECB.EVENT_ID WHERE E.COMPETITION_ID = $1`

	rows, err := db.Query(clause, compID)
	styles := make([]reference.Style, 0)
	if err != nil {
		rows.Close()
		return styles, err
	}

	for rows.Next() {
		each := reference.Style{}
		rows.Scan(
			&each.ID,
			&each.Name,
		)
		styles = append(styles, each)
	}
	rows.Close()
	return styles, err
}
*/
