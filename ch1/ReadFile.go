package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Print("%v\n", err)
		os.Exit(1)
	}

	lines := 0
	for _, s := range strings.Split(string(data), "\n") {
		if s != "" {
			fmt.Println(s)
			counts[s]++
			lines++
		}
	}

	fmt.Printf("Lines = %d\n", lines)

	for s, c := range counts {
		fmt.Println(s, "=", c)
	}

	//	return 9

}
