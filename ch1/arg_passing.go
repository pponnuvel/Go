package main

import (
	"fmt"
)

func main() {
	i := 5
	fmt.Println(i)
	test(&i)
	fmt.Println(i)
}

func test(i *int) {
	*i = *i + 1
}
