package main

import "fmt"

func main() {

	// Sumber 	=> https://dasarpemrogramangolang.novalagung.com/12-seleksi-kondisi.html

	// Switch di golang bisa diimplementasikan dengan style seperti if-else.
	// Implementasinya, nilai yang akan diseleksi atau dibandingkan tidak ditulis setelah keyword "switch"
	// tetapi ditulis langsung dalam bentuk perbandingan setelah keyword "case"

	var point = 6 // mendeklarasikan variable point dengan diinisialisasi nilai 6

	switch { // melakukan proses seleksi
	case point == 8: // jika nilai point sama dengan 8
		fmt.Println("Pefect!!") // cetak ke layar Perfect!!
	case (point < 8) && (point > 3):
		fmt.Println("Awesome!!") // cetak ke layar Awesome!!
	default: // jika kondisi diatas tidak terpenuhi
		{
			fmt.Println("Not Bad!!")           // cetak ke layar Not Bad!!
			fmt.Println("You Can Be Better!!") // cetak ke layar You Can Be Better!!
		}
	}
}
