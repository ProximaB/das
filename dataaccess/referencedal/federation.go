package referencedal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/DancesportSoftware/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
)

const (
	DAS_FEDERATION_TABLE            = "DAS.FEDERATION"
	DAS_FEDERATION_COL_YEAR_FOUNDED = "YEAR_FOUNDED"
)

type PostgresFederationRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresFederationRepository) CreateFederation(federation *businesslogic.Federation) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Insert("").
		Into(DAS_FEDERATION_TABLE).
		Columns(
			common.COL_NAME,
			common.ColumnAbbreviation,
			common.COL_DESCRIPTION,
			DAS_FEDERATION_COL_YEAR_FOUNDED,
			common.COL_COUNTRY_ID,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated,
		).Values(
		federation.Name,
		federation.Abbreviation,
		federation.Description,
		federation.YearFounded,
		federation.YearFounded,
		federation.CreateUserID,
		federation.DateTimeCreated,
		federation.UpdateUserID,
		federation.DateTimeUpdated,
	).Suffix("RETURNING ID")

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&federation.ID)
		tx.Commit()
	}

	return err
}

func (repo PostgresFederationRepository) SearchFederation(criteria businesslogic.SearchFederationCriteria) ([]businesslogic.Federation, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s",
			common.ColumnPrimaryKey,
			common.COL_NAME,
			common.ColumnAbbreviation,
			DAS_FEDERATION_COL_YEAR_FOUNDED,
			common.COL_COUNTRY_ID,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated)).
		From(DAS_FEDERATION_TABLE).OrderBy(common.ColumnPrimaryKey)
	if criteria.CountryID > 0 {
		stmt = stmt.Where(squirrel.Eq{
			common.COL_COUNTRY_ID: criteria.CountryID})
	}
	if len(criteria.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_NAME: criteria.Name})
	}
	if criteria.ID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.ID})
	}

	federations := make([]businesslogic.Federation, 0)
	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return federations, err
	}
	for rows.Next() {
		each := businesslogic.Federation{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Abbreviation,
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

func (repo PostgresFederationRepository) DeleteFederation(federation businesslogic.Federation) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Delete("").From(DAS_FEDERATION_TABLE).Where(squirrel.Eq{common.ColumnPrimaryKey: federation.ID})

	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		tx.Commit()
	}
	return err
}

func (repo PostgresFederationRepository) UpdateFederation(federation businesslogic.Federation) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Update("").Table(DAS_FEDERATION_TABLE)
	if federation.ID > 0 {
		stmt = stmt.Set(common.COL_NAME, federation.Name).
			Set(common.ColumnAbbreviation, federation.Abbreviation).
			Set(common.COL_DESCRIPTION, federation.Description).
			Set(DAS_FEDERATION_COL_YEAR_FOUNDED, federation.YearFounded).
			Set(common.COL_COUNTRY_ID, federation.CountryID).
			Set(common.ColumnUpdateUserID, federation.UpdateUserID).
			Set(common.ColumnDateTimeUpdated, federation.DateTimeUpdated)
		var err error
		if tx, txErr := repo.Database.Begin(); txErr != nil {
			return txErr
		} else {
			_, err = stmt.RunWith(repo.Database).Exec()
			tx.Commit()
		}
		return err
	} else {
		return errors.New("federation is not specified")
	}
}
