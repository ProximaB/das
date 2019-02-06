package eventdal

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/common"
	"github.com/DancesportSoftware/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
)

const (
	dasCompetitionEventTemplateTable = "DAS.COMPETITION_EVENT_TEMPLATE"
)

type PostgresCompetitionEventTemplateRepository struct {
	Database   *sql.DB
	SQLBuilder squirrel.StatementBuilderType
}

func (repo PostgresCompetitionEventTemplateRepository) SearchCompetitionEventTemplates(criteria businesslogic.SearchCompetitionEventTemplateCriteria) ([]businesslogic.CompetitionEventTemplate, error) {
	if repo.Database == nil {
		return nil, errors.New(dalutil.DataSourceNotSpecifiedError(repo))
	}

	templates := make([]businesslogic.CompetitionEventTemplate, 0)
	stmt := repo.SQLBuilder.Select(
		fmt.Sprintf("%s, %s, %s, %s, %s, %s",
			common.ColumnPrimaryKey,
			common.COL_NAME,
			common.COL_DESCRIPTION,
			"FEDERATION",
			"TEMPLATE_EVENTS",
			common.ColumnDateTimeCreated)).
		From(dasCompetitionEventTemplateTable)
	if criteria.ID > 0 {
		stmt = stmt.Where(squirrel.Eq{common.ColumnPrimaryKey: criteria.ID})
	}
	if criteria.Name != "" {
		stmt = stmt.Where(squirrel.Eq{common.COL_NAME: criteria.Name})
	}

	rows, err := stmt.RunWith(repo.Database).Query()
	if err != nil {
		return templates, err
	}
	for rows.Next() {
		each := businesslogic.CompetitionEventTemplate{
			TargetFederation: businesslogic.Federation{},
		}
		list := ""
		scanErr := rows.Scan(
			&each.ID,
			&each.Name,
			&each.Description,
			&each.TargetFederation.Name,
			&list,
			&each.DateTimeCreate)
		if scanErr != nil {
			return templates, nil
		}

		templateEvents := make([]businesslogic.EventTemplate, 0)
		unmarshalErr := json.Unmarshal([]byte(list), templateEvents)
		if unmarshalErr != nil {
			return templates, unmarshalErr
		}
		each.TemplateEvents = templateEvents
		templates = append(templates, each)
	}

	return templates, nil
}
