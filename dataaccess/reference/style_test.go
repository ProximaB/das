package reference

import (
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var styleRepository = PostgresStyleRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresStyleRepository_SearchStyle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	styleRepository.Database = db

	rows := sqlmock.NewRows(
		[]string{
			"ID", "NAME", "DESCRIPTION", "CREATE_USER_ID", "DATETIME_CREATED", "UPDATE_USER_ID", "DATETIME_UPDATED",
		},
	).AddRow(1, "Standard", "International Standard", 3, time.Now(), 4, time.Now())

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	styles, _ := styleRepository.SearchStyle(reference.SearchStyleCriteria{})
	assert.NotZero(t, len(styles))
}
