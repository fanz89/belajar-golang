package main

import "fmt"

func main() {

	// Sumber		=> https://dasarpemrogramangolang.novalagung.com/13-perulangan.html

	// Perulangan 	=> Proses mengulang-ulang eksekusi blok kode tanpa henti sampai kondisi yang ditentukan terpenuhi.
	// Di golang untuk perulangan hanya menggunakan keyword "for" saja tapi kemampuannya gabungan dari "for", "foreach" dan "while" dibahasa pemrograman lain.

	// Perulangan menggunakan keyword "for"

	for i := 0; i < 5; i++ { // melakukan perulangan dari i = 0 sampai dengan i < 5, dimana nilai i akan bertambah 1 ( i++ ) setiap melakukan perulangan

		fmt.Println("Angka : ", i) // cetak ke layar

	}
}
