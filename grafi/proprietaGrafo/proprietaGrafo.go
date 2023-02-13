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
	fmt.Println(degree(g, "2"))
	mappa := make(map[string]bool)
	mappa2 := make(map[string]bool)
	percorso := []string{"1"}
	p := (Path(g, "1", "6", mappa, mappa2, percorso))
	fmt.Println(len(p) == 1)
}

type vertice struct {
	chiave string
	peso   int
}

type grafo struct {
	adiacenti map[string][]*vertice
}

func leggiGrafo() grafo {
	var g grafo
	g.adiacenti = make(map[string][]*vertice)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		num, _ := strconv.Atoi(args[2])
		g.adiacenti[args[0]] = append(g.adiacenti[args[0]], &vertice{args[1], num})
		g.adiacenti[args[1]] = append(g.adiacenti[args[1]], &vertice{args[0], num})
	}
	return g
}

// stampa un grafo
func stampaGrafo(g grafo) {
	for v, ad := range g.adiacenti {
		fmt.Print("adiacenti a ", v, ": ")
		for _, k := range ad {
			fmt.Print(k.chiave, " ")
		}
		fmt.Println()
	}
}

// ritorna il grado di un vertice del grafo
func degree(g grafo, v string) int {
	return len(g.adiacenti[v])
}

// eseguo una DFS per trovare se esiste un cammino all'interno del grafo tra il vertice di chiave v e quello di chiave w
func Path(g grafo, v, w string, aux map[string]bool, mappa map[string]bool, percorso []string) map[string]bool {
	if v == w {
		var s string
		for _, v := range percorso {
			s += v
		}
		mappa[v] = true

	}
	aux[v] = true
	for _, v := range g.adiacenti[v] {
		if !aux[v.chiave] {
			percorso = append(percorso, v.chiave)
			Path(g, v.chiave, w, aux, mappa, percorso)
			percorso = percorso[:len(percorso)-1]
		}
	}
	aux[v] = false
	return mappa
}
