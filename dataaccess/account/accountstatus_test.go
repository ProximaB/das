package account

import (
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var accountStatusRepo = PostgresAccountStatusRepository{
	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresAccountStatusRepository_GetAccountStatus(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	accountStatusRepo.Database = db
	rows := sqlmock.NewRows(
		[]string{
			"ID",
			"NAME",
			"ABBREVIATION",
			"YEAR_FOUNDED",
			"COUNTRY_ID",
			"CREATE_USER_ID",
			"DATETIME_CREATED",
			"UPDATE_USER_ID",
			"DATETIME_UPDATED",
		},
	).AddRow(1,
		"Pre Teen I",
		"Pre",
		1233,
		8,
		2,
		time.Now(),
		3,
		time.Now())
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	status, err := accountStatusRepo.GetAccountStatus()
	assert.Nil(t, err)
	assert.EqualValues(t, 4, len(status))
}
