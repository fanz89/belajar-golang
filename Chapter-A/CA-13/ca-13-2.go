package main

import "fmt"

func main() {

	// Sumber		=> https://dasarpemrogramangolang.novalagung.com/13-perulangan.html

	// Penggunaan keyword "for" dengan argument hanya kondisi
	// Implementasinya seperti while pada pemrograman lain

	var i = 0 // mendeklarasikan variable i dengan diinisialisasi nilai 0

	for i < 5 { // melakukan perulangan sampai nilai i lebih kecil dari 5

		fmt.Println("Angka : ", i) // cetak ke layar
		i++                        // nilai i ditambah 1

	}

}
