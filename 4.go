// 4. Outputkan rata-rata nilai murid.
// Buat function yang menggunakan metode recursive.

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

	fmt.Println("Rata-rata nilai murid:", rataRata(list_murid, jumlah))
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
		sum = float64(list_murid[n-1].nilai) + float64(n-1) * rataRata(list_murid, n-1)
	}

	var rata float64
	rata = sum/float64(n)
	return rata
}
