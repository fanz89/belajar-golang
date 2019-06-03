package main

import "fmt"

func main() {

	// Sumber 	=> https://dasarpemrogramangolang.novalagung.com/9-tipe-data.html

	// Nilai nil Dan Nilai Default Tipe Data
	// Nil merupakan sebuah nilai atau berisi nilai kosong

	/*
		Tipe data yang bisa diisi nilai kosong atau nill
		- pointer
		- tipe data fungsi
		- slice
		- map
		- channel
		- interface kosong atau interface{}
	*/

	//	Nilai default dari tipe data yang pernah digunakan pada latihan sebelumnya :

	// string 	=> "" (string kosong)
	var hello string                   // mendeklarasikan variable hello dengan tipe string
	fmt.Printf("string : %s\n", hello) // mencetak ke layar variable hello

	// bool 	=> false
	var exist bool                   // mendeklarasikan variable exist dengan tipe data boolean
	fmt.Printf("bool : %t\n", exist) // mencetak ke layar variable exist

	// numerik atau non-desimal => 0
	var one uint                                 // mendeklarasikan variable one dengan tipe data uint
	var two int                                  // mendaklarasikan variable two dengan tipe data int
	fmt.Printf("one : %d, two : %d\n", one, two) // mencetak ke layar variable one dan two

	// numerik desimal => 0.0
	var flt32 float32                                        // mendeklarasikan variable flt32 dengan tipe data float32
	var flt64 float64                                        // mendeklarasikan variable flt64 dengan tipe data float64
	fmt.Printf("flt32 : %.1f, flt64 : %.1f\n", flt32, flt64) // mencetak ke layar variable flt32 dan flt64

}
