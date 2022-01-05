package main

import "fmt"

// 3. Urutkan murid dari nilai tertinggi ke terendah.
// Gunakan 2 function dengan selection dan insertion sort.

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

	fmt.Println("\nSetelah diurutkan dari nilai tertinggi ke terendah: ")
	selectionSort(&list_murid, j)
	// insertionSort(&list_murid, j)

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

func selectionSort(list_murid *[100]murid, j int) {
	var i int
	var k int
	var temp murid

	for i = 0; i < j; i++ {
		for k = i + 1; k < j; k++ {
			if list_murid[i].nilai < list_murid[k].nilai {
				temp = list_murid[i]
				list_murid[i] = list_murid[k]
				list_murid[k] = temp
			}
		}
	}
}

func insertionSort(list_murid *[100]murid, n int) {
	var i int
	var j int

	for i = 1; i < n; i++ {
		j = i
		for j > 0 {
			if list_murid[j].nilai > list_murid[j-1].nilai {
				list_murid[j-1], list_murid[j] = list_murid[j], list_murid[j-1]
			}
			j = j - 1
		}
	}
}
