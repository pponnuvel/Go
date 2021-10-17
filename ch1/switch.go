package main

import (
	"fmt"
)

func main() {
	fmt.Println(sw(10))
	fmt.Println(sw(-10))
	fmt.Println(sw(0.0))
}

func sw(x int) int {

	switch {
	case x < 0:
		return -1
	default:
		return 0
	case x > 0:
		return 1
	}
}
