package controller

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"tugas-akhir-enigmacamp-go/controller"
	"tugas-akhir-enigmacamp-go/database"
	"tugas-akhir-enigmacamp-go/entity"
)

func UpdatingBandUi() {
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
	fmt.Println("\nSilahkan masukkan Id band yang ingin diupdate :")
	scanner.Scan()
	idBand := scanner.Text()
	ErrorInput(scanner.Err(), tx)

	// Mengonversi idBand menjadi integer
	bandID, err := strconv.Atoi(idBand)
	ErrorInput(err, tx)
	// =======================================

	// Get Data band wanna update ============
	theBand := GetBandById(bandID, tx)

	if theBand.Id != 0 {
		dataBandBaru := UpdateBand(theBand, tx)
		fmt.Printf("\nData band '%s' telah diperbarui :\n", theBand.NameBand)
		fmt.Println("Nama Band :", dataBandBaru.NameBand)
		fmt.Println("Email Band :", dataBandBaru.Email)
		fmt.Printf("Penanggung Jawab : %s\n\n", dataBandBaru.Captain)

		// shwowing band ========================
		band := GetBand(tx)
		ShowingDataBand(band)
		// ======================================
	} else {
		fmt.Println("Data band tidak ditemukan.")
	}
	// =======================================

	// commit transaction ====================
	err = tx.Commit()
	if err != nil {
		controller.Validate(err, "commit band", tx)
	}
	// =======================================
}

func UpdateBand(band entity.Band, tx *sql.Tx) entity.Band {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\nKosongkan jika tidak dirubah.")

	// Nama band
	fmt.Printf("Masukkan Nama Baru (lama:%s) :\n", band.NameBand)
	scanner.Scan()
	NameBand := scanner.Text()
	if NameBand != "" {
		band.NameBand = NameBand
	}
	ErrorInput(scanner.Err(), tx)

	// Email band
	fmt.Printf("Masukkan Email baru (lama:%s) :\n", band.Email)
	scanner.Scan()
	Email := scanner.Text()
	if Email != "" {
		band.Email = Email
	}
	ErrorInput(scanner.Err(), tx)

	// Penanggung jawab band
	fmt.Printf("Masukkan Penanggung Jawab baru (lama:%s) :\n", band.Captain)
	scanner.Scan()
	Captain := scanner.Text()
	if Captain != "" {
		band.Captain = Captain
	}
	ErrorInput(scanner.Err(), tx)

	// insert
	_, err := insertUpdateBand(tx, band.NameBand, band.Captain, band.Email, band.Id)
	controller.Validate(err, "Insert Band Detail", tx)

	return band
}

/*
* Update data band to table bands
 */
func insertUpdateBand(tx *sql.Tx, bandName, captain, email string, idBand int) (int, error) {
	insertSQL := `
			UPDATE 
				bands 
			SET 
				name_band = $2, 
				captain = $3, 
				email = $4
			WHERE id = $1;`
	var id int
	_, err := tx.Exec(insertSQL, idBand, bandName, captain, email)
	return id, err
}
