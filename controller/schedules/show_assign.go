package controller

import (
	"fmt"
	"studio-room/database"
)

func JadwalTerisiUi() {
	// sambungkan database ===================
	db := database.ConnectDb()
	defer db.Close()

	// buat transaksi
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// =======================================

	jadwal := GetJadwalTerisi(tx)
	fmt.Println("ID | ruangan | Hari | Jam Mulai - Jam selesai")
	fmt.Println("==============================================")
	for _, v := range jadwal {
		fmt.Printf("%d | %s | %s | %s - %s\n", v.Id, v.Rooms, v.Day, v.ClockStart, v.ClockEnd)
	}
	fmt.Println("==============================================")
}
