package functions

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type MyDate struct {
	Year  int
	Month int
	Day   int
	Woy   int
	Dow   int
	Qtr   int
}

func CallendarIn(db *sql.DB, date MyDate, plan string, did string) sql.Result {
	sql := `INSERT INTO callendar
					(year, month, day, woy, dow, qtr) 
	        VALUES(?, ?, ?, ?, ?, ?)`
	result, err := db.Exec(sql,
		date.Year,
		date.Month,
		date.Day,
		date.Woy,
		date.Dow,
		date.Qtr)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
