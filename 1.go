package main

import "fmt"

// Buat array yang berisi murid-murid yang telah mengikuti ujian.
// Setiap murid memiliki nama, mata pelajaran, dan nilai.
// Minta input jumlah murid.
// Minta input nama murid, mata pelajaran, dan nilai dari user untuk setiap murid.
// Tampilkan semua nama murid, mata pelajaran, dan nilainya.

func main() {
	type murid struct {
		nama string
		mapel string
		nilai int
	}

	var list_murid [100]murid

	var jumlah int
	fmt.Print("Masukkan jumlah murid: ")
	fmt.Scanln(&jumlah)

	var i int

	for i = 0; i < jumlah; i++ {
		fmt.Print("Masukkan nama murid: ")
		fmt.Scanln(&list_murid[i].nama)
		fmt.Print("Masukkan mata pelajaran: ")
		fmt.Scanln(&list_murid[i].mapel)
		fmt.Print("Masukkan nilai: ")
		fmt.Scanln(&list_murid[i].nilai)
	}
	
	fmt.Println("\nNama\t\tMata Pelajaran\t\tNilai")
	fmt.Println("---------------------------------------------")
	for i = 0; i < jumlah; i++ {
		fmt.Println(list_murid[i].nama, "\t\t", list_murid[i].mapel, "\t\t", list_murid[i].nilai)
	}
}
