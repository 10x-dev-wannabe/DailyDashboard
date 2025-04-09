package functions

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type day struct {
	id   int
	note string
	toDo string
}

func Init() {
	// Open database connection
	db, err := sql.Open("sqlite3", "data.db")
	// Throw error if neded
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Create table

	// Define the callendar table where we store all data.
	statement, err := db.Prepare(`CREATE TABLE IF NOT EXISTS callendar (
		year  INTEGER,
		month INTEGER,
		day   INTEGER,
		woy   INTEGER,
		dow   STRING,
		qtr   INTEGER,
		plan  STRING,
		did   STRING,
		PRIMARY KEY(year, month, day, qtr)
		)`)
	if err != nil {
		log.Println("Error in creating table")
	}
	statement.Exec()

	statement, err = db.Prepare(`CREATE TABLE IF NOT EXISTS dailyNotes (
		year  INTEGER,
		month INTEGER,
		day   INTEGER,
		note  TEXT,
		toDo  TEXT,
		PRIMARY KEY(year, month, day)
		)`)
	if err != nil {
		log.Println("Error in creating table")
	}
	statement.Exec()

	statement, err = db.Prepare(`CREATE TABLE IF NOT EXISTS weeklyNotes (
		year INTEGER,
		woy  INTEGER,
		note TEXT,
		toDo TEXT,
		PRIMARY KEY(year, woy)
		)`)
	if err != nil {
		log.Println("Error in creating table")
	}
	statement.Exec()

	statement, err = db.Prepare(`CREATE TABLE IF NOT EXISTS monthlyNotes (
		year  INTEGER,
		month INTEGER,
		note  TEXT,
		toDo  TEXT,
		PRIMARY KEY(year, month)		
		)`)
	if err != nil {
		log.Println("Error in creating table")
	}
	statement.Exec()

	fmt.Println("Finished init")
}
