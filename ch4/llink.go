package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type list struct {
	data int
	next *list
}

func getnode(data int) *list {
	return &list{data, nil}
}

func getlist() *list {
	var root *list = nil
	io := bufio.NewScanner(os.Stdin)
	for io.Scan() {
		d, _ := strconv.Atoi(io.Text())
		node := getnode(d)
		if root == nil {
			root = node
		} else {
			node.next = root
			root = node
		}
	}

	return root
}

func print_list(root *list) {
	if root == nil {
		fmt.Println("nil")
		return
	}

	fmt.Printf("%d->", root.data)
	print_list(root.next)
}

func main() {
	var root *list = getlist()
	print_list(root)
}
