package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	var buf bytes.Buffer
	for _, a := range os.Args[1:] {
		fmt.Fprint(&buf, a)
		fmt.Fprint(&buf, " ")
	}

	fmt.Println(strings.TrimSpace(buf.String()))
}
