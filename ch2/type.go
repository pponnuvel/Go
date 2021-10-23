package main

import (
	"fmt"
)

func main() {
	type Index int
	type arr []int

	var i Index
	var a arr

	i = 0
	a = []int{0, 1, 2}
	for i = 0; i < Index(len(a)); i++ {
		fmt.Println(a[i])
	}
}
