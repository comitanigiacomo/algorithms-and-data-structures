// Trovare i cammini minimi tra un vertice e tutti gli altri

// Pesi non negativi

package main

import (
	"fmt"
	"math"
)

type arco struct {
	x, y string
	val  float64
}

type grafo []arco

type vertice struct {
	str string
	val float64
}

func creaVettore(g grafo, s string) map[string]float64 {
	visitati := make(map[string]bool)
	d := make(map[string]float64)
	for _, v := range g {
		if (v.x == s || v.y == s) && !visitati[s] {
			visitati[s] = true
			d[s] = 0
		}

		if !visitati[v.x] {
			d[v.x] = math.Inf(1)
			visitati[v.x] = true
		}
		if !visitati[v.y] {
			d[v.y] = math.Inf(1)
			visitati[v.y] = true
		}
	}
	return d
}

func creaCoda(d map[string]float64) []vertice {
	var c []vertice
	for str, val := range d {
		c = append(c, vertice{str, val})
	}
	return c
}

func riordinaCoda(c []vertice) []vertice {
	for i := 1; i < len(c); i++ {
		if c[i].val < c[i-1].val {
			c[i], c[i-1] = c[i-1], c[i]
			i = 0
		}
	}
	return c
}

func dijkstra(g grafo, s string) map[string]float64 {
	d := creaVettore(g, s)

	c := creaCoda(d)
	c = riordinaCoda(c)

	for len(c) != 0 {
		u := c[0]
		c = c[1:]
		c = riordinaCoda(c)

		for _, v := range g {
			if v.x == u.str { // devo considerare tutti gli archi che partono da u
				if d[u.str]+v.val < d[v.y] {
					d[v.y] = d[u.str] + v.val
				}
			}

		}
	}
	return d
}

func main() {
	g := grafo{
		{"E", "F", 3},
		{"A", "E", 1},
		{"A", "B", 3},
		{"B", "C", 5},
		{"B", "E", 2},
		{"E", "C", 6},
		{"F", "D", 5},
		{"C", "D", 2},
		{"F", "C", 2},
	}
	fmt.Println(dijkstra(g, "A"))
}
