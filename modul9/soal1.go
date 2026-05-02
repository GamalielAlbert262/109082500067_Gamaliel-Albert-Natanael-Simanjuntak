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