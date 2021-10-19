package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {

	var n = flag.Bool("n", false, "Print newline")
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), " "))
	if !*n {
		fmt.Println()
	}
}
