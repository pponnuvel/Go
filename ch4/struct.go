package main

import (
	"fmt"
)

func main() {
	type person struct {
		name, place string
		age         int
	}

	var x person = person{name: "abc", place: "cityx", age: 23}

	fmt.Println(x.age)
}
