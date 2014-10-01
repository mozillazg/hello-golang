package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x, y, r float64
}

func (r *Rectangle) area() float64 {
	return math.Pi * r.r * r.r
}

type Shape interface {
	area() float64
}

// 等价于 python 代码： shapes=[c, r]; sum([s.area() for s in shapes])
// go 中无法将 c 和 r 组合成一个列表，因为它们的数据类型不一样
// 所以通过接口的方式调用
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

// 将接口作为结构体的字段
type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := Circle{1, 2, 3}
	r := Rectangle{2, 3, 4}
	// c 和 r 都有 area 方法，符合 Shape 的定义
	fmt.Println(totalArea(&c, &r))

	m := MultiShape{
		[]Shape{&c, &r},
	}
	fmt.Println(m.area())
}
