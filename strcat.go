/* Concatenate strings */
package main

import (
	"fmt"
	"os"
)

func main() {
	var res string
	for i := 1; i < len(os.Args); i++ {
		if len(res) != 0 {
			res += "-" + os.Args[i]
		} else {
			res = os.Args[i]
		}
	}

	fmt.Println(res)
}
