// go run curl.go https://pponnuvel.com

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		error_exit(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		error_exit(err)
	}

	fmt.Printf("%s\n", body)
}

func error_exit(err error) {
	fmt.Printf("%v\n", err)
	os.Exit(1)
}
