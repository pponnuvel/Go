package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	i := 4
	for i != 0 {
		fmt.Println(i)
		i--
	}

	for i, str := range os.Args[1:] {
		fmt.Println(i, " = ", str)
	}
}
