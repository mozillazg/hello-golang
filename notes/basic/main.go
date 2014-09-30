package main

import "fmt"

// 定义变量
var a int
var b, c, d int
var e int = 'e'
var f, g, h int = 'f', 'g', 'h'
var i, j = 'i', 'j'
var (
	a1 = 1
	a2 = 2
)

func basic() {
	// 只能在函数内部使用
	k, l := 'k', 'l'
	fmt.Printf("%d, %d ", k, l)

	// 常量
	const Pi = 3.1415926

	// 数据类型
	// int
	var t1 int = 1
	var t2 int32 = 1
	var t3 int64 = 1
	var t4 uint = 1
	var t5 uint64 = 1
	fmt.Println(t1, t2, t3, t4, t5)

	// float
	var f1 float32 = 1.2
	var f2 float64 = 1.3
	fmt.Println(f1, f2)

	// bool
	var isActive bool
	var enabled = true
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	f3, f4 := 0, 0
	fmt.Println("f3 == f4: ", f3 == f4)

	// string
	var str string = ""
	str = "abc" /* 只能在函数内重新赋值 */
	str = "efg"
	abc := `abc
abc`
	fmt.Println(str + "123")
	fmt.Println("str[0]: ", str[0])

	// array
	var arr [10]int
	arr[0] = 10

	fmt.Println(isActive, enabled, str, abc, arr)
}

func scanf() {
	fmt.Println("Input a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}

func controlFor() {
	i := 1
	for i <= 10 {
		fmt.Print(i, " ")
		i += 1
	}
	fmt.Println("")

	for n := 1; n <= 10; n++ {
		fmt.Print(n, " ")
	}
}

func controlIf() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print(i, " even, ")
		} else {
			fmt.Print(i, " odd, ")
		}
	}
}

func controlSwitch() {
	for i := 0; i <= 6; i++ {
		switch i {
		case 0:
			fmt.Print("Zero ")
		case 1:
			fmt.Print("One ")
		case 2:
			fmt.Print("Two ")
		case 3:
			fmt.Print("Three ")
		case 4:
			fmt.Print("Four ")
		default:
			fmt.Print("Unknown Number ")
		}
	}
}

// Array
func fArray() {
	var x [5]int
	x[4] = 100
	fmt.Print(x)
	fmt.Print(x[4])
	fmt.Println("")

	y := [5]float64{
		1,
		2,
		3,
		4,
		5,
		// 6,
	}
	var total float64 = 0
	/* range 遍历 array
	_ 定义一个不想使用的变量
	*/
	for _, value := range y {
		total += value
	}
	fmt.Println(total / float64(len(y)))
}

// slice 长度可变的 array
func aSlice() {
	// x := make([]float64, 5)
	arr := []float64{1, 2, 3, 4, 5}
	x := arr[0:5]
	x = arr[:5]
	x = arr[3:]
	y := append(x, 6, 7) // 合并 x 和 6, 7
	fmt.Println(x, y)

	a := []int{1, 2, 3}
	b := make([]int, 2)
	copy(b, a) // 从 a 中复制数据到 b, 由于 b 长度为 2，所以只复制了两个数
	fmt.Println(a, b)
}

// map  key-value
func aMap() {
	x := make(map[string]int) // key 的类型是 string, value 的类型是 int
	x["key"] = 10
	x["abc"] = 1
	fmt.Println(x)
	fmt.Println(x["key"])
	delete(x, "abc")
	fmt.Println(x)
	// 当 key 不存在时, ok 为 false
	value, ok := x["12345"]
	fmt.Println(value, ok)
	value2, ok2 := x["key"]
	fmt.Println(value2, ok2)

	elements := map[string]string{
		"a": "abc",
		"f": "fire",
		"g": "google",
	}
	// 如果存在 key = a, 输出
	if el, ok := elements["a"]; ok {
		fmt.Println(el)
	}

	// value 的值可以是一个 map
	ele2 := map[string]map[string]string{
		"a": map[string]string{
			"name": "abcd",
		},
	}
	if value, ok := ele2["a"]; ok {
		fmt.Println(value["name"])
	}

}

func main() {
	// basic()
	// scanf()
	// controlFor()
	// controlIf()
	// controlSwitch()
	// fArray()
	// aSlice()
	aMap()
}
