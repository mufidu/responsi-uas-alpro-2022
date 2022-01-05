package main

import "fmt"

// 6. Buat function selection sort dengan metode rekursif

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
	selectionSortRecursive(&list_murid, j, 0)

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

func selectionSortRecursive(list_murid *[100]murid, j int, i int) {
	var k int

	if i < j {
		for k = i + 1; k < j; k++ {
			if list_murid[i].nilai < list_murid[k].nilai {
				list_murid[i], list_murid[k] = list_murid[k], list_murid[i]
			}
		}
		i++
		selectionSortRecursive(list_murid, j, i)
	}
}

// func selectionSort(list_murid *[100]murid, j int) {
// 	var i int
// 	var k int
// 	var temp murid

// 	for i = 0; i < j; i++ {
// 		for k = i + 1; k < j; k++ {
// 			if list_murid[i].nilai < list_murid[k].nilai {
// 				temp = list_murid[i]
// 				list_murid[i] = list_murid[k]
// 				list_murid[k] = temp
// 			}
// 		}
// 	}
// }
