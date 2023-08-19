package controller

import (
	"database/sql"
	"fmt"
	"tugas-akhir-enigmacamp-go/controller"
	"tugas-akhir-enigmacamp-go/database"
)

/*
* Ui of deleting data band
 */
func DeletingBandUi() {
	// sambungkan database ===================
	db := database.ConnectDb()
	defer db.Close()

	// buat transaksi
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// =======================================

	// shwowing band =========================
	band := GetBand(tx)
	ShowingDataBand(band)
	// =======================================

	// input id band wanna delete ============
	var idBand int
	fmt.Println("\nSilahkan masukkan Id band yang ingin dihapus :")
	_, err = fmt.Scan(&idBand)
	ErrorInput(err, tx)
	// =======================================

	// Get Data bend wanna delete ============
	deleteBand := GetBandById(idBand, tx)
	// =======================================

	// delete data ===========================
	DeleteInstrumentBand(idBand, tx)
	DeleteBand(idBand, tx)
	fmt.Printf("\n\nData band '%s' dengan ID %d telah dihapus\n\n", deleteBand.NameBand, idBand)
	// =======================================

	// shwowing band =========================
	band = GetBand(tx)
	ShowingDataBand(band)
	// =======================================

	// commit transaction ====================
	err = tx.Commit()
	if err != nil {
		controller.Validate(err, "commit band", tx)
	}
	// =======================================
}

/*
* Deleting data band refering ids
 */
func DeleteBand(ids int, tx *sql.Tx) {
	sqlStatement := "DELETE FROM bands WHERE id = $1;"

	_, err := tx.Exec(sqlStatement, ids)
	if err != nil {
		controller.Validate(err, "delete", tx)
		return
	}
}

/*
* Deleting data instrumen used in band refering bands id
 */
func DeleteInstrumentBand(ids int, tx *sql.Tx) {
	sqlStatement := "DELETE FROM band_instruments WHERE band_id = $1;"

	_, err := tx.Exec(sqlStatement, ids)
	if err != nil {
		controller.Validate(err, "delete", tx)
		return
	}
}
