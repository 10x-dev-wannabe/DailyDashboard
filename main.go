package main

import (
	"database/sql"
	"fmt"
	"github.com/10x-dev-wannabe/DailyDashboard/functions"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

func main() {
	functions.Init()

	var statement string
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	timeNow := time.Now()
	yearNow, monthNow, dayNow := timeNow.Date()
	_, weekNow := timeNow.ISOWeek()
	dow := timeNow.Weekday()
	fmt.Println(yearNow)

	qtr := timeNow.Hour()*4 + timeNow.Minute()/15

	statement = "INSERT INTO callendar(year, month, day, woy, qtr, dow, plan, did) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = db.Exec(statement, yearNow, monthNow, dayNow, weekNow, qtr, dow, " ", " ")
	fmt.Println(err)

	rows, _ := db.Query("SELECT * FROM callendar")
	defer rows.Close()

	var a, b string

	for rows.Next() {
		err := rows.Scan(&yearNow, &monthNow, &dayNow, &weekNow, &qtr, &dow, &b, &a)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(yearNow, monthNow, dayNow, weekNow, qtr, dow, b, a)
	}

}
