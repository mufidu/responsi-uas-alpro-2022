// Dari array berisi murid tadi, minta user memasukkan nama murid yang ingin dicari.
// Lalu tampilkan nama, mata pelajaran, dan nilai murid yang dicari.
// Jika tidak ditemukan, tampilkan pesan "Murid tidak ditemukan".
// Buat function yang menggunakan metode sequential search.

package main

import "fmt"

func main() {

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
	
	var nama_cari string
	fmt.Print("\nMasukkan nama murid yang ingin dicari: ")
	fmt.Scanln(&nama_cari)

	var hasil = seqSearch(nama_cari, list_murid, jumlah)

	if hasil.nilai != 0 {
		fmt.Println("\nNama\t\tMata Pelajaran\t\tNilai")
		fmt.Println("---------------------------------------------")
		fmt.Println(hasil.nama, "\t\t", hasil.mapel, "\t\t", hasil.nilai)
	}
}

type murid struct {
	nama string
	mapel string
	nilai int
}

func seqSearch(nama_cari string, list_murid [100]murid, jumlah int) murid {
	var i int
	for i = 0; i < jumlah; i++ {
		if list_murid[i].nama == nama_cari {
			return list_murid[i]
		}
	}
	fmt.Println("Murid tidak ditemukan")
	return murid{}
}
