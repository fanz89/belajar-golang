package main

import "fmt"

func main() {

	// sumber 			=> https://dasarpemrogramangolang.novalagung.com/8-variabel.html
	// manifest typing 	=> mendaklarasikan variable wajib menggunakan tipe data
	// var 				=> keyword yang digunakan untuk mendeklarasikan variable baru
	// schema 			=>
	// 						var <nama-variabel> <tipe-data>
	// 						var <nama-variabel> <tipe-data> = <nilai>

	var firstname string = "Muhammad" // mendeklarasikan variable firstname dengan tipe data string
	var lastname string               // mendeklarasikan variable lastname
	lastname = "Firdaus"              // mengisi nilai variable lastname

	fmt.Printf("Halo %s %s!\n", firstname, lastname) // mencetak ke layar
}
