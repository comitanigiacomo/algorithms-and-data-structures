package main

func main() {

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

func readTree(s string) *treeNode {

}

func arr2tree(a []int, i int) (root *treeNode) {
	if i >= len(a) {
		return nil
	}
	root = newNode(a[i])
	root.left = arr2tree(a, 2*i+1)
	root.right = arr2tree(a, 2*i+2)
	return root
}
