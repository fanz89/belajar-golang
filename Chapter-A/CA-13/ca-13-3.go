package main

import "fmt"

func main() {

	// Sumber		=> https://dasarpemrogramangolang.novalagung.com/13-perulangan.html

	// Penggunaan Keyword "for" tanpa argumen
	// Cara ini ditulis tanpa menggunakan kondisi di baris "for"
	// Seleksi kondisi akan ditulis didalam blok "for" dan perulangan akan berhenti ketika menggunakan keyword "break"

	var i = 0 // mendeklarasikan variable i dan diinisalisasi dengan nilai 0

	for { // melakukan perulangan

		fmt.Println("Angka : ", i) // cetak ke layar

		i++ // nilai i ditambah 1

		if i == 5 { // jika nilai i sama dengan 5
			break // menghentikan perulangan
		}
	}

}
