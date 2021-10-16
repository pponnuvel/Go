package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	io := bufio.NewScanner(os.Stdin)
	for io.Scan() {
		counts[io.Text()]++
	}

	for line, count := range counts {
		fmt.Println(line, " ", count)
	}
}
