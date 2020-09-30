package referencedal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/dataaccess/common"
	"github.com/ProximaB/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
)

const (
	DAS_SCHOOL_TABLE = "DAS.SCHOOL"
)

type PostgresSchoolRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

func (repo PostgresSchoolRepository) CreateSchool(school *businesslogic.School) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Insert("").Into(DAS_SCHOOL_TABLE).Columns(
		common.COL_NAME,
		common.COL_CITY_ID,
		common.ColumnCreateUserID,
		common.ColumnDateTimeCreated,
		common.ColumnUpdateUserID,
		common.ColumnDateTimeUpdated,
	).Values(
		school.Name,
		school.CityID,
		school.CreateUserID,
		school.DateTimeCreated,
		school.UpdateUserID,
		school.DateTimeUpdated,
	).Suffix(
		"RETURNING ID",
	)

	clause, args, err := stmt.ToSql()
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		row := repo.Database.QueryRow(clause, args...)
		row.Scan(&school.ID)
		tx.Commit()
	}
	return err
}

func (repo PostgresSchoolRepository) UpdateSchool(school businesslogic.School) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.Update("").Table(DAS_SCHOOL_TABLE)
	if school.ID > 0 {
		stmt = stmt.Set(common.COL_NAME, school.Name).
			Set(common.COL_CITY_ID, school.CityID).
			Set(common.ColumnUpdateUserID, school.UpdateUserID).
			Set(common.ColumnDateTimeUpdated, school.DateTimeUpdated)
		var err error
		if tx, txErr := repo.Database.Begin(); txErr != nil {
			return txErr
		} else {
			_, err = stmt.RunWith(repo.Database).Exec()
			if err = tx.Commit(); err != nil {
				tx.Rollback()
			}
		}
		return err
	} else {
		return errors.New("school is not specified")
	}
}

func (repo PostgresSchoolRepository) DeleteSchool(school businesslogic.School) error {
	if repo.Database == nil {
		return errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.
		Delete("").
		From(DAS_SCHOOL_TABLE).
		Where(squirrel.Eq{common.ColumnPrimaryKey: school.ID})
	var err error
	if tx, txErr := repo.Database.Begin(); txErr != nil {
		return txErr
	} else {
		_, err = stmt.RunWith(repo.Database).Exec()
		if err = tx.Commit(); err != nil {
			tx.Rollback()
		}
	}
	return err
}

func (repo PostgresSchoolRepository) SearchSchool(criteria businesslogic.SearchSchoolCriteria) ([]businesslogic.School, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	stmt := repo.SqlBuilder.
		Select(fmt.Sprintf(
			`%s,%s, %s,%s, %s, %s, %s`,
			common.ColumnPrimaryKey,
			common.COL_NAME,
			common.COL_CITY_ID,
			common.ColumnCreateUserID,
			common.ColumnDateTimeCreated,
			common.ColumnUpdateUserID,
			common.ColumnDateTimeUpdated)).
		From(DAS_SCHOOL_TABLE).
		OrderBy(`DAS.SCHOOL.ID`)
	if criteria.ID > 0 {
		stmt = stmt.Where(squirrel.Eq{`DAS.SCHOOL.ID`: criteria.ID})
	}
	if len(criteria.Name) > 0 {
		stmt = stmt.Where(squirrel.Eq{`DAS.SCHOOL.NAME`: criteria.Name})
	}
	if criteria.CityID > 0 {
		stmt = stmt.Where(squirrel.Eq{`DAS.SCHOOL.CITY_ID`: criteria.CityID})
	}
	if criteria.StateID > 0 {
		stmt = stmt.Join(`DAS.CITY C ON C.ID = DAS.SCHOOL.CITY_ID`).
			Where(squirrel.Eq{`C.STATE_ID`: criteria.StateID})
	}
	rows, err := stmt.RunWith(repo.Database).Query()
	schools := make([]businesslogic.School, 0)
	if err != nil {
		return schools, err
	}
	for rows.Next() {
		each := businesslogic.School{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.CityID,
			&each.CreateUserID,
			&each.DateTimeCreated,
			&each.UpdateUserID,
			&each.DateTimeUpdated,
		)
		schools = append(schools, each)
	}
	rows.Close()
	return schools, err
}
