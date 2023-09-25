//go:build (darwin && cgo) || (linux && cgo)

package db

import (
	"os"

	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"gorm.io/gorm"
)

func ConnectSQLite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_FILE")), &gorm.Config{QueryFields: true})
	return db, err
}
