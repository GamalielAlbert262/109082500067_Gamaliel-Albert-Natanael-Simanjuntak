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


### 3. [Soal]
#### soal3.go

```go
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
```

### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/GamalielAlbert262/109082500067_Gamaliel-Albert-Natanael-Simanjuntak/blob/main/modul9/outputscreenshot/soal3.png)

#### [penjelasan]
 Program ini dibuat untuk merekap hasil pertandingan sepak bola antara dua klub yang bertanding dalam beberapa pertandingan. Program bekerja dengan cara meminta pengguna memasukkan nama kedua klub, kemudian dilanjutkan dengan memasukkan skor hasil pertandingan secara berulang.

Setiap input skor terdiri dari dua bilangan bulat yang masing-masing merepresentasikan jumlah gol dari Klub A dan Klub B. Program akan terus menerima input skor hingga ditemukan nilai skor yang tidak valid, yaitu ketika salah satu atau kedua skor bernilai negatif. Kondisi tersebut digunakan sebagai penanda untuk menghentikan proses input.

Untuk setiap pertandingan yang dimasukkan, program akan melakukan proses perbandingan skor untuk menentukan hasil pertandingan. Jika skor Klub A lebih besar dari Klub B, maka Klub A dinyatakan sebagai pemenang. Sebaliknya, jika skor Klub B lebih besar, maka Klub B dinyatakan sebagai pemenang. Jika kedua skor sama, maka pertandingan dinyatakan seri (draw).

Program kemudian menampilkan hasil dari setiap pertandingan secara langsung. Selain itu, program juga menyimpan nama klub yang memenangkan pertandingan ke dalam sebuah array. Hanya pertandingan yang menghasilkan pemenang yang akan disimpan, sedangkan hasil seri tidak dimasukkan ke dalam array.

Setelah proses input selesai, program akan menampilkan pesan bahwa pertandingan telah selesai, kemudian diikuti dengan daftar nama klub yang memenangkan pertandingan berdasarkan data yang telah disimpan di dalam array.

Secara keseluruhan, program ini mengimplementasikan konsep dasar pemrograman berupa penggunaan array untuk penyimpanan data, struktur perulangan untuk memproses input berulang, serta percabangan untuk menentukan hasil berdasarkan kondisi tertentu. Program ini juga menerapkan validasi input sederhana untuk mengontrol jalannya proses eksekusi.


### 4. [Soal]
#### soal4.go

```go
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0

	for {
		fmt.Scanf("%c", &ch)

		if ch == '.' || *n >= NMAX {
			break
		}

		if ch != ' ' && ch != '\n' {
			t[*n] = ch
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]), " ")
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks : ")
	isiArray(&tab, &m)

	fmt.Print("Isi teks : ")
	cetakArray(tab, m)

	if palindrom(tab, m) {
		fmt.Println("Palindrom ? true")
	} else {
		fmt.Println("Palindrom ? false")
	}

	balikanArray(&tab, m)

	fmt.Print("Reverse teks : ")
	cetakArray(tab, m)
}
```

### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/GamalielAlbert262/109082500067_Gamaliel-Albert-Natanael-Simanjuntak/blob/main/modul9/outputscreenshot/soal4.png)

#### [penjelasan]
Program ini dibuat untuk mengolah sekumpulan karakter yang disimpan dalam sebuah array. Program memiliki beberapa fungsi utama, yaitu mengisi array dengan input karakter dari pengguna, menampilkan isi array, membalik urutan elemen array, serta memeriksa apakah susunan karakter tersebut merupakan palindrom.

Pada awalnya, program meminta pengguna untuk memasukkan sejumlah karakter satu per satu. Proses input akan terus berlangsung hingga pengguna memasukkan tanda titik (.) sebagai penanda akhir input atau hingga kapasitas maksimum array tercapai. Karakter yang dimasukkan kemudian disimpan ke dalam array bertipe rune.

Setelah proses input selesai, program akan menampilkan isi array sesuai dengan urutan awal yang dimasukkan. Selanjutnya, program melakukan pemeriksaan apakah susunan karakter tersebut termasuk palindrom. Pemeriksaan dilakukan dengan membandingkan elemen array dari sisi depan dan belakang secara berpasangan. Jika seluruh pasangan karakter tersebut sama, maka data dinyatakan sebagai palindrom. Sebaliknya, jika terdapat perbedaan, maka data tersebut bukan palindrom.

Program juga memiliki prosedur untuk membalik urutan isi array. Proses pembalikan dilakukan dengan menukar elemen pertama dengan elemen terakhir, elemen kedua dengan elemen kedua dari belakang, dan seterusnya hingga mencapai tengah array. Setelah proses pembalikan selesai, program menampilkan kembali isi array dalam urutan yang telah dibalik.

Secara keseluruhan, program ini mengimplementasikan konsep dasar pemrograman seperti penggunaan array untuk penyimpanan data, prosedur dan fungsi untuk modularisasi program, perulangan untuk proses input dan manipulasi data, serta percabangan untuk pengambilan keputusan. Selain itu, program ini juga menunjukkan penerapan algoritma sederhana untuk pembalikan data dan pengecekan palindrom, yang merupakan dasar dalam pengolahan string atau karakter.



