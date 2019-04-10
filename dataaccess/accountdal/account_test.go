package accountdal_test

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/accountdal"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

var accountRepository = accountdal.PostgresAccountRepository{
	Database:   nil,
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresAccountRepository_SearchAccount(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	res1, err := accountRepository.SearchAccount(businesslogic.SearchAccountCriteria{})
	assert.NotNil(t, err, "should return an error when database connection is not specified")
	assert.Nil(t, res1, "should not return a concrete object if database connection does not even exist")

	accountRepository.Database = db
	rows := sqlmock.NewRows(
		[]string{
			"ID",
			"UID",
			"ACCOUNT_STATUS_ID",
			"USER_GENDER_ID",
			"LAST_NAME",
			"MIDDLE_NAMES",
			"FIRST_NAME",
			"DATE_OF_BIRTH",
			"EMAIl",
			"PHONE",
			"EMAIL_VERIFIED",
			"PHONE_VERIFIED",
			"HASH_ALGORITHM",
			"PASSWORD_SALT",
			"PASSWORD_HASH",
			"DATETIME_CREATED",
			"DATETIME_UPDATED",
			"TOS_ACCEPTED",
			"PP_ACCEPTED",
			"BY_GUARDIAN",
			"GUARDIAN_SIGNATURE",
		},
	)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	results, err := accountRepository.SearchAccount(businesslogic.SearchAccountCriteria{
		Email: "test",
	})

	assert.Nil(t, err, "Database schema should match")
	assert.NotNil(t, results, "should return at least empty slice of accounts")

	accountRepository.Database = nil
}

// TODO: test is rigged
func TestPostgresAccountRepository_CreateAccount(t *testing.T) {
	/*db, mock, _ := sqlmock.New()
	defer db.Close()
	rows := sqlmock.NewRows(
		[]string{
			"ID",
		}).AddRow(1)
	account := businesslogic.Account{
		UID:             "abcd-efg",
		AccountStatusID: businesslogic.AccountStatusUnverified,
		FirstName:       "Alice",
		LastName:        "Anderson",
	}
	noDbErr := accountRepository.CreateAccount(&account)
	assert.NotNil(t, noDbErr, "should return an error if the data source is not specified")

	accountRepository.Database = db
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO DAS.ACCOUNT (UID, ACCOUNT_STATUS_ID, USER_GENDER_ID, LAST_NAME, MIDDLE_NAMES, FIRST_NAME, DATE_OF_BIRTH, EMAIL, PHONE, DATETIME_CREATED, DATETIME_UPDATED, TOS_ACCEPTED, PP_ACCEPTED, BY_GUARDIAN, GUARDIAN_SIGNATURE) VALUES `).WillReturnRows(rows)
	mock.ExpectCommit()
	results := accountRepository.CreateAccount(&account)

	assert.Nil(t, results, "should not return an error if data is correct")*/
}
