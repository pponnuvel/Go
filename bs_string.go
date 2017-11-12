/* Bubble sort for strings */
package main

import (
    "fmt"
    "os"
)

func main() {
    var tmp string
    length := len(os.Args)

    for i :=1; i < length; i++ {
        for j := 1; j < i; j++ {
            if os.Args[i] < os.Args[j] {
                tmp = os.Args[i]
                os. Args[i] = os.Args[j]
                os.Args[j] = tmp
            }
        }
    }

    for i := 1; i < length; i++ {
        fmt.Println(os.Args[i])
    }
}
