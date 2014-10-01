package main

/* Android is a Person */

import "fmt"

type Person struct {
	name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is ", p.name)
}

// Android 继承 Person
// 拥有 Person 的所有字段和方法
type Android struct {
	Person
	model string
}

func main() {
	a := new(Android)
	a.Talk()

	b := Android{
		Person{
			"abc",
		},
		"android 4",
	}
	b.Talk()
	fmt.Println(b.name)
}
