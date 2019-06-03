package main

import "fmt"

func main() {

	// Sumber	=> https://dasarpemrogramangolang.novalagung.com/8-variabel.html

	// Deklarasi Variable dengan menggunakan keyword "new"
	// Keyword "new" digunakan untuk mencetak data "pointer" dengan tipe data tertentu
	// Nilai data default-nya akan menyesuaikan dengan tipe datanya

	name := new(string) // variable name menampung nilai pointer string

	fmt.Println(name) // ketika dicetak ke layar maka yang tampil adalah alamat memori dari nilai tersebut (dalam bentuk notasi heksadesimal)

	// Untuk menampilkan nilai dari pointer tersebut perlu di "dereference"
	// Proses dereference menggunakan tanda asterisk ( * )

	fmt.Println(*name) // mencetak ke layar nilai dari variable name

	// Dekalarasi variable dengan menggunakan keword "make"
	// Pemanfaatannya hanya bisa digunakan untuk "channel", "slice" dan "map"

}
