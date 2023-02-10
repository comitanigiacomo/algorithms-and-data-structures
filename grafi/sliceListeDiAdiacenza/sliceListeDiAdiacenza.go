// utilizzpo questa rappresentazione nel caso in cui i vertici del grafo possono essere rappresentati implicitamente, poichè
// immediatamente riducibili a numeri naturali progressivi
// rappresento la relazione di adiacenza mediante un vettore di liste di adiacenza
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	g := leggiGafo()
	stampaGrafo(g)
	fmt.Println(arco(g, 4, 7))
	fmt.Println(arco(g, 2, 6))
}

type listNode struct {
	item int
	next *listNode
}

type linkedList struct {
	head *listNode
}

type grafo struct {
	n         int // numero di vertici
	adiacenti []*linkedList
}

func newNode(n int) *listNode {
	return &listNode{n, nil}
}

// restituisce l’indirizzo di un nuovo grafo con n nodi
func nuovoGrafo(n int) *grafo {
	var lista []*linkedList
	for i := 0; i < n+1; i++ {
		lista = append(lista, &linkedList{newNode(0)})
	}
	return &grafo{n, lista}
}

//aggiunge un nuovo nodo alla lista
func addNewNode(l *linkedList, n int) {
	node := newNode(n)
	if l.head.item != 0 {
		node.next = l.head
	}
	l.head = node
}

// stampa una lista
func stampaLista(l *linkedList) {
	p := l.head
	for p != nil {
		fmt.Print(p.item, " ")
		p = p.next
	}
	fmt.Println()
}

// legge un grafo da standard input
func leggiGafo() grafo {
	var g grafo
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		if len(args) == 1 {
			num, _ := strconv.Atoi(args[0])
			g = *nuovoGrafo(num)
		} else {
			num1, _ := strconv.Atoi(args[0])
			num2, _ := strconv.Atoi(args[1])
			addNewNode(g.adiacenti[num1], num2)
			addNewNode(g.adiacenti[num2], num1)
		}
	}
	return g
}

// stampa un grafo
func stampaGrafo(g grafo) {
	for i := 1; i < g.n+1; i++ {
		fmt.Print("vertice : ", i, " archi : ")
		stampaLista(g.adiacenti[i])
	}
}

//  dati due interi x e y stabilisca se c’è un arco tra x e y
func arco(g grafo, x, y int) bool {
	p := g.adiacenti[x].head
	for p != nil {
		if p.item == y {
			return true
		}
		p = p.next
	}
	return false
}
