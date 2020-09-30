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

type PostgresGenderRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

const (
	DAS_USER_GENDER_TABLE = "DAS.GENDER"
)

func (repo PostgresGenderRepository) GetAllGenders() ([]businesslogic.Gender, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	genders := make([]businesslogic.Gender, 0)
	stmt := repo.SqlBuilder.Select(
		fmt.Sprintf(
			"%s, %s, %s, %s, %s, %s",
			common.ColumnPrimaryKey,
			common.COL_NAME,
			common.ColumnAbbreviation,
			common.COL_DESCRIPTION,
			common.ColumnDateTimeCreated,
			common.ColumnDateTimeUpdated,
		)).From(DAS_USER_GENDER_TABLE)

	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return genders, err
	}

	for rows.Next() {
		each := businesslogic.Gender{}
		rows.Scan(
			&each.ID,
			&each.Name,
			&each.Abbreviation,
			&each.Description,
			&each.DateTimeCreated,
			&each.DateTimeUpdated,
		)
		genders = append(genders, each)
	}
	rows.Close()
	return genders, err
}
