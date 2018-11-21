package env

import (
	"log"
	"os"
	"strings"
)

func init() {
	if val, ok := os.LookupEnv(VarDatabaseConnectionString); ok && len(strings.TrimSpace(val)) != 0 {
		log.Printf("[info] %v is defined", VarDatabaseConnectionString)
		DatabaseConnectionString = strings.TrimSpace(val)
	} else {
		log.Printf("[warning] %v is missing or undefined", VarDatabaseConnectionString)
	}
	if val, ok := os.LookupEnv(VarFirebaseServiceAccountKey); ok && len(strings.TrimSpace(val)) != 0 {
		log.Printf("[info] %v is defined", VarFirebaseServiceAccountKey)
		FirebaseServiceAccountKey = strings.TrimSpace(val)
	} else {
		log.Printf("[warning] %v is missing or undefined", VarFirebaseServiceAccountKey)
	}
}
