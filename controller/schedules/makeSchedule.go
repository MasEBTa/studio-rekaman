package controller

import (
	"database/sql"
	"fmt"
	"studio-room/controller"
	"studio-room/database"
)

func CreteSchedule() {
	if CheckIsScheduleNull() == 0 {
		db := database.ConnectDb()
		defer db.Close()

		tx, err := db.Begin()
		if err != nil {
			panic(err)
		}

		roomId := GetRoomId(tx)
		day := GetDayId(tx)
		time := GetTimeId(tx)
		InputSchedules(roomId, day, time, tx)

		err = tx.Commit()
		if err != nil {
			controller.Validate(err, "commit", tx)
		}
	}
}

func GetRoomId(tx *sql.Tx) (result []int) {
	sqlStatement := "SELECT id FROM rooms;"

	rows, err := tx.Query(sqlStatement)
	if err != nil {
		controller.Validate(err, "select", tx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			controller.Validate(err, "scan", tx)
			continue
		}
		result = append(result, id)
	}

	if err := rows.Err(); err != nil {
		controller.Validate(err, "rows", tx)
	}

	return result
}

func GetDayId(tx *sql.Tx) (result []int) {
	sqlStatement := "SELECT id FROM schedule_days;"

	rows, err := tx.Query(sqlStatement)
	if err != nil {
		controller.Validate(err, "select", tx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			controller.Validate(err, "scan", tx)
			continue
		}
		result = append(result, id)
	}

	if err := rows.Err(); err != nil {
		controller.Validate(err, "rows", tx)
	}

	return result
}

func GetTimeId(tx *sql.Tx) (result []int) {
	sqlStatement := "SELECT id FROM schedule_times;"

	rows, err := tx.Query(sqlStatement)
	if err != nil {
		controller.Validate(err, "select", tx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			controller.Validate(err, "scan", tx)
			continue
		}
		result = append(result, id)
	}

	if err := rows.Err(); err != nil {
		controller.Validate(err, "rows", tx)
	}

	return result
}

func InputSchedules(roomId, dayId, timeId []int, tx *sql.Tx) {
	sqlStatement := "INSERT INTO schedule_assigns(room_id, schedule_day_id, schedule_time_id, band_id) VALUES ($1, $2, $3, null);"

	for _, room := range roomId {
		for _, day := range dayId {
			for _, time := range timeId {
				_, err := tx.Exec(sqlStatement, room, day, time)
				if err != nil {
					controller.Validate(err, "insert", tx)
				}
			}
		}
	}

	fmt.Println("Success Creating schedule")
}
