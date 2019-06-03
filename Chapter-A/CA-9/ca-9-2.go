package main

import "fmt"

func main() {

	// Sumber 	=> https://dasarpemrogramangolang.novalagung.com/9-tipe-data.html

	/*
		Tipe Data Numerik Desimal terdapat 2 tipe
		- float32
		- float64
		Perbedaannya hanya pada cakupan nilai decimal yang bisa ditampung
	*/

	// Ketika tipe data tidak didefinisikan maka compiler akan menentukan tipe data dari nilai tersebut
	// Kode dibawah ini akan memiliki tipe data float32 karena nilai tersebut termasuk kedalam cakupan tipe data tersebut

	var decimalNumber = 2.62 // mendeklarasikan variable decimalNumber

	// %f => untuk memformat data numerik desimal menjadi string

	fmt.Printf("Bilangan desimal : %f\n", decimalNumber)

	// %.nf => memformat data numerik desimal menjadi string dengan menentukan digit yang muncul ( n )

	fmt.Printf("Bilangan desimal : %.3f\n", decimalNumber)

}
