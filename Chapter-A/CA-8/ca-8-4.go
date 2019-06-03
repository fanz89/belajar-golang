package main

import "fmt"

func main() {

	// Sumber 	=> https://dasarpemrogramangolang.novalagung.com/8-variabel.html

	// Golang memiliki aturan jika variable yang sudah dideklrasikan wajib digunakan
	// jika tidak maka akan error => variable_name declared and not used

	// var hello string => jika tidak diberi komentar akan terjadi error karena variable hello tidak digunakan

	// Variable Underscore ( _ ) => predefined variable yang bisa dimanfaatkan untuk menampung nilai yang tidak terpakai.
	// 							 => Bisa dibilang variable ini merupakan keranjang sampah.

	_ = "Belajar Golang"              // nilai ditampung oleh variable underscore
	_ = "Bismillah, Golang itu mudah" // nilai ditampung oleh variable underscore

	// Variable underscore adalah predefined maka tidak menggunakan operand := cukup dengan =

	// _ := "asdf" => jika tidak diberi komentar akan terjadi error "no new variables on left side of :="

	// Variable underscore bisa menggunakan operand := atau type interface ketika akan mengimplentasikan multiple variable

	name, _ := "Adi", "Budi" // nilai "Adi" ditampung oleh variable name & nilai Budi ditampung oleh variable underscore

	// Variable underscore tidak dapat ditampilkan
	fmt.Println(name) // mencetak ke layar
	// fmt.Println(_)  jika tidak diberi komentar akan terjadi error "cannot use _ as value"

}
