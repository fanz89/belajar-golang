package main

import "fmt"

func main() {

	// Sumber	=> https://dasarpemrogramangolang.novalagung.com/9-tipe-data.html

	/* Tipe data numerik atau non-decimal secara umum terdapat 2 tipe data
	  	- uint	=> tipe data untuk bilangan cacah (bilangan positif)
		- int	=> tipe data untuk bilangan bulan (bilangan negatif dan positif)
	*/

	var positiveNumber uint8 = 86 // mendeklarasikan variable positiveNumber dengan tipe data uint8

	// Jika variable tidak secara ekplisit ditentukan tipe data nya maka compiler akan menentukan tipe data tersebut
	// Deklarasi dibawah akan bertipe int32 karena nilai tersebut termasuk ke dalam cakupan tipe data int32

	var negativeNumber = -1243423644 // mendeklarasikan variable negativeNumber

	// %d => untuk memformat data numerik non-desimal

	fmt.Printf("Bilangan positif : %d\n", positiveNumber) // mencetak ke layar
	fmt.Printf("Bilangan negatif : %d\n", negativeNumber) // mencetak ke layar

	// Dalam menentukan tipe data harus selektif atau sesuai kebutuhan data karena akan berkaitan dengan alokasi memori variable.
	// Pemilihan yang tepat akan membuat pemakaian memori lebih optimal dan tidak berlebihan.
}
