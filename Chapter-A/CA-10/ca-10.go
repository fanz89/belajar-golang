package main

import "fmt"

func main() {

	// Sumber 	=> https://dasarpemrogramangolang.novalagung.com/10-konstanta.html

	// Konstanta => jenis variable yang diinisialisasi hanya satu kali dan nilainya tidak bisa diubah

	// implementasi dengan manifest typing

	const firstname string = "Muhammad" // mendeklarasikan konstanta fistname dengan tipe data string
	fmt.Println("Hallo", firstname)     // mencetak variable fistname ke layar

	// firstname = "Adi" => jika tidak diberi komentar akan terjadi error "cannot assign to firstname"

	// implementasi dengan type interface

	const lastname = "Irfan Firdaus" // mendeklarasikan konstanta lastname
	fmt.Println("Hallo", lastname)   // mencetak ke layar

	// fmt.Print => ketika dicetak ke layar maka tidak akan menggunakan pemisah atau penghubung spasi

	fmt.Print(firstname, lastname, "\n") // mencetak ke layar
	fmt.Println(firstname, lastname)     // mencetak ke layar

}
