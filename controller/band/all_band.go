package controller

import (
	"database/sql"
	"fmt"
	"tugas-akhir-enigmacamp-go/controller"
	"tugas-akhir-enigmacamp-go/database"
	"tugas-akhir-enigmacamp-go/entity"
)

/*
* Ui of menu Showing data band
 */
func ShowingBand() {
	// sambungkan database ===================
	db := database.ConnectDb()
	defer db.Close()

	// buat transaksi
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// =======================================
	band := GetBand(tx)
	ShowingDataBand(band)
	// =======================================

	var idBand int
	fmt.Printf("Masukkan Id Band untuk melihat daftar Instrument yang digunakan :\n")
	_, err = fmt.Scan(&idBand)
	ErrorInput(err, tx)

	instrumentOfBand := GetInstrumentOfBand(idBand, tx)
	for _, v := range band {
		if v.Id == idBand {
			fmt.Println("\n======================")
			fmt.Println("Nama Band :", v.NameBand)
		}
	}
	fmt.Println("Instrumen Yang Digunakan :")

	if len(instrumentOfBand) > 0 {
		for _, v := range instrumentOfBand {
			fmt.Println(v.Total, v.Name)
		}
	} else {
		fmt.Println("Instrument tidak ditemukan.")
	}
	fmt.Println("======================")
	// commit transaction
	err = tx.Commit()
	if err != nil {
		controller.Validate(err, "commit band", tx)
	}
}

/*
* Getting All data from Bands Table
 */
func GetBand(tx *sql.Tx) (result []entity.Band) {
	sqlStatement := "SELECT id, name_band, captain, email FROM bands;"

	rows, err := tx.Query(sqlStatement)
	if err != nil {
		controller.Validate(err, "select", tx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var band entity.Band
		if err := rows.Scan(&band.Id, &band.NameBand, &band.Captain, &band.Email); err != nil {
			controller.Validate(err, "scan", tx)
			continue
		}
		result = append(result, band)
	}

	if err := rows.Err(); err != nil {
		controller.Validate(err, "rows", tx)
	}

	return result
}

/*
* Getting datum from Bands Table feference those id's
 */
func GetBandById(ids int, tx *sql.Tx) (band entity.Band) {
	sqlStatment := "SELECT id, name_band, captain, email FROM bands WHERE id = $1"

	rows, err := tx.Query(sqlStatment, ids)
	if err != nil {
		controller.Validate(err, "select", tx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&band.Id, &band.NameBand, &band.Captain, &band.Email); err != nil {
			controller.Validate(err, "scan", tx)
			continue
		}
	}

	if err := rows.Err(); err != nil {
		controller.Validate(err, "rows", tx)
	}

	return band
}

/*
* Getting instrumens which use in the band
* select the band using ids
 */
func GetInstrumentOfBand(idBand int, tx *sql.Tx) (bandInstruments []entity.InstrumentAndTotal) {
	sqlStatement := `
	SELECT
		bi.total,
		i.instrumen_name
	FROM
		band_instruments AS bi
	LEFT JOIN
		instruments AS i ON bi.instrument_type_id = i.id
	WHERE
		bi.band_id = $1;
	`

	rows, err := tx.Query(sqlStatement, idBand)
	if err != nil {
		controller.Validate(err, "select instrument", tx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var instrumenBand entity.InstrumentAndTotal
		if err := rows.Scan(&instrumenBand.Total, &instrumenBand.Name); err != nil {
			controller.Validate(err, "Select Instrument", tx)
			continue
		}
		bandInstruments = append(bandInstruments, instrumenBand)
	}

	if err := rows.Err(); err != nil {
		controller.Validate(err, "rows", tx)
	}

	return bandInstruments
}

/*
* Showing data bands table in terminal
 */
func ShowingDataBand(band []entity.Band) {
	fmt.Println("Daftar Band yang terdaftar ")
	fmt.Println("===========================")
	fmt.Println("No | Id | nama | email | penanggung jawab")
	for i, v := range band {
		fmt.Printf("%d. | %d | %s | %s | %s\n", i+1, v.Id, v.NameBand, v.Email, v.Captain)
	}
	fmt.Println("===========================")
}
