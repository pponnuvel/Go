/* Bubble sort for integers */
package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    var tmp string
    length := len(os.Args)

    for i :=1; i < length; i++ {
        for j := 1; j < i; j++ {
            val, err1 := strconv.Atoi(os.Args[i])
            val2, err2 := strconv.Atoi(os.Args[j])
            if err1 != nil || err2 != nil {
                fmt.Println("Invalid input found")
                os.Exit(1)
            }
            if val < val2 {
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
