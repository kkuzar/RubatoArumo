package pgsql

import (
	_ "database/sql"
	"fmt"
	"strings"
	"github.com/jmoiron/sqlx"
	"github.com/liudng/godump"
)


// Info holds the details for the PGSQL connection.
type Info struct {
	Username  string    `json:"Username"`
	Password  string    `json:"Password"`
	Database  string    `json:"Database"`
	Charset   string    `json:"Charset"`
	Hostname  string    `json:"Hostname"`
	Port      uint16    `json:"Port"`
	Parameter string    `json:"Parameter"`
	Timeout   string    `json:"Timeout"`
}
// *****************************************************************************
// Database Handling
// *****************************************************************************

// Connect to the database.
func (c Info) Connect(specificDatabase bool) (db *sqlx.DB, err error) {
	// Connect to database and ping
	godump.Dump(c.dsn(specificDatabase))

	Db , err := sqlx.Open("postgres",c.dsn(specificDatabase))

	return Db,err
}

// *****************************************************************************
// Postgres Specific
// *****************************************************************************

// DSN returns the Data Source Name.
func (c Info) dsn(includeDatabase bool) string {
	// Set defaults
	ci := c.setDefaults()

	// Build parameters
	param := ci.Parameter

	// If parameter is specified, add a question mark
	// Don't add one if a question mark is already there
	if len(ci.Parameter) > 0 && !strings.HasPrefix(ci.Parameter, "?") {
		param = "?" + ci.Parameter
	}

	// Add timeout
	if !strings.Contains(param, "collation") {
		if len(param) > 0 {
			param += "&connect_timeout=" + ci.Timeout
		} else {
			param = "?connect_timeout=" + ci.Timeout
		}
	}

	// Add charset
	//if !strings.Contains(param, "charset") {
	//	if len(param) > 0 {
	//		param += "&charset=" + ci.Charset
	//	} else {
	//		param = "?charset=" + ci.Charset
	//	}
	//}

	// Example: root:password@tcp(localhost:3306)/test
	s := fmt.Sprintf("%v:%v@tcp(%v:%d)/%v", ci.Username, ci.Password, ci.Hostname, ci.Port, param)

	if includeDatabase {
		// s = fmt.Sprintf("%v:%v@tcp(%v:%d)/%v%v", ci.Username, ci.Password, ci.Hostname, ci.Port, ci.Database, param)
		s = fmt.Sprintf("postgres://%v:%v@%v:%d/%v%v",ci.Username,ci.Password,ci.Hostname, ci.Port, ci.Database, param)
	}

	return s
}

// setDefaults sets the charset and collation if they are not set.
func (c Info) setDefaults() Info {
	ci := c

	if len(ci.Charset) == 0 {
		ci.Charset = "utf8"
	}

	return ci
}
