package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	// Define flags
	maxp := flag.Int("mac", 6, "the max value")
	// Parse
	flag.Parse()
	bs := flag.Args()
	fmt.Println(*maxp)
	fmt.Println(bs)

	// Generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))
}
