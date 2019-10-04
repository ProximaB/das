package database

import (
	"database/sql"
	"github.com/DancesportSoftware/das/env"
	_ "github.com/lib/pq"

	"log"
)

func openDatabaseConnection() {
	if len(env.DatabaseDriver) == 0 {
		log.Println("[error] cannot find database driver")
	}
	// for testing, use default connection
	if len(env.DatabaseConnectionString) == 0 {
		log.Println("[error] cannot find database connection string")
	}

	var err error
	PostgresDatabase, err = sql.Open(env.DatabaseDriver, env.DatabaseConnectionString)
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
	RoleApplicationRepository.Database = PostgresDatabase
	RoleApplicationStatusRepository.Database = PostgresDatabase

	// Partnership request blacklist
	PartnershipRequestBlacklistRepository.Database = PostgresDatabase
	PartnershipRequestBlacklistReasonRepository.Database = PostgresDatabase

	// Partnership request
	PartnershipRequestRepository.Database = PostgresDatabase
	PartnershipRequestStatusRepository.Database = PostgresDatabase

	// Partnership
	PartnershipRoleRepository.Database = PostgresDatabase
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
	EventDanceRepository.Database = PostgresDatabase
	CompetitionEventTemplateRepository.Database = PostgresDatabase

	// competition entry
	AthleteCompetitionEntryRepository.Database = PostgresDatabase
	PartnershipCompetitionEntryRepository.Database = PostgresDatabase

	// event entry
	AthleteEventEntryRepository.Database = PostgresDatabase
	PartnershipEventEntryRepository.Database = PostgresDatabase
}
