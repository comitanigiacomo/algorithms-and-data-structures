/*
A B 3
E F 1
A C 2
B C 4
B F 5
C F 5
B D 6
D E 6
*/

package main

import "fmt"

type arco struct {
	x, y  string
	value int
}

type grafo []arco

func ordinaArchi(g grafo) grafo {
	for i := 1; i < len(g); i++ {
		if g[i].value < g[i-1].value {
			g[i-1], g[i] = g[i], g[i-1]
			i = 0
		}
	}
	return g
}

// Verifica se due nodi sono connessi nel grafo

func BFS(g grafo, start, target string) bool {
	queue := []string{start}
	visited := make(map[string]bool)

	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == target {
			return true
		}

		for _, arco := range g {
			if arco.x == node && !visited[arco.y] {
				queue = append(queue, arco.y)
				visited[arco.y] = true
			}
			if arco.y == node && !visited[arco.x] {
				queue = append(queue, arco.x)
				visited[arco.x] = true
			}
		}
	}

	return false
}

func kruskal(g grafo) grafo {
	tree := []arco{}

	g = ordinaArchi(g)

	for i := 0; i < len(g); i++ {
		if !BFS(tree, g[i].x, g[i].y) {

			tree = append(tree, g[i])

		}
	}
	return tree
}

func main() {
	g := grafo{
		{"E", "F", 1},
		{"A", "C", 2},
		{"A", "B", 3},
		{"B", "C", 4},
		{"B", "F", 5},
		{"C", "F", 5},
		{"B", "D", 6},
		{"D", "E", 6},
	}

	alberoRicoprente := kruskal(g)

	fmt.Println("Albero Ricoprente Minimo:")
	for _, arco := range alberoRicoprente {
		fmt.Println(arco)
	}
}
