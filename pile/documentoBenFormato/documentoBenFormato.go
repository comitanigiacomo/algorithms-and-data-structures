package main

import (
	"fmt"
	"strings"
)

func main() {

	benFormato("<a> <b> </b> </c>")

}

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

// inserisce un nuovo elemento nella coda
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

// stampa il contenuto della pila
func stampaLista(p pila) {
	node := p.head
	for node != nil {
		fmt.Print(node.item, " ")
		node = node.next
	}
	fmt.Println()
}

//ritorna true se la pila è vuota, false altrimenti
func isEmpty(p pila) bool {
	return p.head != nil
}

func benFormato(espr string) bool {
	var p pila
	args := strings.Split(espr, " ")
	for i, v := range args {
		if v[1] == '/' {
			tmp := pop(&p)
			if tmp.item[1] != v[2] {
				fmt.Println("errore in pos", i+1)
				return false
			}
		} else {
			push(&p, v)
		}
	}
	if isEmpty(p) {
		fmt.Print("sono rimasti aperti i seguenti tag: ")
		stampaLista(p)
		return false
	}
	fmt.Println("il documento è ben formato")
	return true
}
