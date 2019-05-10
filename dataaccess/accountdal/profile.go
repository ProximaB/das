package accountdal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
)

type PostgresAthleteProfileRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

func (repo PostgresAthleteProfileRepository) GetAthleteProfile(sid string) (businesslogic.AthleteProfile, error) {
	output := businesslogic.AthleteProfile{}
	if repo.Database == nil {
		return output, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}
	row := repo.Database.QueryRow(fmt.Sprintf("SELECT * FROM GET_ATHLETE_PROFILE(%s)", sid))

	err := row.Scan()
	return output, err
}
