package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[rune]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		for _, v := range s {
			count[v]++
		}
	}

	for k, v := range count {
		fmt.Printf("%c = %d\n", k, v)
	}
}
