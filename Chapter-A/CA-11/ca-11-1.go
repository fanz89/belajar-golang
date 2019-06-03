package main

import "fmt"

func main() {

	// Sumber 	=> https://dasarpemrogramangolang.novalagung.com/11-operator.html

	// Operator aritmatika => operator yang digunakan untuk operasi yang sifatnya perhitungan

	/*
		Operator standar yang didukung oleh golang
		+ => penjumahan
		- => pengurangan
		* => perkalian
		/ => pembagian
		% => modulus / sisa hasil pembagian
	*/

	var value = (((2+6)%3)*4 - 2) / 3 // mendeklarasikan variable value dengan diisi nilai dari hasil perhitungan
	fmt.Println("Result : ", value)   // mencetak ke layar variable value

}
