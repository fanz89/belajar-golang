package main

import "fmt"

func main() {

	// Sumber 	=> https://dasarpemrogramangolang.novalagung.com/11-operator.html

	// Operator pembanding => menentukan kebenaran dari suatu kondisi dan hasilnya berupa nilai boolean (true / false)
	/*
		Operator pembanding yang digunakan di golang
		== 	=> apakah nilai kiri sama dengan nilai kanan
		!=	=> apakah nilai kiri tidak sama dengan nilai kanan
		< 	=> apakah nilai kiri lebih kecil daripada nilai kanan
		<=	=> apakah nilai kiri lebih kecil atau sama dengan nilai kanan
		>	=> apakah nilai kiri lebih besar daripada nilai kanan
		>=	=> apakah nilai kiri lebih besar atau sama dengan nilai kanan
	*/

	var value = (((2+6)%3)*4 - 2) / 3 // mendekarasikan variable value dengan diisi nilai dari hasil perhitungan
	var isEqual = (value == 2)        // mendeklarasikan variable isEqual dengan diisi nilai dari hasil pembanding (value == 2)
	fmt.Printf("nilai %d ( %t ) \n", value, isEqual)

}
