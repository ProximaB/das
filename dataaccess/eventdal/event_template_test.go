package eventdal_test

import (
	"database/sql"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/eventdal"
	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostgresCompetitionEventTemplateRepository_SearchCompetitionEventTemplates(t *testing.T) {
	db, _ := sql.Open("postgres", "user=dasdev password=dAs!@#$1234 dbname=das sslmode=disable")
	repo := eventdal.PostgresCompetitionEventTemplateRepository{
		db,
		squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}

	output, err := repo.SearchCompetitionEventTemplates(businesslogic.SearchCompetitionEventTemplateCriteria{
		ID: 1,
	})
	assert.Nil(t, err, "should be able to retrieve templates with ID = 1")
	assert.Equal(t, 1, len(output), "should return exactly one template when ID is provided")
	assert.True(t, len(output[0].TemplateEvents) > 1, "template has many events")
}
