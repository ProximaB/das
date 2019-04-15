package env

import (
	"log"
	"os"
	"strings"
)

func init() {
	if val, ok := os.LookupEnv(VarDatabaseDriver); ok && len(strings.TrimSpace(val)) != 0 {
		log.Printf("[info] %v is defined", VarDatabaseDriver)
		DatabaseDriver = strings.TrimSpace(val)
	} else {
		log.Printf("[warning] %v is missing or undefined", VarDatabaseDriver)
	}
	if val, ok := os.LookupEnv(VarDatabaseConnectionString); ok && len(strings.TrimSpace(val)) != 0 {
		log.Printf("[info] %v is defined", VarDatabaseConnectionString)
		DatabaseConnectionString = strings.TrimSpace(val)
	} else {
		log.Printf("[warning] %v is missing or undefined", VarDatabaseConnectionString)
	}
	if val, ok := os.LookupEnv(VarFirebaseAuthCredential); ok && len(strings.TrimSpace(val)) != 0 {
		log.Printf("[info] %v is defined", VarFirebaseAuthCredential)
		FirebaseAuthCredential = strings.TrimSpace(val)
	} else {
		log.Printf("[warning] %v is missing or undefined", VarFirebaseAuthCredential)
	}
}
