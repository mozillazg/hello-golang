package main

import (
	"./math"
	// m "./math"   // 别名
	"fmt"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
