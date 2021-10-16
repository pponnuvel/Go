package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	for _, name := range files {
		f, err := os.Open(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			continue
		}
		countLines(f, counts)
		f.Close()
	}

	for line, count := range counts {
		fmt.Println(line, " ", count)
	}
}

func countLines(f *os.File, counts map[string]int) {
	io := bufio.NewScanner(f)
	for io.Scan() {
		counts[io.Text()]++
	}
}
