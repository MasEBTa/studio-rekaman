package controller

import (
	"database/sql"
	"fmt"
	"studio-room/database"
)

func CheckIsScheduleNull() int {
	fmt.Println("Checking is Schedule Null")
	db := database.ConnectDb()
	defer db.Close()

	sqlStatement := "SELECT COUNT(*) AS result FROM schedule_assigns;"

	row, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer row.Close()

	result := scanResult(row)
	defer fmt.Println("Checking Schedule Success")

	return result
}

func scanResult(row *sql.Rows) (result int) {
	row.Next()
	if err := row.Scan(&result); err != nil {
		panic(err)
	}
	return result
}
