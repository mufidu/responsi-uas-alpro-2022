package main

import "fmt"

// 1. Buat array yang berisi murid-murid yang telah mengikuti ujian.
// Setiap murid memiliki nama, mata pelajaran, dan nilai.
// Minta input jumlah murid.
// Minta input nama murid dan mata pelajaran dari user.
// Tampilkan semua nama murid, mata pelajaran, dan nilainya.

func main() {
	var list_murid [100]murid
	var i int
	var j int

	fmt.Println("Masukkan jumlah murid: ")	
	fmt.Scanln(&j)	
	for i = 0; i < j; i++ {	
		fmt.Println("Masukkan nama murid: ")	
		fmt.Scanln(&list_murid[i].nama)	
		fmt.Println("Masukkan mata pelajaran: ")	
		fmt.Scanln(&list_murid[i].mapel)
		fmt.Println("Masukkan nilai: ")
		fmt.Scanln(&list_murid[i].nilai)	
	}	

	fmt.Println("\nNama\t\tMata Pelajaran\t\tNilai")
	fmt.Println("---------------------------------------------")
	for i = 0; i < j; i++ {
		fmt.Println(list_murid[i].nama, "\t\t", list_murid[i].mapel, "\t\t\t", list_murid[i].nilai)
	}
}

type murid struct {
	nama string
	mapel string
	nilai int
}