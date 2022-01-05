package main

import "fmt"

// 2. Dari array berisi murid tadi, minta user memasukkan nama murid yang ingin dicari.
// Lalu tampilkan nama dan mata pelajaran murid yang dicari.
// Jika tidak ditemukan, tampilkan pesan "Murid tidak ditemukan".
// Buat function yang menggunakan metode sequential search.

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

	fmt.Println("\nMasukkan nama murid yang ingin dicari: ")
	var nama_cari string
	fmt.Scanln(&nama_cari)
	var hasil = sequentialSearch(nama_cari, list_murid, j)
	fmt.Println("Nama\t\tMata Pelajaran\t\tNilai")
	fmt.Println("---------------------------------------------")
	fmt.Println(hasil.nama, "\t\t", hasil.mapel, "\t\t\t", hasil.nilai)
}

type murid struct {
	nama string
	mapel string
	nilai int
}

func sequentialSearch(nama_cari string, list_murid [100]murid, j int) murid {
	var i int
	for i = 0; i < j; i++ {
		if nama_cari == list_murid[i].nama {
			return list_murid[i]
		}
	}
	fmt.Println("Murid tidak ditemukan")
	return murid{}
}
