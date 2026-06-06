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
Program pilkart digunakan untuk menghitung hasil pemilihan ketua RT. Program membaca data suara hingga ditemukan angka 0, memvalidasi suara (hanya nomor calon 1–20 yang sah), menghitung jumlah suara masuk dan suara sah, lalu menampilkan perolehan suara setiap calon yang memperoleh suara.

Program menerapkan konsep array, perulangan, percabangan, dan validasi data untuk melakukan rekapitulasi suara secara otomatis.



### 2. [Soal]
#### soal2.go

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

	ketua := 1
	for i := 2; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			ketua = i
		}
	}

	wakil := -1
	for i := 1; i <= 20; i++ {
		if i != ketua {
			if wakil == -1 || suara[i] > suara[wakil] {
				wakil = i
			}
		}
	}

	fmt.Println("Suara masuk:", suaraMasuk)
	fmt.Println("Suara sah:", suaraSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/GamalielAlbert262/109082500067_Gamaliel-Albert-Natanael-Simanjuntak/blob/main/modull12/outputscreenshot/soal2.png)

#### [penjelasan]
Program membaca data suara sampai angka 0, menghitung suara sah (nomor 1–20), lalu mencari calon dengan suara terbanyak sebagai Ketua RT dan calon dengan suara terbanyak kedua sebagai Wakil Ketua RT. Jika jumlah suara sama, calon dengan nomor lebih kecil akan dipilih terlebih dahulu.



### 3. [Soal]
#### soal3.go

```go
package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func main() {
	var n, k int

	fmt.Scan(&n, &k)

	isiArray(n)

	pos := posisi(n, k)

	if pos == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(pos)
	}
}

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	kr := 0
	kn := n - 1

	for kr <= kn {
		med := (kr + kn) / 2

		if data[med] == k {
			return med
		} else if data[med] < k {
			kr = med + 1
		} else {
			kn = med - 1
		}
	}

	return -1
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/GamalielAlbert262/109082500067_Gamaliel-Albert-Natanael-Simanjuntak/blob/main/modull12/outputscreenshot/soal3.png)

#### [penjelasan]
Program membaca jumlah data n dan nilai yang dicari k, kemudian mengisi array yang sudah terurut membesar. Pencarian dilakukan menggunakan algoritma Binary Search dengan membandingkan k terhadap elemen tengah array. Jika ditemukan, fungsi mengembalikan indeksnya; jika tidak ditemukan, fungsi mengembalikan -1 dan program menampilkan "TIDAK ADA".
