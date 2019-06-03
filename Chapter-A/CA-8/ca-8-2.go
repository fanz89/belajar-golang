package main

import "fmt"

var helloWorld1 string = "Hello World 1!!" // manifest typing

var helloWorld2 = "Hello World 2!!" // type interface menggunakan keyword var dan operand =

// helloWorld3 := "Hello World 3!!" => syntax error: non-declaration statement outside function body

func main() {

	// Sumber 			=> https://dasarpemrogramangolang.novalagung.com/8-variabel.html
	// Type interface 	=> metode deklarasi variable yang tipe datanya ditentukan oleh tipe data nilainya

	// Penggunaan type interface menggunakan operand := dan tidak menggunakan keyword var

	firstname1 := "Muhammad" // mendeklarasikan variable firstname1
	lastname1 := "Firdaus"   // mendeklarasikan variable lastname1

	fmt.Printf("Hallo %s %s!! \n", firstname1, lastname1) // mencatak ke layar

	// Penggunaan type interface bisa juga menggunakan keyword var
	// mengganti operand := dengan =

	var firstname2 = "Adi"   // mendeklarasikan variable firstname2
	var lastname2 = "Wibowo" // mendeklarasikan variable lastname2

	fmt.Printf("Hallo %s %s!! \n", firstname2, lastname2) // mencetak ke layar

	// Penggunaan operand := hanya bisa digunakan dalam block function

	fmt.Println(helloWorld1) // mencetak ke layar
	fmt.Println(helloWorld2) // mencetak ke layar

}
