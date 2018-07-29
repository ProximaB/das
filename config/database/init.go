// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
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

package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const envarDBConnectionString = "POSTGRES_CONNECTION"

var dbConnectionString = os.Getenv(envarDBConnectionString)

func openDatabaseConnection() {
	// for testing, use default connection
	if len(dbConnectionString) == 0 {
		log.Println("[error] cannot find database connection string")
	}

	var err error
	PostgresDatabase, err = sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Printf("[error] cannot establish connection to database: %s\n", err)
	}
	err = PostgresDatabase.Ping()
	if err != nil {
		log.Printf("[error] cannot ping database without error: %s\n", err.Error())
	}
	if err == nil {
		log.Println("[success] connected to database with the given connection string")
	}
	if PostgresDatabase == nil {
		log.Fatal("cannot create connection to the database")
	}
}

func init() {
	openDatabaseConnection()

	// Reference data
	CountryRepository.Database = PostgresDatabase
	StateRepository.Database = PostgresDatabase
	CityRepository.Database = PostgresDatabase
	FederationRepository.Database = PostgresDatabase
	DivisionRepository.Database = PostgresDatabase
	AgeRepository.Database = PostgresDatabase
	ProficiencyRepository.Database = PostgresDatabase
	StyleRepository.Database = PostgresDatabase
	DanceRepository.Database = PostgresDatabase
	SchoolRepository.Database = PostgresDatabase
	StudioRepository.Database = PostgresDatabase

	// account
	AccountRepository.Database = PostgresDatabase
	AccountTypeRepository.Database = PostgresDatabase
	AccountRoleRepository.Database = PostgresDatabase
	GenderRepository.Database = PostgresDatabase
	UserPreferenceRepository.Database = PostgresDatabase

	// Partnership request blacklist
	PartnershipRequestBlacklistRepository.Database = PostgresDatabase
	PartnershipRequestBlacklistReasonRepository.Database = PostgresDatabase

	// Partnership request
	PartnershipRequestRepository.Database = PostgresDatabase
	PartnershipRequestStatusRepository.Database = PostgresDatabase

	// Partnership
	PartnershipRepository.Database = PostgresDatabase

	// organizer
	OrganizerProvisionRepository.Database = PostgresDatabase
	OrganizerProvisionHistoryRepository.Database = PostgresDatabase

	// competition
	CompetitionStatusRepository.Database = PostgresDatabase
	CompetitionRepository.Database = PostgresDatabase

	// event
	EventRepository.Database = PostgresDatabase
	EventMetaRepository.Database = PostgresDatabase

	// competition entry
	AthleteCompetitionEntryRepository.Database = PostgresDatabase
	PartnershipCompetitionEntryRepository.Database = PostgresDatabase

	// event entry
	PartnershipEventEntryRepository.Database = PostgresDatabase
}
