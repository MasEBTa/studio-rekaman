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

/*
* Ui of add new data band
 */
func AddBand() {
	// sambungkan database ===================
	db := database.ConnectDb()
	defer db.Close()
	// buat transaksi
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	// =======================================
	// ambil data instrument =================
	instruments := GetInstrument(tx)

	// masukkan detail band ==================
	band, bandId := InputBandDetail(tx)

	// masukkan instrument yang digunakan ====
	instrumentsInBand := InputInstrument(bandId, instruments, tx)

	// commit transaction
	err = tx.Commit()
	if err != nil {
		controller.Validate(err, "commit band", tx)
	}

	fmt.Printf("\n\n==============\n\n")
	fmt.Println("Data berhasil ditambahkan :")
	fmt.Println("*) Nama Band :", band.NameBand)
	fmt.Println("*) Email Band :", band.Email)
	fmt.Println("*) Penanggung Jawab :", band.Captain)
	fmt.Println("*) Daftar Instrumen :")
	for _, v := range instruments {
		for _, i := range instrumentsInBand {
			if v.Id == i.InstrumentTypeId {
				fmt.Println("   -", i.Total, v.Instruments)
			}
		}
	}
	fmt.Printf("\n==============\n\n\n")
}

/*
* Getting data from table instruments
 */
func GetInstrument(tx *sql.Tx) (result []entity.Instrument) {
	sqlStatement := "SELECT id, instrumen_name FROM instruments;"

	rows, err := tx.Query(sqlStatement)
	if err != nil {
		controller.Validate(err, "select", tx)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var instrument entity.Instrument
		if err := rows.Scan(&instrument.Id, &instrument.Instruments); err != nil {
			controller.Validate(err, "scan", tx)
			continue
		}
		result = append(result, instrument)
	}

	if err := rows.Err(); err != nil {
		controller.Validate(err, "rows", tx)
	}

	return result
}

/**
* Getting input of data band
 */
func InputBandDetail(tx *sql.Tx) (entity.Band, int) {
	// variable
	var band entity.Band
	scanner := bufio.NewScanner(os.Stdin)

	// name band
	fmt.Println("Masukkan Nama Band :")
	scanner.Scan()
	band.NameBand = scanner.Text()
	ErrorInput(scanner.Err(), tx)

	// email band
	fmt.Println("Masukkan Email :")
	scanner.Scan()
	band.Email = scanner.Text()
	ErrorInput(scanner.Err(), tx)

	// captain of band
	fmt.Println("Masukkan Penanggung Jawab :")
	scanner.Scan()
	band.Captain = scanner.Text()
	ErrorInput(scanner.Err(), tx)

	// insert
	BandId, err := insertBandAndGetID(tx, band.NameBand, band.Captain, band.Email)
	controller.Validate(err, "Insert Band Detail", tx)

	return band, BandId
}

/*
* Insert data band to table bands
 */
func insertBandAndGetID(tx *sql.Tx, bandName, captain, email string) (int, error) {
	insertSQL := "INSERT INTO bands (name_band, captain, email) VALUES ($1, $2, $3) RETURNING id;"
	var id int
	err := tx.QueryRow(insertSQL, bandName, captain, email).Scan(&id)
	return id, err
}

/**
* Getting input instrument to used in band
 */
func InputInstrument(bandId int, instrumen []entity.Instrument, tx *sql.Tx) []entity.BandInstruments {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nDaftar Instrument :")
	for _, instrument := range instrumen {
		fmt.Printf("%d. %s\n", instrument.Id, instrument.Instruments)
	}
	fmt.Println("Tuliskan instrumen yang digunakna (jika menggunakan drum, gitar dan bass tulis : 1 2 3) :")
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

	// jumlah instrumen
	fmt.Println("\n\nMasukkan jumlah instrumen yang dibutuhkan.")
	var instrumentsOfBand []entity.BandInstruments
	for _, id := range selectedInstruments {
		for _, instrument := range instrumen {
			if id == instrument.Id {
				var jml int
				fmt.Printf("Jumlah %s:\n", instrument.Instruments)
				_, err := fmt.Scan(&jml)

				ErrorInput(err, tx)

				instrumentOfBand, err := insertBandInstrument(tx, bandId, id, jml)
				ErrorInput(err, tx)

				instrumentsOfBand = append(instrumentsOfBand, instrumentOfBand)
			}
		}
	}

	return instrumentsOfBand
}

/*
* Insserting each Instrument used in the band to band_instruments table
 */
func insertBandInstrument(tx *sql.Tx, bandId, instrumentTypeId, total int) (entity.BandInstruments, error) {
	insertSQL := "INSERT INTO band_instruments (band_id, instrument_type_id, total) VALUES ($1, $2, $3) RETURNING id;"
	var insertedID int // Ini adalah variabel untuk menyimpan ID yang dihasilkan dari RETURNING clause
	err := tx.QueryRow(insertSQL, bandId, instrumentTypeId, total).Scan(&insertedID)
	if err != nil {
		return entity.BandInstruments{}, err
	}

	// Membuat objek BandInstruments baru dengan ID yang dihasilkan
	newBandInstrument := entity.BandInstruments{
		Id:               insertedID,
		BandId:           bandId,
		InstrumentTypeId: instrumentTypeId,
		Total:            total,
	}

	return newBandInstrument, nil
}

/**
* Throwing error if you scanning input from user
 */
func ErrorInput(err error, tx *sql.Tx) {
	if err != nil {
		defer controller.End()
		tx.Rollback()
		panic(err)
	}
}
