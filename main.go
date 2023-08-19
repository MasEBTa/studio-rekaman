package main

import (
	"fmt"
	"os"

	band "tugas-akhir-enigmacamp-go/controller/band"
	schedule "tugas-akhir-enigmacamp-go/controller/schedules"

	_ "github.com/lib/pq" // Impor driver PostgreSQL
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "makeSchedule":
			schedule.CreteSchedule()
		case "newBand":
			band.AddBand()
		case "daftarBand":
			band.ShowingBand()
		case "deleteBand":
			band.DeletingBandUi()
		case "updateBand":
			band.UpdatingBandUi()
		case "updateInstruments":
			band.UpdateInstruentBandUI()
		case "jadwal":
			schedule.ScheduleInterface()
		case "jadwalTerisi":
			schedule.JadwalTerisiUi()
		default:
			fmt.Println("Tidak ada fungsi yang dijalankan")
		}
	} else {
		fmt.Println("Fungsi main dijalankan")
	}
}
