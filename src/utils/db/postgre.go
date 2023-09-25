package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgre() (*gorm.DB, error) {
	var err error
	var db *gorm.DB
	var dsn string
	if os.Getenv("PG_DB_URL") != "" {
		dsn = os.Getenv("PG_DB_URL")
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("PG_DB_HOST"),
			os.Getenv("PG_DB_USER"),
			os.Getenv("PG_DB_PASS"),
			os.Getenv("PG_DB_NAME"),
			os.Getenv("PG_DB_PORT"),
		)
	}
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{QueryFields: true})
	if err != nil {
		log.Println(err)
		panic("Failled to connect to Database. ")
	}
	return db, nil

}
