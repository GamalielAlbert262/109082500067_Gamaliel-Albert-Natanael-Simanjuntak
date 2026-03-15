# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">[Gamaliel Albert Natanael Simanjuntak] - [109082500067]</p>

## Unguided 

### 1. [Soal]
#### soal1.go
Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan
masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang
diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?
```go
package main
import "fmt"
func main() {
var (
satu, dua, tiga string
temp string
)
fmt.Print("Masukan input string: ")
fmt.Scanln(&satu)
fmt.Print("Masukan input string: ")
fmt.Scanln(&dua)
fmt.Print("Masukan input string: ")
fmt.Scanln(&tiga)
fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
temp = satu
satu = dua
dua = tiga
tiga = temp
fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/GamalielAlbert262/109082500067_Gamaliel-Albert-Natanael-Simanjuntak/blob/main/modul2/outputscreenshot/soal1.png)
#### Penjelasan
Menerima Input: Program akan meminta pengguna untuk mengetikkan tiga data teks satu per satu. Data ini disimpan berturut-turut dalam tiga variabel: satu, dua, dan tiga.

Menampilkan Urutan Awal: Program mencetak urutan teks asli yang baru saja dimasukkan oleh pengguna.

Contoh: Output awal = [teks 1] [teks 2] [teks 3]

Proses Pertukaran (Pergeseran): Inilah bagian inti dari program tersebut. Program menggunakan variabel bantuan bernama temp untuk memindahkan isi variabel-variabel tersebut tanpa kehilangan datanya.

Nilai yang awalnya di variabel pertama (satu) dipindahkan ke variabel terakhir (tiga).

Nilai dari variabel kedua (dua) dipindahkan menjadi isi variabel pertama (satu).

Nilai dari variabel ketiga (tiga) dipindahkan menjadi isi variabel kedua (dua).

Menampilkan Urutan Akhir: Terakhir, program menampilkan urutan teks setelah posisinya diubah.

Ringkasnya, urutan yang awalnya adalah satu -> dua -> tiga, setelah program dijalankan akan berubah menjadi dua -> tiga -> satu. Setiap data bergeser satu posisi ke kiri, dan data yang paling kiri pindah ke paling kanan secara melingkar.

