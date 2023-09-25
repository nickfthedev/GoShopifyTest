package db

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

// Database instance
var DB *gorm.DB

// Connect to database and put instance into var DB
func ConnectDB() {
	var err error

	switch os.Getenv("DB_DRIVER") {
	case "POSTGRESQL":
		fmt.Println("DB Driver: PostgreSQL")
		DB, err = ConnectPostgre()
		if err != nil {
			panic(err)
		}

	case "SQLITE":
		fmt.Println("DB Driver: SQLite")
		DB, err = ConnectSQLite()
		if err != nil {
			panic(err)
		}
	default:
		fmt.Println("Check your Databasedriver!")
		panic("no db driver")
	}

}
