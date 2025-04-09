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
	fmt.Println("init...")
	functions.Init()

	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("oppened database...")

	timeNow := time.Now()
	var today functions.MyDate

	today.Month = int(timeNow.Month())
	today.Day = timeNow.Day()
	today.Dow = int(timeNow.Weekday())
	today.Year, today.Woy = timeNow.ISOWeek()
	today.Qtr = timeNow.Hour()*4 + timeNow.Minute()/15

	var a, b string
	a = "helu"
	b = ""

	fmt.Println("Writing to file...")
	functions.CallendarIn(db, today, a, b)

	fmt.Println("Finished writing to file...")
	fmt.Println("exiting...")
}
