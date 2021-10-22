package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("fib(%d) = %d\n", n, fib(n))
}

func fib(n int) int {
	x := 1
	y := 1

	for i := 1; i < n; i++ {
		x, y = y, x+y
	}

	return x
}
