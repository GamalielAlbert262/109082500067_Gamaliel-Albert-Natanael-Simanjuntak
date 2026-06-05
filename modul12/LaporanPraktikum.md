# <h1 align="center">Laporan Praktikumm Modul 12  - SEARCHING </h1>
<p align="center">[Gamaliel Albert Natanael Simanjuntak] - [109082500067]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/GamalielAlbert262/109082500067_Gamaliel-Albert-Natanael-Simanjuntak/blob/main/modul12/outputscreenshot/soal1.png)

#### [penjelasan]
