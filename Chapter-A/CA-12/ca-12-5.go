package main

import "fmt"

func main() {

	// Sumber 	=> https://dasarpemrogramangolang.novalagung.com/12-seleksi-kondisi.html

	// Tanda kurung kurawal {} bisa diimplementasikan pada keyword "case" dan default.
	// Penggunaanya bersifat optional, bisa digunakan atau tidak.
	// Baik digunakan ketika terdapat banyak statement agar kode terlihat lebih rapih.

	var point = 6 // mendeklarasikan variable point dengan diinisialisasi nilai 6

	switch point { // seleksi nilai point
	case 8: // jika nilai point sama dengan 8
		fmt.Println("Perfect") // cetak ke layar Perfect!!
	case 7, 6, 5, 4: // jika nilai point sama dengan 7 atau 6 atau 5 atau 4
		fmt.Println("Awesome!!") // cetak ke layar Awesome!!
	default: // jika kondisi diatas tidak terpenuhi
		{
			fmt.Println("Not Bad!!")           // cetak ke layar Not Bad!!
			fmt.Println("You Can Be Better!!") // cetak ke layar You Can Be Better!!
		}
	}

}
