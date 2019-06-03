package main

import "fmt"

func main() {

	// Sumber 	=> https://dasarpemrogramangolang.novalagung.com/11-operator.html

	// Operator logika 	=> mencari hasil benar atau tidak nya kombinasi dari nilai yang bertipe boolean
	//					=> ataupun hasil dari operator pembanding
	// Hasil dari operator logika sama dengan hasil dari operator pembanding yaitu nilai bertipe bolean

	/*
		Beberapa operator yang sering digunakan di golang
		&&	=> kiri dan kanan
		|| 	=> kiri atau kanan
		!	=> negasi / nilai kebalikan
	*/

	var left = false // mendeklarasikan variable left dengan diisi nilai bertipe boolean
	var right = true // mendaklarasikan variable right dengan diisi nilai bertipe boolean

	// Hasil code dibawah ini akan bernilai false
	// karena false && true = false

	var leftAndRight = left && right                    // mendeklarasikan variable leftAndRight dengan diisi nilai dari hasil operator logika
	fmt.Printf("left && right \t(%t) \n", leftAndRight) // mencetak ke layar nilai dari variable leftAndRight

	// Hasil code dibawah ini akan bernilai true
	// karena false || true = true

	var leftOrRight = left || right                    // mendeklarasikan variable leftOrRight dengan diisi nilai dari hasil operator logika
	fmt.Printf("left || right \t(%t) \n", leftOrRight) // mencetak ke layar nilai dari variable leftOrRight

	// Hasil code dibawah ini akan bernilai true
	// karena negasi atau nilai kebalikan dari false adalah true

	var leftReverse = !left                      // mendaklarasikan variable leftReverse dengan diisi nilai dari hasil operator logika
	fmt.Printf("!left \t\t(%t) \n", leftReverse) // mencetak ke layar nilai dari variable leftReverse

	// \t => untuk menambahkan "indent tabulasi"
}
