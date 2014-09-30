/* Q2 */

package main

import "fmt"


// 1
func q1() {
    for i := 0; i < 10; i++ {
        fmt.Printf("%d ", i)
    }
}


// 2
func q2() {
    i := 0
here:
    if i < 10 {
        fmt.Printf("%d ", i)
        i++
        goto here
    }
}


// 3
func q3() {
    arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    for v := range arr {
        fmt.Printf("%d ", v)
    }
}


func main() {
    fmt.Println("q1")
    q1()
    println("\nq2")
    q2()
    println("\nq3")
    q3()
}
