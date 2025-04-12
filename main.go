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

	//timeObj := time.Now()
	var today functions.MyDate

	/*
		today.Month = int(timeNow.Month())
		today.Day = timeNow.Day()
		today.Dow = int(timeNow.Weekday())
		today.Year, today.Woy = timeNow.ISOWeek()
		today.Qtr = timeNow.Hour()*4 + timeNow.Minute()/15
	*/

	fmt.Println("Input Year:")
	_, _ = fmt.Scanf("%d", &today.Year)
	fmt.Println("Input Month:")
	_, _ = fmt.Scanf("%d", &today.Month)
	fmt.Println("Input Day:")
	_, _ = fmt.Scanf("%d", &today.Day)
	fmt.Println("Input Qtr:")
	_, _ = fmt.Scanf("%d", &today.Qtr)

	timeObj, err := time.Parse("2006-1-2",
		fmt.Sprintf("%d-%d-%d", today.Year, today.Month, today.Day))
	today.Dow = int(timeObj.Weekday())
	_, today.Woy = timeObj.ISOWeek()
	if err != nil {
		log.Fatal(err)
	}

	var plan, did string
	fmt.Println("input plan:")
	_, err = fmt.Scanln(&plan)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("input did:")
	_, err = fmt.Scanln(&did)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Writing to file...")
	functions.CallendarIn(db, today, plan, did)

	fmt.Println("Finished writing to file...")
	fmt.Println("exiting...")
}
