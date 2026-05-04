package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang [100]string
	var i int = 0
	var pertandingan int = 1

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	for {
		fmt.Print("Pertandingan ", pertandingan, " : ")
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			fmt.Println("Hasil", pertandingan, ":", klubA)
			pemenang[i] = klubA
			i++
		} else if skorB > skorA {
			fmt.Println("Hasil", pertandingan, ":", klubB)
			pemenang[i] = klubB
			i++
		} else {
			fmt.Println("Hasil", pertandingan, ": Draw")
		}

		pertandingan++
	}

	fmt.Println("Pertandingan selesai")
	fmt.Println("Daftar klub pemenang:")

	for j := 0; j < i; j++ {
		fmt.Println(j+1, ".", pemenang[j])
	}
}