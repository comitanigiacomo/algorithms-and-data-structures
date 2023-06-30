// Trova i cammini minimi tra ogni coppia di vertici

// Supporta pesi negativi, non supporta cicli

package main

import (
	"fmt"
	"math"
)

type vertice struct {
	str string
	val float64
}

type arco struct {
	x, y string
	val  float64
}

type grafo []arco

func numeroVertici(g grafo) int {
	vertici := make(map[string]bool)
	for _, v := range g {
		if !vertici[v.x] {
			vertici[v.x] = true
		}
		if !vertici[v.y] {
			vertici[v.y] = true
		}
	}
	return len(vertici)
}

func contiene(g grafo, i, j int) (bool, float64) {
	x, y := string('A'+i), string('A'+j)
	for _, v := range g {
		if x == v.x && y == v.y {
			return true, v.val
		}
	}
	return false, -1
}

func creaMatrici(g grafo) ([][]float64, [][]float64, int) {
	n := numeroVertici(g)
	costi := make([][]float64, n)
	pesi := make([][]float64, n)

	for i := 0; i < n; i++ {
		costi[i] = make([]float64, n)
		pesi[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			if i == j {
				costi[i][j] = 0
			} else if ok, v := contiene(g, i, j); ok == true {
				costi[i][j] = v
			} else {
				costi[i][j] = math.Inf(1)
			}
			pesi[i][j] = 0
		}
	}
	return costi, pesi, n
}

func floydWarshall(g grafo) ([][]float64, [][]float64) {
	costi, pesi, n := creaMatrici(g)

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if costi[i][k]+costi[k][j] < costi[i][j] {
					costi[i][j] = costi[i][k] + costi[k][j]
					pesi[i][j] = float64(k)
				}
			}
		}
	}
	return costi, pesi
}

func main() {

	g := grafo{
		{"A", "B", 5},
		{"B", "C", 7},
		{"C", "B", 1},
		{"D", "C", 2},
		{"D", "A", 2},
		{"B", "D", 4},
	}

	fmt.Println(floydWarshall(g))

}
