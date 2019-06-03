package main

import "fmt"

func main() {

	// Sumber		=> https://dasarpemrogramangolang.novalagung.com/13-perulangan.html

	// Penggunaan keyword "break" dan "continue"
	// Break 		=> memaksa menghentikan perulangan.
	// continue		=> memaksa maju ke perulangan selanjutnya.

	for i := 1; i < 10; i++ { // melakukan perulangan dari i == 1 sampai dengan i < 10, dimana nilai i akan bertambah 1 ( i++ ) setiap melakukan perulangan

		if i%2 == 1 { // jika hasil sisa bagi dari nilai i % 2 sama dengan 1
			continue // maju ke perulangan selanjutnya
		}

		if i > 8 { // jika nilai i lebih kecil dari 8
			break // menghentikan perulangan
		}

		fmt.Println("Angka : ", i) // cetak ke layar

	}

}
