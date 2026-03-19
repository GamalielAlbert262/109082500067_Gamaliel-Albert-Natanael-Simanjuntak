package main
import (
	"fmt"
)

func factorial(n int64) int64 {

}
	if n == 0 || n == 1 {
		return 1
	}
     var result int64 = 1
	for i := int64(2); i <= n; i++ {
		result *= i
	}
	return result
   func permutation(n, r int64) int64 {
	if n < r {
		return 0
	}
    return factorial(n) / factorial(n-r)
}
func combination(n, r int64) int64 {
    if n < r {
		return 0
	}
	return factorial(n) / (factorial(r) * factorial(n-r))
}
func main() {
	var a,b,c,d int64
   fmt.Print("Masukkan 4 bilangan (a b c d): ")
    _, err := fmt.Scan(&a, &b, &c, &d)
    
	if err != nil {
	   fmt.Println("Error membaca input. Pastikan Anda memasukkan 4 bilangan bulat.")
       return
	}
	if a < c || b < d {
	   fmt.Println("Nilai a harus lebih besar atau sama dengan c, dan nilai b harus lebih besar atau sama dengan d.")
	   return
	}
    p1 := permutation(a, c)
	c1 := combination(a, c)
	p2 := permutation(b, d)
	c2 := combination(b, d)
    fmt.Println("\nkeluaran:")
    fmt.Printf("%d %d\n", p1, c1)
    fmt.Printf("%d %d\n", p2, c2)
}
