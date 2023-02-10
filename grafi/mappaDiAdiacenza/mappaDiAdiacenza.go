// utilizzo questa rappresentazione nel caso in cui i vertici del grafo
// non possono essere immediatamente ridotti a numeri naturali progressivi e quindi rappresentati implicitamente
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	g := leggiGrafo()
	stampaGrafo(g)
	hobbySeguiti("Riccardo", g)
	hobbyCheSeguono("Riccardo", g)
}

type twit struct {
	nome  string
	eta   int
	hobby []string
}

type vertice struct {
	car    twit
	chiave string
}

type grafo struct {
	vertici   map[string]*vertice
	adiacenti map[*vertice][]*vertice
}

//restituisce l’indirizzo di un nuovo grafo con n nodi.
func graphNew(n int) *grafo {
	return &grafo{make(map[string]*vertice), make(map[*vertice][]*vertice)}
}

//legge un grafo da standard input
func leggiGrafo() grafo {
	var g grafo
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		if len(args) == 1 {
			num, _ := strconv.Atoi(args[0])
			g = *graphNew(num)
		} else {
			num, err := strconv.Atoi(args[1])
			if err == nil {
				hob := strings.Split(args[2], ",")
				g.vertici[args[0]] = &vertice{twit{args[0], num, hob}, args[0]}
			} else {
				for i := 1; i < len(args); i++ {
					g.adiacenti[g.vertici[args[0]]] = append(g.adiacenti[g.vertici[args[0]]], g.vertici[args[i]])
				}
			}
		}
	}
	return g
}

// stampa un grafo
func stampaGrafo(g grafo) {
	for _, v := range g.vertici {
		fmt.Print(v.chiave, " segue : ")
		stampaSliceGrafo(g.adiacenti[v])
	}
}

// stampa il contenuto di una slice del grafo
func stampaSliceGrafo(l []*vertice) {
	if len(l) == 0 {
		fmt.Print("non segue nessuno")
	} else {
		for _, v := range l {
			fmt.Print(v.chiave, " ")
		}
	}
	fmt.Println()
}

func stampaSlice(l []string) {
	for _, v := range l {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

//Scrivete una funzione che data una stringa A stampi gli hobby dell’utente di nome A e
//l’elenco di tutti gli hobby delle persone seguite da A.
func hobbySeguiti(n string, g grafo) {
	fmt.Println("gli hobby di ", n, " sono : ")
	stampaSlice(g.vertici[n].car.hobby)
	fmt.Print("gli hobby delle persone seguite da ", n, " sono : ")
	for _, v := range g.adiacenti[g.vertici[n]] {
		stampaSlice(v.car.hobby)
	}
}

//Scrivete una funzione che data una stringa A stampi gli hobby dell’utente di nome A e
//l’elenco di tutti gli hobby delle persone che seguono A.
func hobbyCheSeguono(n string, g grafo) {
	fmt.Print("gli hobby di ", n, " sono : ")
	stampaSlice(g.vertici[n].car.hobby)
	fmt.Println("gli hobby delle persone che seguono ", n, " sono: ")
	for _, v := range g.vertici {
		for _, seguiti := range g.adiacenti[v] {
			if seguiti.car.nome == n {
				stampaSlice(v.car.hobby)
			}
		}
	}
}
