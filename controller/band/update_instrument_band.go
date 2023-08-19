package controller

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"studio-room/controller"
	"studio-room/database"
	"studio-room/entity"
)

func UpdateInstruentBandUI() {
	// sambungkan database ===================
	db := database.ConnectDb()
	defer db.Close()

	// buat transaksi
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// =======================================
	// showing band =========================
	band := GetBand(tx)
	ShowingDataBand(band)
	// =======================================

	// input id band wanna update ============
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\nSilahkan masukkan Id band Untuk melihat detail Instrument :")
	scanner.Scan()
	idBand := scanner.Text()
	ErrorInput(scanner.Err(), tx)

	// Mengonversi idBand menjadi integer
	bandID, err := strconv.Atoi(idBand)
	ErrorInput(err, tx)
	// =======================================

	// Get Data band wanna update ============
	theBand := GetBandById(bandID, tx)
	// Get Data Instrument band ==============
	instrumentBand := GetInstrumentOfBand(bandID, tx)

	fmt.Printf("\n%s, Instrument Yang Digunakan :\n", theBand.NameBand)
	if len(instrumentBand) == 0 {
		fmt.Println("Belum Memilih Instrument")
	}
	for _, v := range instrumentBand {
		fmt.Printf(" - %d %s\n", v.Total, v.Name)
	}
	fmt.Println("\nSilahkan masukkan update yang diperlukan (ketik kata kuncinya) :")
	fmt.Println(" - tambah : menambah instrument baru")
	fmt.Println(" - edit : mengubah jumlah instrument yang digunakan")
	fmt.Println(" - delete : menghapus instrument yang telah dipilih")

	// input id band wanna update ============
	fmt.Println("\n(tambah/edit/delete) :")
	scanner.Scan()
	kataKunci := scanner.Text()
	ErrorInput(scanner.Err(), tx)

	switch kataKunci {
	case "tambah":
		TambahInstrument(instrumentBand, tx, bandID)
	case "edit":
		TambahJumlah(instrumentBand, tx, bandID)
	case "delete":
		DeleteInstrumentInBand(tx, bandID, instrumentBand)
	default:
		fmt.Println("Tidak ada fungsi yang dijalankan")
	}

	// Get Data Instrument band ==============
	instrumentBand = GetInstrumentOfBand(bandID, tx)

	fmt.Printf("\n%s, Instrument Yang Digunakan :\n", theBand.NameBand)
	if len(instrumentBand) == 0 {
		fmt.Println("Belum Memilih Instrument")
	}
	for _, v := range instrumentBand {
		fmt.Printf(" - %d %s\n", v.Total, v.Name)
	}

	// commit transaction ====================
	err = tx.Commit()
	if err != nil {
		controller.Validate(err, "commit band", tx)
	}
	// =======================================
}

func TambahInstrument(instrument []entity.InstrumentAndTotal, tx *sql.Tx, bandId int) {
	if len(instrument) < 4 {
		fmt.Println("\nInstrument yang bisa digunakan :")
		// ambil data instrument =================
		instruments := GetInstrument(tx)
		for _, v := range instruments {
			if isHasData(instrument, v.Instruments) {
				continue
			} else {
				fmt.Printf("key:%d. %s\n", v.Id, v.Instruments)
			}
		}

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Tuliskan instrumen yang digunakan :")
		scanner.Scan()
		instrumentUsed := scanner.Text()
		ErrorInput(scanner.Err(), tx)

		// Pisahkan input menjadi slice berdasarkan spasi
		inputInstrumentIDs := strings.Split(instrumentUsed, " ")

		var selectedInstruments []int

		for _, idStr := range inputInstrumentIDs {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Input instrumen tidak valid:", err)
			}
			selectedInstruments = append(selectedInstruments, id)
		}

		var newData []int
		if len(selectedInstruments) == 1 {
			newData = selectedInstruments
		}
		if len(instrument) > 0 {
			for _, v := range selectedInstruments {
				for _, w := range instruments {
					if v == w.Id {
						if isHasData(instrument, w.Instruments) {
							newData = removeFromSlice(selectedInstruments, v)
						} else {
							continue
						}
					}
				}
			}
		} else {
			newData = selectedInstruments
		}

		// jumlah instrumen
		fmt.Println("\nMasukkan jumlah instrumen yang dibutuhkan.")
		for _, id := range newData {
			for _, v := range instruments {
				if id == v.Id {
					var jml int
					fmt.Printf("Jumlah %s:\n", v.Instruments)
					if isHasData(instrument, v.Instruments) {
						fmt.Printf("instrument %s, sudah digunakan.\n", v.Instruments)
					} else {
						_, err := fmt.Scan(&jml)

						ErrorInput(err, tx)

						_, err = insertBandInstrument(tx, bandId, id, jml)
						ErrorInput(err, tx)
					}
				}
			}
		}
	}
}

