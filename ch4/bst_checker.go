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

	fmt.Println(bst_checker(&r))
}

func bst_checker(root *tree) bool {
	if root == nil || root.left == nil || root.right == nil {
		return true
	}

	return (root.data > root.left.data) &&
		(root.data < root.right.data) &&
		bst_checker(root.left) &&
		bst_checker(root.right)
}
