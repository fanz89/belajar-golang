package main

import "fmt"

func main() {

	// Sumber	=> https://dasarpemrogramangolang.novalagung.com/12-seleksi-kondisi.html

	// variable temporary => variable yang hanya bisa digunakan dalam blok seleksi kondisi dimana variable tersebut di deklarasikan
	/*
		Pemanfaatan variable temporary :
		- Menegaskan scope atau cakupan dari variable karena hanya bisa digunakan dalam blok seleksi kondisi
		- Kode program menjadi lebih rapih
		- Ketika nilai variable tersebut didapat dari sebuah komputasi, perhitungan tidak perlu dilakukan di dalam block masing-masing kondisi
	*/

	var point = 6840.0 // deklarasi variable point

	if percent := point / 100; percent >= 100 { // mendaklarasikan variable percent dengan di isi nilai hasil dari point / 100 lalu lakukan operasi pembanding apakah nilai percent > 100
		fmt.Printf("%.1f%s Perfect!\n", percent, "%") // mencetak ke layar nilai dari percent
	} else if percent >= 70 { // jika nilai percent lebih dari atau sama dengan 70
		fmt.Printf("%.1f%s Good!\n", percent, "%") // mencetak ke layar nilai dari percent
	} else { // jika kondisi di atas tidak terpenuhi
		fmt.Printf("%.1f%s Not Bad!\n", percent, "%") // mencetak ke layar nilai dari percent
	}

	// variable percent didapat dari hasil perhitungan dan hanya bisa digunakan di deretan block seleksi kondisi itu saja.
	// jika kode dibawah ini tidak diberi komentar akan terjadi error "undefined: percent"

	// fmt.Printf("%.1f%s \n", percent, "%")

	// Deklarasi variable temporary hanya bisa digunakan dengan metode type interface ( := )
	// Jika kode dibawah ini tidak diberi komentar akan terjadi error "var declaration not allowed in if initializer"

	// if var percent := point / 100 ; percent > 70 {}

}
