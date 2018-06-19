// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func openDatabaseConnection() {
	connectionString := os.Getenv("POSTGRES_CONNECTION")
	// for testing, use default connection
	if len(connectionString) == 0 {
		log.Println("using default connection string")
		connectionString = `user=dasdev password=dAs\!@#\$1234 dbname=das sslmode=disable`
	} else {
		log.Println("using connection string from environment variable")
	}

	var err error
	PostgresDatabase, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Printf("[error] cannot establish connection to database: %s\n", err)
	}
	err = PostgresDatabase.Ping()
	if err != nil {
		log.Printf("[error] cannot ping database without error: %s\n", err.Error())
	}
	if err == nil {
		log.Println("[success] connected to database with given connection string")
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
	GenderRepository.Database = PostgresDatabase

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
	CompetitionEntryRepository.Database = PostgresDatabase

	// event entry
	EventEntryRepository.Database = PostgresDatabase
}
