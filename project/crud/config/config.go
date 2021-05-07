package config

import (
	"database/sql"
)

func Connect() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "1000sean334"
	dbName := "department"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
