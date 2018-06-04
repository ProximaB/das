package reference

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/Masterminds/squirrel"
)

const (
	DAS_CITY_TABLE = "DAS.CITY"
)

type PostgresCityRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresCityRepository) CreateCity(city *reference.City) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresCityRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Insert("").
		Into(DAS_CITY_TABLE).
		Columns(common.COL_NAME,
			common.COL_STATE_ID,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED).
		Values(
			city.Name,
			city.StateID,
			city.CreateUserID,
			city.DateTimeCreated,
			city.UpdateUserID,
			city.DateTimeUpdated).Suffix("RETURNING ID")

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&city.CityID)
		tx.Commit()
	}
	return err
}

func (repo PostgresCityRepository) DeleteCity(city reference.City) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresCityRepository is not specified")
	}
	stmt := repo.SqlBuilder.Delete("").From(DAS_CITY_TABLE)
	if city.CityID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.PRIMARY_KEY: city.CityID})
	}
	if len(city.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_NAME: city.Name})
	} else {
		return errors.New("cannot identify City")
	}

	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		tx.Commit()
	}

	return err
}

func (repo PostgresCityRepository) UpdateCity(city reference.City) error {
	if repo.Database == nil {
		return errors.New("data source of PostgresCityRepository is not specified")
	}
	stmt := repo.SqlBuilder.Update("").Table(DAS_CITY_TABLE).
		SetMap(squirrel.Eq{common.COL_NAME: city.Name, common.COL_STATE_ID: city.StateID}).
		SetMap(squirrel.Eq{common.COL_DATETIME_UPDATED: city.DateTimeUpdated}).Where(squirrel.Eq{common.PRIMARY_KEY: city.CityID})

	if city.UpdateUserID != nil {
		stmt = stmt.SetMap(squirrel.Eq{common.COL_UPDATE_USER_ID: city.UpdateUserID})
	}

	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		tx.Commit()
	}
	return err

}

func (repo PostgresCityRepository) SearchCity(criteria reference.SearchCityCriteria) ([]reference.City, error) {
	if repo.Database == nil {
		return nil, errors.New("data source of PostgresCityRepository is not specified")
	}
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s",
			common.PRIMARY_KEY,
			common.COL_NAME,
			common.COL_STATE_ID,
			common.COL_CREATE_USER_ID,
			common.COL_DATETIME_CREATED,
			common.COL_UPDATE_USER_ID,
			common.COL_DATETIME_UPDATED)).
		From(DAS_CITY_TABLE).OrderBy(common.PRIMARY_KEY)
	if len(criteria.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_NAME: criteria.Name})
	}
	if criteria.StateID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.COL_STATE_ID: criteria.StateID})
	}
	if criteria.CityID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.PRIMARY_KEY: criteria.CityID})
	}

	rows, err := stmt.RunWith(repo.Database).Query()
	cities := make([]reference.City, 0)
	if err != nil {
		return cities, err
	}
	for rows.Next() {
		each := reference.City{}
		scanErr := rows.Scan(
			&each.CityID,
			&each.Name,
			&each.StateID,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		if scanErr != nil {
			rows.Close()
			return cities, scanErr
		}
		cities = append(cities, each)
	}
	rows.Close()
	return cities, nil
}
