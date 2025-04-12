package functions

import (
	"database/sql"
	"fmt"
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

	// We update the data in the callendar
	sql := `UPDATE callendar SET
					plan = ?, did = ? WHERE
					year = ? AND
					month= ? AND
					day  = ? AND
					woy  = ? AND
					dow  = ? AND
					qtr  = ?`
	result, err := db.Exec(sql,
		plan,
		did,
		date.Year,
		date.Month,
		date.Day,
		date.Woy,
		date.Dow,
		date.Qtr)
	if err != nil {
		log.Fatal(err)
	}

	// If there is no data at the specified date, create row
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		sql = `INSERT INTO callendar
					(year, month, day, woy, dow, qtr, plan, did) 
	        VALUES(?, ?, ?, ?, ?, ?, ?, ?);`
		result, err = db.Exec(sql,
			date.Year,
			date.Month,
			date.Day,
			date.Woy,
			date.Dow,
			date.Qtr,
			plan,
			did)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("data modified sucessfully")
	return result
}
