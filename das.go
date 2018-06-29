// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package main

import (
	"log"
	"net/http"

	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/config/routes"
	"google.golang.org/appengine"
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

	http.Handle("/", router)
	appengine.Main() // to run this on app engine, do not make router listen to any particular port
}
