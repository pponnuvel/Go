package main

import (
	"fmt"
)

func main() {
	var sl []int

	for i := 1; i < 10; i++ {
		sl = appendInt(sl, i)
	}
	printSlice(sl)

	sl2 := []int{21, 22, 23}
	sl = appendInt(sl, sl2...)
	printSlice(sl)
}

func printSlice(sl []int) {
	fmt.Printf(":")
	for _, v := range sl {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func appendInt(sl []int, x ...int) []int {
	var newSl []int
	old_len := len(sl)
	new_len := old_len + len(x)

	if new_len <= cap(sl) {
		newSl = sl[:new_len]
	} else {
		new_cap := new_len
		if new_cap < 2*old_len {
			new_cap *= 2
		}
		newSl = make([]int, new_len, new_cap)
		copy(newSl, sl)
	}

	copy(newSl[old_len:], x)
	return newSl
}
