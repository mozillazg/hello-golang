package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)    // 同步阻塞
	c3 := make(chan string, 1) // 异步非阻塞

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			c3 <- "from 3"
			time.Sleep(time.Second * 4)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case msg3 := <-c3:
				fmt.Println(msg3)
			case <-time.After(time.Second):
				fmt.Println("timeout")
				// default:
				// 	fmt.Println("nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
