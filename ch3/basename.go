package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(basename(s))
	}
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	return s[slash+1:]
}
