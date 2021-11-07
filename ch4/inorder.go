package main

import (
	"fmt"
)

type tree struct {
	data        int
	left, right *tree
}

func main() {
	var r tree = tree{data: 8}
	r.left = &tree{data: 4}
	r.right = &tree{data: 16}

	inorder(&r)
}

func inorder(root *tree) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Println(root.data)
	inorder(root.right)
}
