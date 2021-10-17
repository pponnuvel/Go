// go run curl.go https://pponnuvel.com

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		error_exit(err)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	resp.Body.Close()
}

func error_exit(err error) {
	fmt.Printf("%v\n", err)
	os.Exit(1)
}
