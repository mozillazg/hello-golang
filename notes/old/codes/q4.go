/* Q4 */

package main

import (
    "fmt"
    "unicode/utf8"
)


func q1() {
    // 建立一个 Go 程序打印下面的内容（到 100 个字符）：
    // A
    // AA
    // AAA
    // AAAA
    // AAAAA
    // AAAAAA
    // AAAAAAA
    // ...

    for i := 0; i < 100; i++ {
        a := "A"
        for x :=0; x < i; x++ {
            a += "A"
        }
        fmt.Printf("%s\n", a)
    }
}


func q2() {
    // 建立一个程序统计字符串里的字符数量：
    // asSASA ddd dsjkdsjs dk
    // 同时输出这个字符串的字节数。
    // 提示：
    // 看看 unicode/utf8 包。

    str := "asSASA dd dsjkdsjs dk 你好"
    fmt.Println(str)
    fmt.Printf("%d ", len(str))
    n := utf8.RuneCount([]byte(str))
    //n := utf8.RuneCountInString(str)
    fmt.Printf("%d \n", n)
}


func q3() {
    // 扩展上一个问题的程序，替换位置 4 开始的三个字符为 “abc”。

    str := "asSASA dd dsjkdsjs dk 你好"
    s := str[:4] + "abc" + str[7:]
    fmt.Println(str)
    fmt.Println(s)
}


func q4() {
    // 编写一个 Go 程序可以逆转字符串，例如 “foobar” 被打印成 “raboof”。
    // 提示：不幸的是你需要知道一些关于转换的内容，
    // 参阅 “转换” 第 58 页的内容。

    str := "foobar 你好"
    b := []rune(str)
    length := len(b)
    s := ""
    for i := length - 1; i >= 0; i-- {
        s += string(b[i])
    }
    fmt.Println(str)
    fmt.Println(s)
}


func main() {
    q1()
    q2()
    q3()
    q4()
}
