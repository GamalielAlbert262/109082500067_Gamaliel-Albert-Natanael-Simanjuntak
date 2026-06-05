package main

import "fmt"

func main() {
	var suara [21]int
	var x int
	var suaraMasuk, suaraSah int

	for {
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		suaraMasuk++

		if x >= 1 && x <= 20 {
			suaraSah++
			suara[x]++
		}
	}

	fmt.Println("Suara masuk:", suaraMasuk)
	fmt.Println("Suara sah:", suaraSah)

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}