package database

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost" // default
	port     = 5432        // default
	user     = "postgres"  // drfault
	password = "12345678"
	dbname   = "studio" // sesuaikan
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// connection to database
func ConnectDb() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
