package main

import (
	"fmt"
)

func main() {
	a := [3]int{1, 2, 3}
	var b [3]int
	b = [3]int{1, 2, 3}

	fmt.Println(a == b)

	c := [10]int{1, 7: -1}

	for _, v := range c {
		fmt.Println(v)
	}
}
