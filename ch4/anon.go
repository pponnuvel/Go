package main

import "fmt"

type shape struct {
	height, width int
}

type rect struct {
	shape
}

func main() {
	var r rect
	r.height = 4
	r.width = 10

	fmt.Printf("Height = %d, Width = %d\n", r.height, r.width)
}
