package main

import "fmt"

// 4. Outputkan rata-rata nilai murid.
// Buat function yang menggunakan metode recursive.

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

	fmt.Println("Rata rata nilai murid: ", rataRata(list_murid, j))
}

type murid struct {
	nama string
	mapel string
	nilai int
}

func rataRata(list_murid [100]murid, n int) float64 {
	var sum float64
	if n == 1 {
		sum = float64(list_murid[0].nilai)
	} else {
		// Hitung sum dari n-1 angka = (n-1) * (rata rata dari n-1 angka)
		// Tambahkan angka ke n ( i.e. list_murid[n-1])
		sum = float64(list_murid[n-1].nilai) + float64(n-1) * rataRata(list_murid, n-1)
	}
	return sum/float64(n)
}
