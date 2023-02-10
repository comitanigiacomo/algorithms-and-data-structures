package main

import (
	"fmt"
	"strings"
)

func main() {
	var i int
	var t albero
	t.root = newNode(78)
	t.root.left = newNode(54)
	t.root.right = newNode(21)
	t.root.left.right = newNode(90)
	t.root.left.right.left = newNode(19)
	t.root.left.right.right = newNode(95)
	t.root.right.left = newNode(16)
	t.root.right.left.left = newNode(5)
	t.root.right.right = newNode(19)
	t.root.right.right.left = newNode(56)
	t.root.right.right.right = newNode(43)
	stampaAlbero(t.root, i)
}

type treeNode struct {
	item  int
	left  *treeNode
	right *treeNode
}

type albero struct {
	root *treeNode
}

func newNode(val int) *treeNode {
	return &treeNode{val, nil, nil}
}

func stampaAlbero(root *treeNode, i int) {
	if root == nil {
		return
	}
	fmt.Println(strings.Repeat(" ", i), "*", root.item)
	stampaAlbero(root.left, i+1)
	stampaAlbero(root.right, i+1)
}

func visitaPreOrdine(root *treeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.item, " ")
	visitaPreOrdine(root.left)
	visitaPreOrdine(root.right)
}

func visitaMidOrder(root *treeNode) {
	if root == nil {
		return
	}
	visitaMidOrder(root.left)
	fmt.Print(root.item, " ")
	visitaMidOrder(root.right)
}

func postOrder(root *treeNode) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Print(root.item)
}
