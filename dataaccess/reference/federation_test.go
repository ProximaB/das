package reference

import (
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var federationRepo = PostgresFederationRepository{

	Database:   nil,
	SqlBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresFederationRepository_SearchFederation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	federationRepo.Database = db
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
	federations, err := federationRepo.SearchFederation(&reference.SearchFederationCriteria{})
	assert.Nil(t, err, "should be able to read federation table")
	assert.NotZero(t, len(federations), "Database has more than 1 federation")
}
