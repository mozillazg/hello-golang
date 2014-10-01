package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

// Circle method, c.area
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circle{1, 2, 3}
	fmt.Println(c)
	fmt.Println(c.area())
}
