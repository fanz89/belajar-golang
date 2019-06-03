package main

import "fmt"

func main() {

	// Sumber		=> https://dasarpemrogramangolang.novalagung.com/13-perulangan.html

	// Pemanfaatan label dalam perulangan
	/*
		Pada perulangan bersarang break dan continue akan berlaku pada blok dimana ia digunakan saja.
		Pemberian "label" dimanfaatkan pada break dan continue untuk menunju pada perulangan terluar atau perulangan tertentu saja.
	*/

outerloop: // mendeklarasikan label outerloop

	for i := 0; i < 5; i++ { // melakukan perulangan dari i == 0 sampai dengan i < 5, dimana nilai i akan bertambah 1 ( i++ ) setiap melakukan perulangan

		for j := 0; j < 5; j++ { // melakukan perulangan dari j == 0 sampai dengan j < 5, dimana nilai j akan bertambah 1 ( j++ ) setiap melakukan perulangan

			if i == 3 { // jika nilai i sama dengan 3
				break outerloop // menghentian perulangan dan menuju pada blok for sesuai dengan label outerloop dideklarasikan
			}

			fmt.Print("matriks [", i, "][", j, "]", "\n") // mencetak ke layar

		}

	}

}
