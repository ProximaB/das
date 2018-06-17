package main

import (
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/config/routes"
	"google.golang.org/appengine"
	"log"
	"net/http"
)

func main() {

	defer database.PostgresDatabase.Close() // database connection will not close until server is shutdown
	router := routes.DasRouter()

	if database.PostgresDatabase == nil {
		log.Fatal("database connection is closed before service started")
	}
	if database.PostgresDatabase.Ping() != nil {
		log.Fatal("cannot establish connectionto the database")
	}

	http.Handle("/", router)
	appengine.Main() // to run this on app engine, do not make router listen to any particular port
}
