//Trova i cammini minimi tra un vertice e tutti gli altri

// tecnica basata sulla propagazione delle informazioni e programmazione dinamica

package main

import (
	"fmt"
	"math"
)

type arco struct {
	x, y string
	val  float64
}

type vertice struct {
	str string
	val int
}
type grafo []arco

func creaVettore(g grafo, s string, mappa map[string]bool) map[string]float64 {
	d := make(map[string]float64)
	for _, v := range g {
		if !mappa[v.x] && v.x != s {
			mappa[v.x] = true
			d[v.x] = math.Inf(1)
		}
		if !mappa[v.y] && v.y != s {
			mappa[v.y] = true
			d[v.y] = math.Inf(1)
		}
		if (v.x == s || v.y == s) && !mappa[s] {
			mappa[s] = true
			d[s] = 0
		}
	}
	return d
}

func BellmanFord(g grafo, s string) map[string]float64 {
	mappa := make(map[string]bool)
	d := creaVettore(g, s, mappa)

	for range g {
		for _, v := range g {

			if d[v.x]+v.val < d[v.y] {
				d[v.y] = d[v.x] + v.val
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

	fmt.Println(BellmanFord(g, "A"))
}
