/* Q5 */

package main

import "fmt"


func q1() {
    // 编写计算一个类型是 float64 的 slice 的平均值的代码。
    se := [...]float64 {1, 2, 3, 4, 5, 6}
    avg := 0.0
    length := len(se)
    if length > 0 {
        sum := 0.0
        for _, v := range se {
            sum += float64(v)
        }
        avg = sum / float64(length)
    }
    fmt.Println(se)
    fmt.Println(avg)
}


func main() {
    q1()
}
