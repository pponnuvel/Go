package main

import (
	"fmt"
)

var lec int

func init() {
	fmt.Println("main initialized")
	lec = 99
}

func main() {
	fmt.Printf("lec = %d\n", lec)
}
