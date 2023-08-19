package controller

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"
	"tugas-akhir-enigmacamp-go/controller"
	band "tugas-akhir-enigmacamp-go/controller/band"
	"tugas-akhir-enigmacamp-go/database"
	"tugas-akhir-enigmacamp-go/entity"
)

func ScheduleInterface() {
	// sambungkan database ===================
	db := database.ConnectDb()
	defer db.Close()

	// buat transaksi
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// =======================================

	// Mendapatkan waktu sekarang
	now := time.Now()

	// Mendapatkan nama hari dalam bahasa Inggris
	dayName := now.Weekday().String()
	allday := GetAllDay(tx)
	var hariTersedia []entity.DayOnSchadule
	// var choosenDay []entity.DayOnSchadule

	//=================================================
	fmt.Println("Hari yang Tersedia.")
	for _, v := range allday {
		if v.Id >= GetIdByNameDay(GetIndDay(dayName), tx) {
			hariTersedia = append(hariTersedia, v)
			fmt.Printf(" - Hari : %s\n", v.Name)
		}
	}
	//=================================================
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\nTulis hari yang dinginkan :")
	scanner.Scan()
	searchDay := scanner.Text()
	band.ErrorInput(scanner.Err(), tx)

	// =======================================

	fmt.Println("\nJadwal yang tersedia :")
	jadwal, _ := SearchByDay(searchDay, tx, hariTersedia)
	fmt.Println("ID | ruangan | Hari | Jam Mulai - Jam selesai")
	fmt.Println("==============================================")
	for _, v := range jadwal {
		fmt.Printf("%d | %s | %s | %s - %s\n", v.Id, v.Rooms, v.Day, v.ClockStart, v.ClockEnd)
	}
	fmt.Println("==============================================")

	var id int
	fmt.Print("\n\nPilih ID jadwal yang diinginkan: ")
	_, err = fmt.Scan(&id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// =======================================
	bandData := band.GetBand(tx)
	band.ShowingDataBand(bandData)
	var idband int
	fmt.Println("\nMasukkan Id band yang sesuai :")
	_, err = fmt.Scan(&idband)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// =======================================
	err = ChooseSchedule(idband, id, tx)
	controller.Validate(err, "Choose Schadule", tx)

	// commit transaction
	err = tx.Commit()
	if err != nil {
		controller.Validate(err, "commit band", tx)
	}

	fmt.Println("Berhasil claim jadwal.")
	JadwalTerisiUi()
}

func GetIndDay(day string) string {
	switch day {
	case "Sunday":
		return "minggu"
	case "Monday":
		return "senin"
	case "Tuesday":
		return "selasa"
	case "Wednesday":
		return "rabu"
	case "Thursday":
		return "kamis"
	case "Friday":
		return "jumat"
	case "Saturday":
		return "sabtu"
	default:
		return "Tidak ditemukan"
	}
}

func GetAllDay(tx *sql.Tx) (result []entity.DayOnSchadule) {
	sqlStatment := "SELECT id, name FROM schedule_days"

	rows, err := tx.Query(sqlStatment)
	if err != nil {
		controller.Validate(err, "select", tx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var days entity.DayOnSchadule
		if err := rows.Scan(&days.Id, &days.Name); err != nil {
			controller.Validate(err, "scan", tx)
			continue
		}
		result = append(result, days)
	}

	if err := rows.Err(); err != nil {
		controller.Validate(err, "rows", tx)
	}

	return result
}

func GetIdByNameDay(day string, tx *sql.Tx) int {
	allDays := GetAllDay(tx)
	var id int
	for _, v := range allDays {
		if v.Name == day {
			id = v.Id
		} else {
			continue
		}
	}
	return id
}

func SearchByDay(name string, tx *sql.Tx, hariTersedia []entity.DayOnSchadule) (result []entity.Schedule, err error) {
	sqlStatement := `
		SELECT 
			schedule_assigns.id, 
			schedule_days.name, 
			rooms.name,
			schedule_times.clock_start,
			schedule_times.clock_end
		FROM schedule_assigns
		LEFT JOIN schedule_days ON schedule_assigns.schedule_day_id = schedule_days.id
		LEFT JOIN rooms ON schedule_assigns.room_id = rooms.id
		LEFT JOIN schedule_times ON schedule_assigns.schedule_time_id = schedule_times.id
		WHERE schedule_days.name ILIKE $1
		AND schedule_assigns.band_id IS null;
	`

	rows, err := tx.Query(sqlStatement, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []entity.Schedule
	for rows.Next() {
		var schedule entity.Schedule
		err := rows.Scan(
			&schedule.Id,
			&schedule.Day,
			&schedule.Rooms,
			&schedule.ClockStart,
			&schedule.ClockEnd,
		)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	var newScheduleDay []entity.Schedule
	for _, v := range schedules {
		for _, w := range hariTersedia {
			if v.Day == w.Name {
				newScheduleDay = append(newScheduleDay, v)
			}
		}
	}

	var newScheduleTime []entity.Schedule
	// Mendapatkan waktu sekarang
	now := time.Now()
	// Mendapatkan nama hari dalam bahasa Inggris
	dayName := now.Weekday().String()

	fmt.Println(now)

	for _, v := range newScheduleDay {
		// Ambil bagian jam, menit, dan detik dari targetTime
		targetTime, _ := time.Parse("15:04:05", v.ClockStart[11:19])
		targetStartHour := targetTime.Hour()
		targetStartMinute := targetTime.Minute()
		targetStartSecond := targetTime.Second()

		TimeEnd, _ := time.Parse("15:04:05", v.ClockEnd[11:19])
		EndHour := TimeEnd.Hour()
		EndMinute := TimeEnd.Minute()
		EndSecond := TimeEnd.Second()

		// Ambil waktu saat ini
		currentHour := now.Hour()
		currentMinute := now.Minute()
		currentSecond := now.Second()

		band.ErrorInput(err, tx)

		if v.Day == GetIndDay(dayName) {
			if targetStartHour > currentHour ||
				(targetStartHour == currentHour && targetStartMinute > currentMinute) ||
				(targetStartHour == currentHour && targetStartMinute == currentMinute &&
					targetStartSecond > currentSecond) {
				v.ClockStart = strconv.Itoa(targetStartHour) + ":" + strconv.Itoa(targetStartMinute) + ":" + strconv.Itoa(targetStartSecond)
				v.ClockEnd = strconv.Itoa(EndHour) + ":" + strconv.Itoa(EndMinute) + ":" + strconv.Itoa(EndSecond)
				newScheduleTime = append(newScheduleTime, v)
			}
		} else {
			v.ClockStart = strconv.Itoa(targetStartHour) + ":" + strconv.Itoa(targetStartMinute) + ":" + strconv.Itoa(targetStartSecond)
			v.ClockEnd = strconv.Itoa(EndHour) + ":" + strconv.Itoa(EndMinute) + ":" + strconv.Itoa(EndSecond)
			newScheduleTime = append(newScheduleTime, v)
		}
	}
	return newScheduleTime, nil
}

func ChooseSchedule(bandId, scheduleId int, tx *sql.Tx) error {
	insertSQL := `
			UPDATE 
				schedule_assigns
			SET 
				band_id = $2
			WHERE 
				id = $1;`
	_, err := tx.Exec(insertSQL, scheduleId, bandId)
	return err
}

func GetJadwalTerisi(tx *sql.Tx) (result []entity.Schedule) {
	var hasil []entity.Schedule
	sqlStatement := `
		SELECT 
			schedule_assigns.id, 
			schedule_days.name, 
			rooms.name,
			schedule_times.clock_start,
			schedule_times.clock_end
		FROM schedule_assigns
		LEFT JOIN schedule_days ON schedule_assigns.schedule_day_id = schedule_days.id
		LEFT JOIN rooms ON schedule_assigns.room_id = rooms.id
		LEFT JOIN schedule_times ON schedule_assigns.schedule_time_id = schedule_times.id
		WHERE schedule_assigns.band_id IS NOT null;
	`

	rows, err := tx.Query(sqlStatement)
	if err != nil {
		controller.Validate(err, "select", tx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var Schedule entity.Schedule
		if err := rows.Scan(&Schedule.Id, &Schedule.Day, &Schedule.Rooms, &Schedule.ClockStart, &Schedule.ClockEnd); err != nil {
			controller.Validate(err, "scan", tx)
			continue
		}
		hasil = append(hasil, Schedule)
	}

	if err := rows.Err(); err != nil {
		controller.Validate(err, "rows", tx)
	}

	for _, v := range hasil {
		// Ambil bagian jam, menit, dan detik dari targetTime
		targetTime, _ := time.Parse("15:04:05", v.ClockStart[11:19])
		targetStartHour := targetTime.Hour()
		targetStartMinute := targetTime.Minute()
		targetStartSecond := targetTime.Second()

		TimeEnd, _ := time.Parse("15:04:05", v.ClockEnd[11:19])
		EndHour := TimeEnd.Hour()
		EndMinute := TimeEnd.Minute()
		EndSecond := TimeEnd.Second()

		band.ErrorInput(err, tx)

		v.ClockStart = strconv.Itoa(targetStartHour) + ":" + strconv.Itoa(targetStartMinute) + ":" + strconv.Itoa(targetStartSecond)

		v.ClockEnd = strconv.Itoa(EndHour) + ":" + strconv.Itoa(EndMinute) + ":" + strconv.Itoa(EndSecond)
		result = append(result, v)

	}

	return result
}
