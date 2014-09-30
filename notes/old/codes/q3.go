/* Q3 */

package main

import "fmt"


func q1() {
    for i := 1; i < 101; i++ {
        if i % 3 == 0 && i % 5 == 0 {
            fmt.Printf("FizzBuzz ")
        } else if i % 3 == 0 {
            fmt.Printf("Fizz ")
        } else if i % 5 == 0 {
            fmt.Printf("Buzz ")
        } else {
            fmt.Printf("%d ", i)
        }
    }
}


func main() {
    q1()
}
