package eventdal_test

import (
	"github.com/ProximaB/das/dataaccess/eventdal"
	"github.com/Masterminds/squirrel"
	"testing"
)

var eventRepo = eventdal.PostgresEventRepository{
	Database:   nil,
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresEventRepository_CreateEvent(t *testing.T) {
	/*	db, mock, _ := sqlmock.New()
		defer db.Close()

		event := businesslogic.Event{
			CompetitionID: 12,
			FederationID:  18,
			DivisionID:    11,
			AgeID:         8,
			ProficiencyID: 1,
			StyleID:       12,
		}

		err := eventRepo.CreateEvent(&event)
		assert.NotNil(t, err, dalutil.ErrorNilDatabase)

		eventRepo.Database = db

		mock.ExpectBegin()
		mock.ExpectExec(`INSERT INTO DAS.EVENT (COMPETITION_ID, CATEGORY_ID, FEDERATION_ID, DIVISION_ID, AGE_ID,
			PROFICIENCY_ID, STYLE_ID, DESCRIPTION, EVENT_STATUS_ID, CREATE_USER_ID, DATETIME_CREATED, UPDATE_USER_ID,
			DATETIME_UPDATED)`)
		mock.ExpectCommit()

		err = eventRepo.CreateEvent(&event)
		assert.Nil(t, err, "should insert legitimate Event data without error")*/
}
