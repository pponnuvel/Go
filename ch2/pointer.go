package main

import (
	"fmt"
)

func main() {
	i := 9
	fmt.Println(test(&i))
}

func test(p *int) int {
	*p++
	return *p
}
