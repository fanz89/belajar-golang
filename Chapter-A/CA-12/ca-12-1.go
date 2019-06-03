package main

import "fmt"

func main() {

	// Sumber	=> https://dasarpemrogramangolang.novalagung.com/12-seleksi-kondisi.html

	// Seleksi kondisi => digunakan untuk mengatur alur program.
	// Yang menjadi acuan yaitu nilai bertipe boolean yang bisa berasal dari variable ataupun hasil operasi pembanding.
	// Hasil dari nilai yang diseleksi akan mengacu kepada blok kode mana yang nantinya akan dieksekusi.
	/*
		Golang memiliki 2 keyword untuk seleksi kondisi
		- if else
		- switch
		Golang tidak mendukung seleksi kondisi dengan menggunakan ternary
	*/

	// Seleksi kondisi menggunakan if, else if dan else
	// Penerapan di golang "parenteheses" atau tanda kurung () tidak perlu dituliskan

	var point = 8 // mendeklarasikan variable point dengan diinisalisasi nilai 8

	if point == 10 { // jika nilai point sama dengan nilai 10
		fmt.Println("Lulus dengan nilai sempurna") // cetak ke layar
	} else if point > 5 { // jika nilai point lebih besar dari 5
		fmt.Println("Lulus") // cetak ke layar
	} else if point == 4 { // jika nilai point sama dengan 4
		fmt.Println("Hampir lulus") // cetak ke layar
	} else { // jika kondisi semua kondisi diatas tidak terpenuhi
		fmt.Printf("Tidak lulus. Nilai anda %d\n", point) // cetak ke layar
	}

	// Penerapan di golang jika kondisi if else hanya memiliki 1 block statment wajib menggunakan kurung kurawal {}

	if point > 5 { // jika nilai point > 5
		fmt.Printf("%d > 5\n", point) // cetak ke layar
	}

	// Kode dibawah ini jika tidak diberi komentar akan terdapat error "unexpected newline, expecting { after if clause"

	// if point == 10 fmt.Printf("%d > 5\n", point)

}
