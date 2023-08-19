package controller

import (
	"database/sql"
	"fmt"
	"os"
)

func Validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		defer End()
		tx.Rollback()
		panic(err)
	} else {
		fmt.Println("Succesfully " + message + " data!")
	}
}
func End() {
	message := recover() // menangkap trow massage dari panic (jika tidak terjadi panic akan terisi nil)
	if message != nil {
		fmt.Println("Aplikasi tidak selesai dijalankan")
		fmt.Println("error dengan massage : ", message)
		os.Exit(0)
	}
}
