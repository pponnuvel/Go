package main

import (
	"fmt"
)

func main() {
	s := "abcd efghij"
	fmt.Println(s[1:3])
	fmt.Println(s[0:])
	fmt.Println(s[:4])
	fmt.Println(s[:])

	t := "root"
	fmt.Println("left " + t + " right")

	raw := `hi\tnew\nmore"`
	fmt.Println(raw)
}
