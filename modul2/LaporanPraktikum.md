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
[penjelasan]
testttt