package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var arr []int

	for i := range os.Args[1:] {
		x, err := strconv.Atoi(os.Args[i+1])
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		arr = append(arr, x)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}

}
