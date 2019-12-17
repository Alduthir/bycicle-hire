package connection

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const dbName = "vanderbinckes"
const dbUser = "gouser"
const dbPassword = "test123"

// Uses the stored connectionstring to create a DB object and returns it.
func FetchDatabaseConnection() *sql.DB {
	connectionString := fmt.Sprintf("%s:%s@/%s?parseTime=true", dbUser, dbPassword, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic("Geen verbinding kunnen maken met de database")
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
