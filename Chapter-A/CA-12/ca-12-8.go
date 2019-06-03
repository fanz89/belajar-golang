package main

import "fmt"

func main() {

	// Sumber 	=> https://dasarpemrogramangolang.novalagung.com/12-seleksi-kondisi.html

	// Seleksi kondisi bersarang => Suatu seleksi kondisi yang berada dalam seleksi kondisi dan seterusnya.
	// Implementasinya bisa dilakukan pada if-else, switch atau kombinasi dari keduanya.

	var point = 8 // mendeklarasikan variable point dengan diinisialisai nilai 8

	if point > 7 { // jika point lebih dari 7
		switch point { // seleksi nilai point
		case 10:                     // jika point sama dengan 10
			fmt.Println("Perfect!!") // cetak ke layar
		default:                  // jika kondisi case diatas tidak terpenuhi
			fmt.Println("Nice!!") // cetak ke layar
		}
	} else { // jika kondisi di atas tidak terpenuhi
		if point == 5 { // jika nilai point sama dengan 5
			fmt.Println("Not Bad!!") // cetak ke layar
		} else if point == 3 { // jika nilai point sama dengan 3
			fmt.Println("Keep Trying!!") // cetak ke layar
		} else { // jika kondisi diatas tidak terpenuhi
			fmt.Println("You Can Do It!!") // cetak ke layar
			if point == 0 { // jika nilai point sama dengan 0
				fmt.Println("Try Harder!!") // cetak ke layar
			}
		}
	}

}
