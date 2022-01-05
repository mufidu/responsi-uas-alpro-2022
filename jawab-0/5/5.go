package main

import "fmt"

// 5. Cari murid berdasarkan nilainya.
// Gunakan binary search secara rekursif.

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

	selectionSort(&list_murid, j)

	fmt.Println("\nMasukkan nilai yang ingin dicari: ")
	var nilai int
	fmt.Scanln(&nilai)

	var hasil = binSearch(&list_murid, 0, j-1, nilai)
	fmt.Println("\nNama\t\tMata Pelajaran\t\tNilai")
	fmt.Println("---------------------------------------------")
	fmt.Println(hasil.nama, "\t\t", hasil.mapel, "\t\t\t", hasil.nilai)
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

func binSearch(list_murid *[100]murid, l int, r int, nilai int) murid {	
	if r >= l {
		var mid int
		mid = l + (r - l)/2

		if list_murid[mid].nilai == nilai {
			return list_murid[mid]
		} else if list_murid[mid].nilai > nilai {
			return binSearch(list_murid, l, mid-1, nilai)
		} else {
			return binSearch(list_murid, mid+1, r, nilai)
		}
	}
	return murid{}
}

func binSearchBiasa(list_murid [100]murid, n int, nilai int) int {
	var l int
	var r int
	var mid int

	l = 0
	r = n - 1

	for l <= r {
		mid = l + (r - l)/2

		if list_murid[mid].nilai == nilai {
			return mid
		} else if list_murid[mid].nilai > nilai {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
