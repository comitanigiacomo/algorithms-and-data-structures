package main

import "fmt"

func main() {
	var lista linkedList
	var l *linkedList = &lista
	addNewNode(l, 5)
	addNewNode(l, 10)
	addNewNode(l, 15)
	addNewNode(l, 20)
	printList(lista)
	fmt.Println(searchList(lista, 10))
	fmt.Println(searchList(lista, 27))

}

type listNode struct {
	item int
	next *listNode
}

type linkedList struct {
	head *listNode
}

// crea un nuovo nodo di lista
func newNode(val int) *listNode {
	return &listNode{val, nil}
}

//inserisce un nuovo nodo in testa alla lista
func addNewNode(l *linkedList, val int) {
	node := newNode(val)
	node.next = l.head
	l.head = node
}

// stampa una lista
func printList(l linkedList) {
	p := l.head
	for p != nil {
		fmt.Print(p.item, " ")
		p = p.next
	}
	fmt.Println()
}

// cerca un elemento nella lista
func searchList(l linkedList, val int) (bool, *listNode) {
	p := l.head
	for p != nil {
		if p.item == val {
			return true, p
		}
		p = p.next
	}
	return false, nil

}
