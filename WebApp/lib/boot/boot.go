// Package boot handles the initialization of the web components.
package boot

import (
	"log"

	"github.com/huzhaer/qianxun/controller"
	"github.com/huzhaer/qianxun/lib/env"
	"github.com/huzhaer/qianxun/lib/flight"
	"github.com/huzhaer/teamlite_core/xsrf"

)

// RegisterServices sets up all the web components.
func RegisterServices(config *env.Info) {
	// Set up the session cookie store
	err := config.Session.SetupConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the MySQL database
	 mysqlDB, _ := config.MySQL.Connect(true)
	// mysqlDB.SetMaxIdleConns(100)
	// mysqlDB.SetMaxOpenConns(300)
	// Connect to the PgSql database
	// pgsqlDB, _ := config.PgSQL.Connect(true)

	// PostgreSQL set the Connetionpools variables
	// pgsqlDB.SetMaxIdleConns(100)
	// pgsqlDB.SetMaxOpenConns(300)

	// Load the controller routes
	controller.LoadRoutes()
	// Set up the views

	// Store the variables in flight
	flight.StoreConfig(*config)
	// Store the database connection in flight
	// flight.StoreDB(mysqlDB)
	//flight.StorePqsqlPoolConnection(pgsqlDB)

	flight.StoreDB(mysqlDB)

	// Store the csrf information
	flight.StoreXsrf(xsrf.Info{
		AuthKey: config.Session.CSRFKey,
		Secure:  config.Session.Options.Secure,
	})
}
