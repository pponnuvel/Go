package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Lines to file\nSecond line\n")
	err := ioutil.WriteFile(os.Args[1], data, 0644)
	if err != nil {
		fmt.Print("%v\n", err)
		os.Exit(1)
	}

}
