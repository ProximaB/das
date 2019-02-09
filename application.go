package main

import (
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/config/routes"
	"github.com/gorilla/csrf"
	"log"
	"net/http"
	"os"
)

const (
	envAppPort = "APP_PORT"
	envCsrfKey = "CSRF_KEY"
)

func main() {
	defer database.PostgresDatabase.Close() // database connection will not close until server is shutdown
	router := routes.NewDasRouter()

	if database.PostgresDatabase == nil {
		log.Println("[error] database connection is closed")
	}
	if database.PostgresDatabase.Ping() != nil {
		log.Println("[error] database is not responding to ping")
	}

	csrfKey := os.Getenv(envCsrfKey)
	if csrfKey != "" {
		csrfProtector := csrf.Protect([]byte(csrfKey))
		http.Handle("/", csrfProtector(router))
		log.Printf("[info] CSRF_KEY is added to request handlers")
	} else {
		http.Handle("/", router)
		log.Printf("[warning] CSRF_KEY is not defined and DAS is not protected from CSRF")
	}

	port := os.Getenv(envAppPort)
	if port == "" {
		port = "8080" // default port for Google CLoud
	}

	log.Printf("[info] DAS will be running on port " + port)
	log.Fatalf("[fatal] %v", http.ListenAndServe(":"+port, nil))
}
