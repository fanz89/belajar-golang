package main

import "fmt"

func main() {

	// Sumber 	=> https://dasarpemrogramangolang.novalagung.com/8-variabel.html

	// Deklarasi multiple variable dengan menggunakan tanda ( , )
	// Implementasi dengan manifest typing

	var first, second, third string              // mendeklarasikan 3 variable dengan tipe data string
	first, second, third = "satu", "dua", "tiga" // mengisi 3 variable tersebut

	fmt.Printf("%s %s %s \n", first, second, third) // mencetak ke layar

	var fourth, fifth, sixth string = "empat", "lima", "enam" // mendeklarasikan 3 variabel dan langsung mengisi nilainya

	fmt.Printf("%s %s %s \n", fourth, fifth, sixth) // mencetak ke layar

	// Implementasi dengan type interface

	seventh, eight, ninth := "tujuh", "delapan", "sembilan" // mendeklarasikan 3 variable dan langsung mengisi nilainya

	fmt.Printf("%s %s %s \n", seventh, eight, ninth) // mencetak ke layar

	// Bisa mengimplementasi kan dengan tipe data yang berbeda
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	// %v the value in a default format (General)
	// %t the word true or false (Boolean)
	// %s the uninterpreted bytes of the string or slice (String and slice of bytes )

	fmt.Printf("%v %t %v %s \n", one, isFriday, twoPointTwo, say) // mencetak ke layar
}
