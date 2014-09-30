package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, z float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.z * c.z
}

func circleNew(c *Circle) {
	c.x = 1
	c.y = 5
}

func main() {
	// var c Circle
	// c := new(Circle)
	// c := Circle{0, 1, 5}
	c := Circle{x: 0, y: 1, z: 2}
	fmt.Println(c.x, c.y, c.z)
	fmt.Println(c)

	fmt.Println(circleArea(c))

	circleNew(&c) // 修改
	fmt.Println(c)
}
