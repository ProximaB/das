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

package eventdal

import (
	"database/sql"
	"errors"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
)

type PostgresEventMetaRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresEventMetaRepository) GetEventUniqueFederations(competition businesslogic.Competition) ([]businesslogic.Federation, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}

	clause := `SELECT DISTINCT E.FEDERATION_ID, F.NAME, F.ABBREVIATION, F.DESCRIPTION, F.YEAR_FOUNDED, F.COUNTRY_ID,
				F.CREATE_USER_ID, F.DATETIME_CREATED, F.UPDATE_USER_ID, F.DATETIME_UPDATED FROM DAS.EVENT E JOIN DAS.FEDERATION F
				ON E.FEDERATION_ID = F.ID WHERE E.COMPETITION_ID = $1`

	rows, err := repo.Database.Query(clause, competition.ID)
	federations := make([]businesslogic.Federation, 0)
	if err != nil {
		rows.Close()
		return federations, err
	}

	for rows.Next() {
		each := businesslogic.Federation{}
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

func (repo PostgresEventMetaRepository) GetEventUniqueDivisions(competition businesslogic.Competition) ([]businesslogic.Division, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	clause := `SELECT DISTINCT E.DIVISION_ID, D.NAME, D.ABBREVIATION, D.DESCRIPTION, D.NOTE, D.FEDERATION_ID, 
				D.DATETIME_CREATED, D.DATETIME_UPDATED 
				FROM DAS.EVENT E JOIN DAS.DIVISION D
				ON E.DIVISION_ID = D.ID WHERE E.COMPETITION_ID = $1`
	rows, err := repo.Database.Query(clause, competition.ID)
	divisions := make([]businesslogic.Division, 0)
	if err != nil {
		rows.Close()
		return divisions, err
	}

	for rows.Next() {
		each := businesslogic.Division{}
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

func (repo PostgresEventMetaRepository) GetEventUniqueAges(competition businesslogic.Competition) ([]businesslogic.Age, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	clause := `SELECT DISTINCT E.AGE_ID, A.NAME, A.DESCRIPTION, A.DIVISION_ID, A. ENFORCED, A.MINIMUM_AGE, 
			A.MAXIMUM_AGE, A.CREATE_USER_ID, A.DATETIME_CREATED, A.UPDATE_USER_ID, A.DATETIME_UPDATED
			FROM DAS.EVENT E
				JOIN DAS.AGE A ON E.AGE_ID = A.ID WHERE E.COMPETITION_ID = $1`
	rows, err := repo.Database.Query(clause, competition.ID)
	ages := make([]businesslogic.Age, 0)
	if err != nil {
		rows.Close()
		return ages, err
	}

	for rows.Next() {
		each := businesslogic.Age{}
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

func (repo PostgresEventMetaRepository) GetEventUniqueProficiencies(competition businesslogic.Competition) ([]businesslogic.Proficiency, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	clause := `SELECT DISTINCT E.PROFICIENCY_ID, P.NAME, P.DIVISION_ID P.DESCRIPTION, P.CREATE_USER_ID, P.DATETIME_CREATED,
			P.UPDATE_USER_ID, P.DATETIME_UPDATED FROM DAS.EVENT E JOIN DAS.PROFICIENCY P ON E.PROFICIENCY_ID = P.ID 
			WHERE E.COMPETITION_ID = $1`
	rows, err := repo.Database.Query(clause, competition.ID)
	proficiencies := make([]businesslogic.Proficiency, 0)
	if err != nil {
		rows.Close()
		return proficiencies, err
	}

	for rows.Next() {
		each := businesslogic.Proficiency{}
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

func (repo PostgresEventMetaRepository) GetEventUniqueStyles(competition businesslogic.Competition) ([]businesslogic.Style, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	clause := `SELECT DISTINCT E.STYLE_ID, S.NAME, S.DESCRIPTION, S.CREATE_USER_ID, S.DATETIME_CREATED,
				S.UPDATE_USER_ID, S.DATETIME_UPDATED FROM DAS.EVENT E JOIN DAS.STYLE S ON E.STYLE_ID = S.ID
				WHERE E.COMPETITION_ID = $1`

	rows, err := repo.Database.Query(clause, competition)
	styles := make([]businesslogic.Style, 0)
	if err != nil {
		rows.Close()
		return styles, err
	}

	for rows.Next() {
		each := businesslogic.Style{}
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
