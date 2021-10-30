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

	arr := [3]int{5, 6, 7}
	sl = append(sl, arr[:]...)
	printSlice(sl)

	sl2 := []int{8, 9}
	sl = append(sl, sl2...)
	printSlice(sl)
}

func printSlice(sl []int) {
	fmt.Printf(":")
	for _, v := range sl {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
