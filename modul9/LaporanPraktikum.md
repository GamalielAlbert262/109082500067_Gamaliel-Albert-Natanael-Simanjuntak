# <h1 align="center">Laporan Praktikum Modul 9  - ARRAY </h1>
<p align="center">[Gamaliel Albert Natanael Simanjuntak] - [109082500067]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat Titik
	r     int
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

func diDalam(c Lingkaran, p Titik) bool {
	return jarak(c.pusat, p) <= float64(c.r)
}

func main() {
	var c1x, c1y, r1 int
	var c2x, c2y, r2 int
	var px, py int

	fmt.Scan(&c1x, &c1y, &r1)
	fmt.Scan(&c2x, &c2y, &r2)
	fmt.Scan(&px, &py)

	ling1 := Lingkaran{Titik{c1x, c1y}, r1}
	ling2 := Lingkaran{Titik{c2x, c2y}, r2}
	titik := Titik{px, py}

	dalam1 := diDalam(ling1, titik)
	dalam2 := diDalam(ling2, titik)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/GamalielAlbert262/109082500067_Gamaliel-Albert-Natanael-Simanjuntak/blob/main/modul9/outputscreenshot/soal1.png)

#### [penjelasan]
Program ini dibuat untuk menentukan posisi sebuah titik terhadap dua lingkaran pada bidang koordinat. Setiap lingkaran didefinisikan oleh titik pusat (cx,cy)(cx, cy)(cx,cy) dan jari-jari rrr, sedangkan titik yang diuji memiliki koordinat (x,y)(x, y)(x,y).
Proses utama dalam program ini adalah menghitung jarak antara titik dengan pusat masing-masing lingkaran menggunakan rumus jarak Euclidean:
d=(x−cx)2+(y−cy)2d = \sqrt{(x - cx)^2 + (y - cy)^2}d=(x−cx)2+(y−cy)2​
Selanjutnya, jarak tersebut dibandingkan dengan jari-jari lingkaran. Jika jarak kurang dari atau sama dengan jari-jari, maka titik berada di dalam lingkaran. Sebaliknya, jika lebih besar, maka titik berada di luar lingkaran.
Program kemudian melakukan pengecekan terhadap kedua lingkaran dan menghasilkan keluaran berupa informasi posisi titik, yaitu:


Titik berada di dalam lingkaran 1 dan 2


Titik berada di dalam lingkaran 1 saja


Titik berada di dalam lingkaran 2 saja


Titik berada di luar kedua lingkaran


Dengan demikian, program ini mengimplementasikan konsep struktur data, fungsi, operasi matematika, dan percabangan untuk menyelesaikan permasalahan geometri sederhana.



### 2. [Soal]
#### soal2.go

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Isi array:", arr)

	fmt.Print("Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("Masukkan x: ")
	fmt.Scan(&x)

	fmt.Print("Indeks kelipatan x: ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var idx int
	fmt.Print("Masukkan indeks yang dihapus: ")
	fmt.Scan(&idx)

	arr = append(arr[:idx], arr[idx+1:]...)
	fmt.Println("Array setelah dihapus:", arr)

	sum := 0
	for _, v := range arr {
		sum += v
	}
	mean := float64(sum) / float64(len(arr))
	fmt.Println("Rata-rata:", mean)

	var variance float64
	for _, v := range arr {
		variance += math.Pow(float64(v)-mean, 2)
	}
	variance /= float64(len(arr))
	stdDev := math.Sqrt(variance)
	fmt.Println("Standar deviasi:", stdDev)

	var target, count int
	fmt.Print("Masukkan angka yang dicari: ")
	fmt.Scan(&target)

	for _, v := range arr {
		if v == target {
			count++
		}
	}
	fmt.Println("Frekuensi:", count)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/GamalielAlbert262/109082500067_Gamaliel-Albert-Natanael-Simanjuntak/blob/main/modul9/outputscreenshot/soal2.png)

#### [penjelasan]
Program ini digunakan untuk mengolah data bilangan bulat dalam sebuah array. Pengguna memasukkan N elemen, kemudian program menampilkan berbagai informasi seperti isi array, elemen berdasarkan indeks (ganjil, genap, dan kelipatan x), serta melakukan penghapusan elemen pada indeks tertentu.

Selain itu, program juga menghitung nilai rata-rata, standar deviasi, dan frekuensi kemunculan suatu bilangan dalam array.

Program memanfaatkan konsep array, perulangan, percabangan, dan operasi matematika dasar untuk pengolahan data.


