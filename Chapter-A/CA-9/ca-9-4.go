package main

import "fmt"

func main() {

	// Sumber 	=> https://dasarpemrogramangolang.novalagung.com/9-tipe-data.html

	// Tipe data string memiliki ciri khas dengan diapit oleh tanda "quote" ( " )

	var messageOne string = "Hello World!!"

	fmt.Printf("Message : %s\n", messageOne)

	// Tipe data string bisa juga dideklarasikan dengan tanda "grave accent/backticks" ( ` )

	var messageTwo string = `Nama saya "Muhammad Irfan Firdaus".
	Salam kenal.
	Bismillah, mari belajar "Golang".`

	fmt.Println(messageTwo)

	// Jika menggunakan string backticks maka akan membuat semua karakter di dalam nya tidak di "escape" (semua akan terdeteksi string)
	// termasuk \n, tanda petik 2 ( " ), tanda petik satu ( ' ), baris baru dan lainnya
}
