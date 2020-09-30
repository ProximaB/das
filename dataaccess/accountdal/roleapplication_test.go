package accountdal_test

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/dataaccess/accountdal"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestPostgresRoleApplicationRepository_CreateApplication(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	repo := accountdal.PostgresRoleApplicationRepository{
		Database:   db,
		SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}

	app := businesslogic.RoleApplication{}
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO DAS.ACCOUNT_ROLE_APPLICATION (ACCOUNT_ID, APPLIED_ROLE_ID, DESCRIPTION,
		STATUS_ID, APPROVAL_USER_ID, DATETIME_APPROVED, CREATE_USER_ID, DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED) VALUES`)
	mock.ExpectCommit()

	result := repo.CreateApplication(&app)
	assert.Nil(t, result, "should not throw an error even the data is empty")
}

func TestPostgresRoleApplicationRepository_SearchApplication(t *testing.T) {
	/*db, mock, _ := sqlmock.New()
	defer db.Close()
	repo := accountdal.PostgresRoleApplicationRepository{
		Database:   db,
		SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
	rows := sqlmock.NewRows(
		[]string{
			"ID", "ACCOUNT_ID", "APPLIED_ROLE_ID", "DESCRIPTION", "STATUS_ID", "APPROVAL_USER_ID", "DATETIME_APPROVED",
			"CREATE_USER_ID", "DATETIME_CREATED", "UPDATE_USER_ID", "DATETIME_UPDATED",
		},
	).AddRow(
		1, 33, businesslogic.AccountTypeOrganizer, "A brief Description", businesslogic.RoleApplicationStatusPending,
		nil, time.Now(), 33, time.Now(), 33, time.Now(),
	)

	accountRows := sqlmock.NewRows([]string{
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
		"DATETIME_CREATED",
		"DATETIME_UPDATED",
		"TOS_ACCEPTED",
		"PP_ACCEPTED",
		"BY_GUARDIAN",
		"GUARDIAN_SIGNATURE",
	}).AddRow(1, "abcd", 1, 3, "Alice", "", "Anderson", nil, "alice.anderson@email.com", "111-222-3333",
		 nil, nil, true, true, false, nil)

	criteria := businesslogic.SearchRoleApplicationCriteria{
		AccountID:      33,
		StatusID:       businesslogic.RoleApplicationStatusPending,
		ApprovalUserID: 21,
		AppliedRoleID:  businesslogic.AccountTypeOrganizer,
	}
	mock.ExpectQuery(`SELECT ID, ACCOUNT_ID, APPLIED_ROLE_ID, DESCRIPTION, STATUS_ID, APPROVAL_USER_ID, DATETIME_APPROVED,
		CREATE_USER_ID, DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED FROM DAS.ACCOUNT_ROLE_APPLICATION WHERE`).WithArgs(criteria.AccountID, criteria.AppliedRoleID, criteria.StatusID, criteria.ApprovalUserID).WillReturnRows(rows)
	mock.ExpectQuery(`SELECT ID, UID, ACCOUNT_STATUS_ID, USER_GENDER_ID, LAST_NAME, MIDDLE_NAMES, FIRST_NAME, DATE_OF_BIRTH,
		EMAIL, PHONE, DATETIME_CREATED, DATETIME_UPDATED, TOS_ACCEPTED, PP_ACCEPTED, BY_GUARDIAN, GUARDIAN_SIGNATURE FROM DAS.ACCOUNT WHERE ACCOUNT_ID = `).WithArgs(sqlmock.AnyArg()).WillReturnRows(accountRows)

	results, err := repo.SearchApplication(criteria)
	assert.Nil(t, err, "should not result in any error in creating the query")
	assert.NotNil(t, results, "should not result in empty results")*/
}

func TestPostgresRoleApplicationRepository_UpdateApplication(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	repo := accountdal.PostgresRoleApplicationRepository{
		Database:   db,
		SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}

	app := businesslogic.RoleApplication{}

	result := repo.UpdateApplication(app)
	assert.Error(t, result, "should not be able to update an empty application")

	app.ID = 3
	app.StatusID = businesslogic.RoleApplicationStatusPending
	userID := 84
	app.ApprovalUserID = &userID

	assert.Equal(t, *app.ApprovalUserID, 84)

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE DAS.ACCOUNT_ROLE_APPLICATION SET`).WillReturnResult(sqlmock.NewResult(3, 1))
	mock.ExpectCommit()

	result = repo.UpdateApplication(app)
	assert.Nil(t, result, "should not result in error if ID, Status, and Approval User are provided")

}
