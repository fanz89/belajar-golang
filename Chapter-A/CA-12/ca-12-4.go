package main

import "fmt"

func main() {

	// Sumber 	=> https://dasarpemrogramangolang.novalagung.com/12-seleksi-kondisi.html

	// Case bisa dimanfaatkan untuk menampung banyak kondisi
	// Implementasi nya dengan menuliskan pembanding-pembanding variable yang di-switch setelah keyword "case"]
	// Dipisah dengan tanda koma ( , ) => case 1,2,3,4

	var point = 2 // mendaklarasikan variable point dengan diinisialisasi nilai 2

	switch point { // menseleksi nilai point
	case 8: // jika nilai poin sama dengan 8
		fmt.Println("Perfect!!") // cetak ke layar Perfect!!
	case 7, 6, 5, 4: // jika nilai point sama dengan 8 atau 6 atau 5 atau 4
		fmt.Println("Awesome!!") // cetak ke layar Awesome!!
	default: // jika kondisi di atas tidak terpenuhi
		fmt.Println("Not Bad!!") // cetak ke layar Not Bad!!
	}

}
