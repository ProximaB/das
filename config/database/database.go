package database

import (
	"database/sql"
)

// PostgresDatabase is the database connection used by the entire DAS system
var PostgresDatabase *sql.DB
