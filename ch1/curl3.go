// go run curl.go https://pponnuvel.com

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := os.Args[1]

	if !strings.HasPrefix(url, "http://") {
		url = "https://" + url
	}

	resp, err := http.Get(url)
	fmt.Printf("Https status: %s\n", resp.Status)
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
