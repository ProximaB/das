// Dancesport Application System (DAS)
// Copyright (C) 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package entrydal_test

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/entrydal"
	"github.com/DancesportSoftware/das/dataaccess/util"
	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var partnershipEventEntryRepo = entrydal.PostgresPartnershipEventEntryRepository{
	Database:   nil,
	SQLBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
}

func TestPostgresPartnershipEventEntryRepository_CreatePartnershipEventEntry(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	entry := businesslogic.PartnershipEventEntry{
		PartnershipID: 37,
		EventEntry: businesslogic.EventEntry{
			EventID:     991,
			CheckInTime: time.Now(),
		},
	}

	err := partnershipEventEntryRepo.CreatePartnershipEventEntry(&entry)
	assert.NotNil(t, err, dalutil.ErrorNilDatabase)

	partnershipEventEntryRepo.Database = db

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO DAS.EVENT_ENTRY_PARTNERSHIP (EVENT_ID, PARTNERSHIP_ID, LEADTAG, CHECKIN_IND,
		CHECKIN_DATETIME, CREATE_USER_ID, DATETIME_CREATED, UPDATE_USER_ID, DATETIME_UPDATED)`)
	mock.ExpectCommit()

	err = partnershipEventEntryRepo.CreatePartnershipEventEntry(&entry)
	assert.Nil(t, err, "should insert legitimate PartnershipEventEntry data without error")
}
