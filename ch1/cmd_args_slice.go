package main

import (
	"fmt"
	"os"
)

func main() {
	var l = len(os.Args)
	fmt.Println(os.Args[l-1 : l])
}
