package main

import (
	"fmt"
	"strings"
)

func main() {
	var i, count int
	a := []int{69, 89, 28, 39, 66, 44, 12, 2, 71}
	root := arr2tree(a, i)
	stampaAlbero(root, count)
}

type treeNode struct {
	item  int
	left  *treeNode
	right *treeNode
}

func newNode(val int) *treeNode {
	return &treeNode{val, nil, nil}
}

// costruisce un albero binario a partire da una slice di interi
func arr2tree(a []int, i int) (root *treeNode) {
	if i >= len(a) {
		return nil
	}
	root = newNode(a[i])
	root.left = arr2tree(a, 2*i+1)
	root.right = arr2tree(a, 2*i+2)
	return root
}

func stampaAlbero(root *treeNode, i int) {
	if root == nil {
		return
	}
	fmt.Println(strings.Repeat(" ", i), "*", root.item)
	stampaAlbero(root.left, i+1)
	stampaAlbero(root.right, i+1)
}
