package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	log.Println("Opening the database file...")
	DB, err = sql.Open("sqlite3", "./forum.db") // open db if it exist, create and open it if it does not exist
	if err != nil {
		return fmt.Errorf("error opening database --InitDB()--%w: ", err)
	}

	log.Println("Reading the internal/db/schema.sql  file...")
	schema, err := os.ReadFile("internal/db/schema.sql")
	if err != nil {
		return fmt.Errorf("error reading schema --InitDB()-->--ReadFile-- %w", err)
	}

	log.Println("Executing the internal/db/schema.sql  file...")
	_, err = DB.Exec(string(schema))
	if err != nil {
		return fmt.Errorf("error executing schema --InitDB()-->--DB.Exec--%w", err)
	}

	log.Println("Database initialized")

	return nil
}

//resource management -> avoid file corruption or memory leak
func CloseDB() {
	err := DB.Close()
	if err != nil {
		log.Printf("error closing database --CloseDB()-- %v", err)
	}
}