func TambahJumlah(instrument []entity.InstrumentAndTotal, tx *sql.Tx, bandId int) {
	fmt.Println(instrument)
	// ambil data instrument =================
	instruments := GetInstrument(tx)
	var newData []int
	for _, v := range instruments {
		for _, vv := range instrument {
			if v.Instruments == vv.Name {
				if isHasData(instrument, v.Instruments) {
					newData = append(newData, v.Id)
				} else {
					continue
				}
			}
		}
		if isHasData(instrument, v.Instruments) {
			continue
		} else {
			fmt.Printf("key:%d. %s\n", v.Id, v.Instruments)
		}
	}
	// jumlah instrumen
	fmt.Println("\nMasukkan jumlah instrumen yang dibutuhkan.")
	for _, id := range newData {
		for _, v := range instruments {
			if id == v.Id {
				var jml int
				fmt.Printf("Jumlah %s:\n", v.Instruments)
				_, err := fmt.Scan(&jml)

				ErrorInput(err, tx)

				err = InputUpdateJumlahBand(tx, bandId, id, jml)
				ErrorInput(err, tx)
			}
		}
	}
}

func isHasData(data []entity.InstrumentAndTotal, target string) bool {
	for _, v := range data {
		if v.Name == target {
			return true
		}
	}
	return false
}

func removeFromSlice(slice []int, element int) []int {
	var result []int
	for _, val := range slice {
		if val != element {
			result = append(result, val)
		}
	}
	return result
}

func InputUpdateJumlahBand(tx *sql.Tx, bandId int, instrumentId int, total int) error {
	insertSQL := `
			UPDATE 
				band_instruments
			SET 
				total = $3
			WHERE 
				band_id = $1
			AND 
				instrument_type_id = $2;`
	_, err := tx.Exec(insertSQL, bandId, instrumentId, total)
	return err
}

func DeleteInstrumentInBand(tx *sql.Tx, bandId int, instrumentBand []entity.InstrumentAndTotal) {
	if len(instrumentBand) > 0 {
		fmt.Println("\nInstrument yang bisa digunakan :")
		// ambil data instrument =================
		instruments := GetInstrument(tx)
		for _, v := range instruments {
			if isHasData(instrumentBand, v.Instruments) {
				fmt.Printf("Id:%d. Name:%s\n", v.Id, v.Instruments)
			} else {
				continue
			}
		}
		// input id instrument wanna delete ============
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Tuliskan instrumen yang ingin duhapus :")
		scanner.Scan()
		instrumentUsed := scanner.Text()
		ErrorInput(scanner.Err(), tx)

		// Pisahkan input menjadi slice berdasarkan spasi
		inputInstrumentIDs := strings.Split(instrumentUsed, " ")

		var selectedInstruments []int

		for _, idStr := range inputInstrumentIDs {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Input instrumen tidak valid:", err)
			}
			selectedInstruments = append(selectedInstruments, id)
		}
		// =======================================
		err := UpdateDeleteInstrumentInBand(tx, bandId, selectedInstruments)
		if err != nil {
			controller.Validate(err, "delete", tx)
		} else {
			fmt.Println("Success Delete Instrument\n=================")
		}
	} else {
		fmt.Println("Instrument Kosong")
	}
}

func UpdateDeleteInstrumentInBand(tx *sql.Tx, idBand int, idInst []int) error {
	var err error
	for _, v := range idInst {
		deleteSql := `
			DELETE FROM 
				band_instruments
			WHERE 
				band_id = $1
			AND 
				instrument_type_id = $2;`
		_, err = tx.Exec(deleteSql, idBand, v)
		return err
	}
	return err
}
