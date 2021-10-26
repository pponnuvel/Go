package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	digx384 := sha512.Sum384([]byte("x"))
	digx512 := sha512.Sum512([]byte("x"))
	fmt.Printf("%x\n", digx384)
	fmt.Printf("%x\n", digx512)
}
