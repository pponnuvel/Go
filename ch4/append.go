package main

import (
	"fmt"
)

func main() {
	var sl []int
	printSlice(sl)

	sl = append(sl, 1)
	printSlice(sl)

	sl = append(sl, 2, 3, 4)
	printSlice(sl)
}

func printSlice(sl []int) {
	fmt.Printf(":")
	for _, v := range sl {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
