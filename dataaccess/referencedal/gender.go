package referencedal

import (
	"database/sql"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/Masterminds/squirrel"
)

type PostgresGenderRepository struct {
	Database   *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

const (
	DAS_USER_GENDER_TABLE = "DAS.GENDER"
)

func (repo PostgresGenderRepository) GetAllGenders() ([]referencebll.Gender, error) {
	genders := make([]referencebll.Gender, 0)
	stmt := repo.SqlBuilder.Select(
		fmt.Sprintf(
			"%s, %s, %s, %s, %s, %s",
			common.PRIMARY_KEY,
			common.COL_NAME,
			common.COL_ABBREVIATION,
			common.COL_DESCRIPTION,
			common.COL_DATETIME_CREATED,
			common.COL_DATETIME_UPDATED,
		)).From(DAS_USER_GENDER_TABLE)

	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return genders, err
	}

	for rows.Next() {
		each := referencebll.Gender{}
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
