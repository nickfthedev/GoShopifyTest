package db

import (
	"fmt"
	"os"
	"testing"
)

func TestConnectDB(t *testing.T) {

	os.Setenv("DB_DRIVER", "SQLITE")
	dbfilename := "../../tmp/test.db"
	os.Setenv("DB_FILE", dbfilename)
	os.Remove(dbfilename)
	fmt.Println("tmp/test.db has been removed")
	// Test SQLite in Folder tests/test.db
	ConnectDB()
	fmt.Println("tmp/test.db has been created & connected")
	ConnectDB()
	fmt.Println("tmp/test.db has connected")

	// Test PostgreSQL
	// ConnectDB()
	// fmt.Println("Connection test to postgresql db succeded")

}
