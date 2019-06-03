package main

import "fmt"

func main() {

	// Sumber		=> https://dasarpemrogramangolang.novalagung.com/13-perulangan.html

	// Perulangan bersarang

	for i := 0; i < 5; i++ { // melakukan perulangan dari i == 0 sampai dengan i < 5, dimana nilai i akan bertambah 1 ( i++ ) setiap melakukan perulangan

		for j := i; j < 5; j++ { // melakukan perulangan dari j == i sampai dengan j < 5, dimana nilai j akan bertambah 1 ( j++ ) setiap melakukan perulangan

			fmt.Print(j, " ") // cetak ke layar nilai j

		}

		fmt.Println() // menambahkan baris baru

	}

}
