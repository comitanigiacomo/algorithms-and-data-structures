package main

import (
	"fmt"
	"math"
)

type grafo []arco

type arco struct {
	x, y string
	val  float64
}

type vertice struct {
	str string
	val float64
}

func creaVettore(g grafo) (map[string]float64, []vertice) {
	c := []vertice{}
	mappa := make(map[string]bool)
	d := make(map[string]float64)

	for _, v := range g {
		if !mappa[v.x] {
			c = append(c, vertice{v.x, math.Inf(1)})
			mappa[v.x] = true
			d[v.x] = math.Inf(1)
		}

		if !mappa[v.y] {
			c = append(c, vertice{v.y, math.Inf(1)})
			d[v.y] = math.Inf(1)
			mappa[v.y] = true
		}
	}
	return d, c
}

// ordina la coda in ordine crescente in base ai valori dei vertici
// Essendo la coda una coda di priorità, è opportuno riordinarla ogni volta che viene modificata
func ordinaCoda(c []vertice) []vertice {
	for i := 1; i < len(c); i++ {
		if c[i].val < c[i-1].val {
			c[i-1], c[i] = c[i], c[i-1]
			i = 0
		}
	}
	return c
}

// ritorna true se il vertice è già contenuto nell'albero, false altrimenti
func contiene(tree []arco, s string) bool {
	for _, v := range tree {
		if v.x == s || v.y == s {
			return true
		}
	}
	return false
}

// aggiorna la coda di priorità
func changeKey(c []vertice, val float64, z string) []vertice {
	for i := range c {
		if c[i].str == z {
			c[i].val = val
			break
		}
	}
	return c
}

func Prim(g grafo) []arco {

	d, c := creaVettore(g)
	vicino := make(map[string]string)
	var tree grafo

	for {
		c = ordinaCoda(c)
		v := c[0]
		c = c[1:]
		if d[v.str] != math.Inf(1) {
			x := vicino[v.str]
			tree = append(tree, arco{x, v.str, d[v.str]})
		}

		for _, a := range g {
			if a.x == v.str {
				if !contiene(tree, a.y) && a.val < d[a.y] {
					d[a.y] = a.val
					vicino[a.y] = v.str
					c = changeKey(c, a.val, a.y)
				}
			}
		}

		if len(c) == 0 {
			break
		}
	}
	return tree
}

func aggiungiArco(g []arco, a, b string, val float64) []arco {
	g = append(g, arco{a, b, val})
	g = append(g, arco{b, a, val})
	return g
}

func main() {
	var g grafo
	g = aggiungiArco(g, "A", "B", 17)
	g = aggiungiArco(g, "B", "E", 3)
	g = aggiungiArco(g, "E", "G", 15)
	g = aggiungiArco(g, "G", "F", 20)
	g = aggiungiArco(g, "F", "C", 23)
	g = aggiungiArco(g, "C", "A", 28)
	g = aggiungiArco(g, "A", "D", 25)
	g = aggiungiArco(g, "B", "D", 16)
	g = aggiungiArco(g, "E", "D", 9)
	g = aggiungiArco(g, "G", "D", 4)
	g = aggiungiArco(g, "F", "D", 1)
	g = aggiungiArco(g, "C", "D", 36)

	fmt.Println(Prim(g))

}
