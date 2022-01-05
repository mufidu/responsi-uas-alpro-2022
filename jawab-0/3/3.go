// 3. Urutkan murid dari nilai tertinggi ke terendah.
// Gunakan 2 function dengan selection dan insertion sort.

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
	
	fmt.Println("\nSetelah diurutkan dari nilai tertinggi ke terendah: ")
	insertionSort(&list_murid, jumlah)

	fmt.Println("\nNama\t\tMata Pelajaran\t\tNilai")
	fmt.Println("---------------------------------------------")
	for i = 0; i < jumlah; i++ {
		fmt.Println(list_murid[i].nama, "\t\t", list_murid[i].mapel, "\t\t", list_murid[i].nilai)
	}
}

type murid struct {
	nama string
	mapel string
	nilai int
}

func insertionSort(list_murid *[100]murid, n int) {
	var i int
	var j int

	for i = 1; i < n; i++ {
		for j > 0 {
			if list_murid[j].nilai > list_murid[j-1].nilai {
				list_murid[j], list_murid[j-1] = list_murid[j-1], list_murid[j]
			}
			j--
		}
	}
}
