package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	digx := sha256.Sum256([]byte("x"))
	fmt.Printf("%x\n", digx)
}
