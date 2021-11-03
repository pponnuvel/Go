package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	names := make(map[string]bool)

	for scanner.Scan() {
		s := scanner.Text()
		if !names[s] {
			names[s] = true
			fmt.Println(s)
		}
	}
}
