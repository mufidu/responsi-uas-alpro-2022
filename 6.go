// 6. Buat function selection sort dengan metode rekursif.

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

	fmt.Println("\nSetelah diurutkan dari nilai terendah ke tertinggi: ")
	selSort(&list_murid, jumlah, 0)

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

func selSort(list_murid *[100]murid, jumlah int, i int) {
	var j int

	if i < jumlah {
		for j = i + 1; j < jumlah; j++ {
			if list_murid[j].nilai < list_murid[i].nilai {
				list_murid[i], list_murid[j] = list_murid[j], list_murid[i]
			}
		}
		i++
		selSort(list_murid, jumlah, i)
	}
}
