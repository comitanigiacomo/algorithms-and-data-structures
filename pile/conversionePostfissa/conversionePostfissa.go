package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(converti("( ( 5 - 3 ) * 2 )"))
}

// questa volta implemento la pila mediante lista concatenata

type node struct {
	item string
	next *node
}
type pila struct {
	head *node
}

func newNode(val string) *node {
	return &node{val, nil}
}

// ritorna true se la pila è vuota, false altrimenti
/*func isEmpty(p pila) bool {
	return p.head == nil
}*/

// inserisce un nuovo elemento nella pila
func push(p *pila, val string) {
	node := newNode(val)
	node.next = p.head
	p.head = node
}

// rimuove l'elemento in cima alla pila
func pop(p *pila) *node {
	tmp := p.head
	p.head = p.head.next
	return tmp
}

//  riceve una espressione in notazione infissa e restituisce l’espressione equivalente in notazione postfissa
func converti(espressione string) string {
	var s string
	var p pila
	for _, v := range espressione {
		if unicode.IsNumber(v) {
			s += string(v) + " "
		} else if v == '*' || v == '/' || v == '+' || v == '-' {
			push(&p, string(v))
		} else if v == ')' {
			node := pop(&p)
			s += node.item + " "
		}
	}
	return s
}
