//go:build !(cgo && (linux || darwin))

package db

import (
	"os"

	"github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

func ConnectSQLite() (*gorm.DB, error) {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_FILE")), &gorm.Config{QueryFields: true})
	return db, err
}
