// go run curl.go https://pponnuvel.com

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("Elased main: %fs\n", time.Since(start).Seconds())
}

func error_exit(err error) {
	fmt.Printf("%v\n", err)
	os.Exit(1)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		error_exit(err)
	}

	bytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	elapsed := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%fs %d %s\n", elapsed, bytes, url)
}
