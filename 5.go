// 5. Cari murid berdasarkan nilainya.
// Gunakan binary search secara rekursif.

package main

import (
	"fmt"
)

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

	fmt.Print("\nMasukkan nilai yang ingin dicari: ")
	var nilai_cari int
	fmt.Scanln(&nilai_cari)

	selectionSort(&list_murid, jumlah)

	var hasil murid
	hasil = binSearch(&list_murid, 0, jumlah-1, nilai_cari)
	
	fmt.Println("\nNama\t\tMata Pelajaran\t\tNilai")
	fmt.Println("---------------------------------------------")
	fmt.Println(hasil.nama, "\t\t", hasil.mapel, "\t\t", hasil.nilai)
}

type murid struct {
	nama string
	mapel string
	nilai int
}

func selectionSort(list_murid *[100]murid, jumlah int) {
	var i int
	var j int
	var temp murid

	for i = 0; i < jumlah; i++ {
		for j = i; j < jumlah; j++ {
			if list_murid[i].nilai < list_murid[j].nilai {
				temp = list_murid[i]
				list_murid[i] = list_murid[j]
				list_murid[j] = temp
			}
		} 
	}
}

func binSearch(list_murid *[100]murid, l int, r int, nilai_cari int) murid {
	if r >= l {
		var mid int
		mid = l + (r-l) / 2

		if list_murid[mid].nilai == nilai_cari {
			return list_murid[mid]
		} else if list_murid[mid].nilai > nilai_cari {
			return binSearch(list_murid, l, mid-1, nilai_cari)
		} else {
			return binSearch(list_murid, mid+1, r, nilai_cari)
		}
	}
	return murid{}
}
