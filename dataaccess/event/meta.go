// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package event

import (
	"database/sql"
	"errors"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/Masterminds/squirrel"
)

type PostgresEventMetaRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresEventMetaRepository) GetEventUniqueFederations(competition businesslogic.Competition) ([]referencebll.Federation, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresEventRepository is not specified")
	}

	clause := `SELECT DISTINCT E.FEDERATION_ID, F.NAME, F.ABBREVIATION, F.DESCRIPTION, F.YEAR_FOUNDED, F.COUNTRY_ID,
				F.CREATE_USER_ID, F.DATETIME_CREATED, F.UPDATE_USER_ID, F.DATETIME_UPDATED FROM DAS.EVENT E JOIN DAS.FEDERATION F
				ON E.FEDERATION_ID = F.ID WHERE E.COMPETITION_ID = $1`

	rows, err := repo.Database.Query(clause, competition.ID)
	federations := make([]referencebll.Federation, 0)
	if err != nil {
		rows.Close()
		return federations, err
	}

	for rows.Next() {
		each := referencebll.Federation{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Abbreviation,
			&each.Description,
			&each.YearFounded,
			&each.CountryID,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		federations = append(federations, each)
	}
	rows.Close()
	return federations, err
}

func (repo PostgresEventMetaRepository) GetEventUniqueDivisions(competition businesslogic.Competition) ([]referencebll.Division, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresEventRepository is not specified")
	}
	clause := `SELECT DISTINCT E.DIVISION_ID, D.NAME, D.ABBREVIATION, D.DESCRIPTION, D.NOTE, D.FEDERATION_ID, 
				D.DATETIME_CREATED, D.DATETIME_UPDATED 
				FROM DAS.EVENT E JOIN DAS.DIVISION D
				ON E.DIVISION_ID = D.ID WHERE E.COMPETITION_ID = $1`
	rows, err := repo.Database.Query(clause, competition.ID)
	divisions := make([]referencebll.Division, 0)
	if err != nil {
		rows.Close()
		return divisions, err
	}

	for rows.Next() {
		each := referencebll.Division{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Abbreviation,
			&each.Description,
			&each.Note,
			&each.FederationID,
			&each.DateTimeCreated,
			&each.DateTimeUpdated,
		)
		divisions = append(divisions, each)
	}
	rows.Close()
	return divisions, err
}

func (repo PostgresEventMetaRepository) GetEventUniqueAges(competition businesslogic.Competition) ([]referencebll.Age, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresEventRepository is not specified")
	}
	clause := `SELECT DISTINCT E.AGE_ID, A.NAME, A.DESCRIPTION, A.DIVISION_ID, A. ENFORCED, A.MINIMUM_AGE, 
			A.MAXIMUM_AGE, A.CREATE_USER_ID, A.DATETIME_CREATED, A.UPDATE_USER_ID, A.DATETIME_UPDATED
			FROM DAS.EVENT E
				JOIN DAS.AGE A ON E.AGE_ID = A.ID WHERE E.COMPETITION_ID = $1`
	rows, err := repo.Database.Query(clause, competition.ID)
	ages := make([]referencebll.Age, 0)
	if err != nil {
		rows.Close()
		return ages, err
	}

	for rows.Next() {
		each := referencebll.Age{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Description,
			&each.DivisionID,
			&each.Enforced,
			&each.AgeMinimum,
			&each.AgeMaximum,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		ages = append(ages, each)
	}
	rows.Close()
	return ages, err
}

func (repo PostgresEventMetaRepository) GetEventUniqueProficiencies(competition businesslogic.Competition) ([]referencebll.Proficiency, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresEventRepository is not specified")
	}
	clause := `SELECT DISTINCT E.PROFICIENCY_ID, P.NAME, P.DIVISION_ID P.DESCRIPTION, P.CREATE_USER_ID, P.DATETIME_CREATED,
			P.UPDATE_USER_ID, P.DATETIME_UPDATED FROM DAS.EVENT E JOIN DAS.PROFICIENCY P ON E.PROFICIENCY_ID = P.ID 
			WHERE E.COMPETITION_ID = $1`
	rows, err := repo.Database.Query(clause, competition.ID)
	proficiencies := make([]referencebll.Proficiency, 0)
	if err != nil {
		rows.Close()
		return proficiencies, err
	}

	for rows.Next() {
		each := referencebll.Proficiency{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.DivisionID,
			&each.Description,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTImeUpdated,
		)
		proficiencies = append(proficiencies, each)
	}
	rows.Close()
	return proficiencies, err
}

func (repo PostgresEventMetaRepository) GetEventUniqueStyles(competition businesslogic.Competition) ([]referencebll.Style, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresEventRepository is not specified")
	}
	clause := `SELECT DISTINCT E.STYLE_ID, S.NAME, S.DESCRIPTION, S.CREATE_USER_ID, S.DATETIME_CREATED,
				S.UPDATE_USER_ID, S.DATETIME_UPDATED FROM DAS.EVENT E JOIN DAS.STYLE S ON E.STYLE_ID = S.ID
				WHERE E.COMPETITION_ID = $1`

	rows, err := repo.Database.Query(clause, competition)
	styles := make([]referencebll.Style, 0)
	if err != nil {
		rows.Close()
		return styles, err
	}

	for rows.Next() {
		each := referencebll.Style{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Description,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		styles = append(styles, each)
	}
	rows.Close()
	return styles, err
}
