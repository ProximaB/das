package referencedal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/DancesportSoftware/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
	"time"
)

const (
	DAS_COUNTRY_TABLE = "DAS.COUNTRY"
)

// PostgresCountryRepository implements the ICountryRepository with a Postgres database
type PostgresCountryRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

// CreateCountry inserts a Country object into a Postgres database
func (repo PostgresCountryRepository) CreateCountry(country *businesslogic.Country) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.
		Insert("").
		Into(DAS_COUNTRY_TABLE).
		Columns(common.COL_NAME,
			common.ColumnAbbreviation,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated).
		Values(
			country.Name,
			country.Abbreviation,
			country.CreateUserID,
			country.DateTimeCreated,
			country.CreateUserID,
			country.DateTimeUpdated,
		).Suffix(
		"RETURNING ID",
	)
	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&country.ID)
		tx.Commit()
	}
	return err
}

// DeleteCountry deletes a Country object from a Postgres database
func (repo PostgresCountryRepository) DeleteCountry(country businesslogic.Country) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Delete("").From(DAS_COUNTRY_TABLE).
		Where(squirrel.Eq{common.ColumnPrimaryKey: country.ID})
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		tx.Commit()
	}
	return err
}

// UpdateCountry updates a Country object in a Postgres database
func (repo PostgresCountryRepository) UpdateCountry(country businesslogic.Country) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Update("").Table(DAS_COUNTRY_TABLE)
	if country.ID > 0 {
		stmt = stmt.Set(common.COL_NAME, country.Name).
			Set(common.ColumnAbbreviation, country.Abbreviation).
			Set(common.ColumnDateTimeUpdated, time.Now()).
			Where(squirrel.Eq{common.ColumnPrimaryKey: country.ID})
		if country.UpdateUserID != nil {
			stmt = stmt.Set(common.ColumnUpdateUserID, country.UpdateUserID)
		}

		var err error
		if tx, txErr := repo.Database.Begin(); txErr != nil {
			return txErr
		} else {
			_, err = stmt.RunWith(repo.Database).Exec()
			tx.Commit()
		}

		return err
	} else {
		return errors.New("country is not specified")
	}
}

// SearchCountry searches the Country object in a Postgres database with the provided criteria
func (repo PostgresCountryRepository) SearchCountry(criteria businesslogic.SearchCountryCriteria) ([]businesslogic.Country, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s",
			common.ColumnPrimaryKey,
			common.COL_NAME,
			common.ColumnAbbreviation,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated)).
		From(DAS_COUNTRY_TABLE).
		OrderBy(common.ColumnPrimaryKey)
	if criteria.CountryID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.CountryID})
	}
	if len(criteria.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_NAME: criteria.Name})
	}
	if len(criteria.Abbreviation) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnAbbreviation: criteria.Abbreviation})
	}

	rows, err := stmt.RunWith(repo.Database).Query()
	countries := make([]businesslogic.Country, 0)
	if err != nil {
		return countries, err
	}
	for rows.Next() {
		each := businesslogic.Country{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Abbreviation,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		countries = append(countries, each)
	}
	rows.Close()
	/*row := repo.Database.QueryRow(fmt.Sprintf("SELECT * FROM GET_COUNTRY_BY_ID (%d)", criteria.CountryID))
	countries := make([]reference.Country, 0)
	each := reference.Country{}
	err := row.Scan(
		&each.ID,
		&each.Name,
		&each.Abbreviation,
		&each.CreateUserID,
		&each.DateTimeCreated,
		&each.UpdateUserID,
		&each.DateTimeUpdated,
	)
	countries = append(countries, each)*/
	return countries, err
}
