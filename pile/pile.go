package main

import "fmt"

func main() {

	var p pila
	var punt *pila = &p
	fmt.Println(isEmpty(p))
	push(punt, 5)
	push(punt, 15)
	push(punt, 20)
	fmt.Println(isEmpty(p))
	pop(punt)
	pop(punt)
	pop(punt)

}

//posso implementare una pila come una lista concatenata, sulla quale posso eseguire tre semplici operazioni:

// 1) empty: verificare se la pila è vuota. quindi verificare se la coda è vuota

// 2) push: inserire un elemento nella pila: quindi inserire un nodo in testa alla coda

// 3) pop: estrarre un elemento dalla cima della pila. quindi cancellare l'elemento in testa alla coda

type node struct {
	item int
	next *node
}

type pila struct {
	head *node
}

// aggiunge un elemento alla pila
func newNode(val int) *node {
	return &node{val, nil}
}

// ritorna true se la pila è vuota, false altrimenti
func isEmpty(p pila) bool {
	first := p.head
	return first == nil
}

// inserisce un nuovo elemento nella pila
func push(p *pila, val int) {
	node := newNode(val)
	node.next = p.head
	p.head = node
}

// estrae un elemento dalla pila
func pop(p *pila) *node {
	node := p.head
	p.head = p.head.next
	fmt.Println(node.item)
	return node
}

// eventualmente posso anche implementarla mediante una slice nella quale ho due puntatori testa e coda che puntano rispettivamente al primo e all'ultimo elemento
// se uno dei due puntatori è nullo, la pila è vuota
// inserisco in testa e tolgo in testa
