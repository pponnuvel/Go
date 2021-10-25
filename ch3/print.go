package main

import (
	"fmt"
)

func main() {
	f := 5
	fmt.Printf("%d %#[1]x %#[1]X\n", f)
	fmt.Printf("%d %o %X\n", f, 0, 0)
	fmt.Printf("%g\n", 1.2020202020)
}
