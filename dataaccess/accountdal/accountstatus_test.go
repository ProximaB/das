package accountdal_test

import (
	"github.com/DancesportSoftware/das/dataaccess/accountdal"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var accountStatusRepo = accountdal.PostgresAccountStatusRepository{
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
		[]string{"ID", "NAME", "ABBREVIATION", "DESCRIPTION", "DATETIME_CREATED", "DATETIME_UPDATED"},
	).AddRow(1, "Activated", "A", "Account is activated and is usable", time.Now(), time.Now()).AddRow(1, "Suspended", "S", "Account is suspended for violation of ToS", time.Now(), time.Now())
	mock.ExpectQuery(`SELECT ID, NAME, ABBREVIATION, DESCRIPTION, DATETIME_CREATED, 
		DATETIME_UPDATED FROM DAS.ACCOUNT_STATUS`).WillReturnRows(rows)
	status, err := accountStatusRepo.GetAccountStatus()
	assert.Nil(t, err)
	assert.EqualValues(t, 2, len(status))
}
