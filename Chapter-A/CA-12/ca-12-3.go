package main

import "fmt"

func main() {

	// Sumber 	=> https://dasarpemrogramangolang.novalagung.com/12-seleksi-kondisi.html

	// Switch 	=> seleksi kondisi yang berfokus pada satu variable, kemudian dicek nilainya.
	// Switch di golang ketika case terpenuhi maka tidak akan dilanjutkan pada case atau kondisi selanjutnya
	// Jadi tidak perlu menggunakan keyword "break" seperti bahasa pemrogram lainnya, misalkan java.

	var point = 6 // mendeklarsikan variable point dengan diinisialisasi nilai 6

	switch point { // menseleksi nilai point
	case 8: // jika nilai point sama dengan 8
		fmt.Println("Perfect!!") // cetak ke layar Perfect!!
	case 7: // jika nilai point sama dengan 7
		fmt.Println("Awesome!!") // cetak ke layar Awesome!!
	default: // jika case atau kondisi di atas tidak terpenuhi
		fmt.Println("Not Bad!!") // cetak ke layar Not Bad!!
	}

}
