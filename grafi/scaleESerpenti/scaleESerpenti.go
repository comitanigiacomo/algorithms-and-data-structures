// MODELLAZIONE

/*
Rappresento il tabellone di gioco come un grafo e eseguo una BFS per trovare il percorso minimo dalla tabella di partenza a quella di arrivo.

	- Ogni casella del tabellone rappresenta il nodo di un grafo

	- La connessione tra due caselle rappresenta un arco del grafo. Quindi se la casella 3 è collegata alla casella nove 9 da un arco o da una scala, ci sarà un arco dal nodo 3 al nodo 9


*/

// Implementazione

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// associo ad ogni casella del tabellone tutte quelle che posso raggiungere mediante il lancio di un dado, prendendo scale e serpenti, o non
type graph map[int][]int

// Crea un nuovo grafo
func newGraph(r, c int) graph {
	mappa := make(map[int][]int, (r*c)+1)
	return mappa
}

// Legge da standard input le caratteristiche del grafo, lo crea e ne restituisce l'indirizzo
func readGraph() graph {

	var r, c int
	fmt.Scan(&r, &c)
	g := newGraph(r, c)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := strings.Split(scanner.Text(), " ")
		num1, _ := strconv.Atoi(str[0])
		num2, _ := strconv.Atoi(str[1])
		g[num1] = append(g[num1], num2)
	}
	return g
}

// Trova il percorso migliore per arrivare alla fine del tabellone utilizzando la visita in ampiezza. utilizzo una mappa path per asssociare ad ogni casella, quella da cui effettivamente provengo
func BFS(graph map[int][]int, start, end int) []int {
	coda := []int{start}
	aux := make(map[int]bool)
	aux[start] = true
	path := make(map[int]int) // ad ogni casella associo quella da cui sono arrivato, facendo la BFS

	for len(coda) > 0 {
		nodo := coda[0]
		coda = coda[1:]

		if nodo == end {
			return recuperaPercorso(graph, start, end, path)
		} else {
			// controllo tutte le caselle che posso raggiugere mediante scale o serpenti
			for _, v := range graph[nodo] {
				if !aux[v] {
					aux[v] = true

					coda = append(coda, v)
					path[v] = nodo
				}
			}
			// controllo le caselle che posso raggiungere senza scale
			for i := 1; i <= 6; i++ {
				if nodo+i <= end {
					if !aux[nodo+i] {
						aux[nodo+i] = true

						coda = append(coda, nodo+i)
						path[nodo+i] = nodo
					}
				}
			}
		}
	}
	return []int{}
}

// ritorna un array contenente le caselle del tabellone da cui passo, seguendo il percorso minimo
func recuperaPercorso(g map[int][]int, start, end int, path map[int]int) []int {
	nodo := path[end]
	moves := []int{}

	for nodo != start {
		moves = append(moves, nodo)
		nodo = path[nodo]
	}
	sort.Ints(moves)

	return moves
}

// calcola il numero minimo di lanci necessari del dado, e ritorna un array contenente i numeri che sono usciti lanciando il dado
func calcolaMosseMinime(moves []int, g graph, end, start int) (n int, dado []int) {
	dado = append(dado, moves[0]-start)
	for i := 1; i < len(moves); i++ {
		if moves[i]-moves[i-1] < 6 { // se prendo una scala non conto il dado
			dado = append(dado, moves[i]-moves[i-1])
		}
	}
	dado = append(dado, end-moves[len(moves)-1])
	return len(dado), dado
}

// Calcola il numero minimo di mosse necessarie a vincere la partita senza usare scale ne serpenti, e ritorna un array contenente i lanci del dado
func MosseMinimeSenzascaleoSerpenti(g graph, start, end int) []int {
	var posizione int = 1
	var mossine []int
	var count int

	for posizione != 30 {
		_, ok := g[posizione+6]
		if !ok {
			mossine = append(mossine, 6-count)
			posizione = posizione + 6
			count = 0
		} else {
			posizione--
			count++
		}
	}
	return mossine
}

func main() {
	g := readGraph()
	moves := (BFS(g, 1, 30))
	numero_lanci, lanci := (calcolaMosseMinime(moves, g, 30, 1))
	fmt.Println("---")
	fmt.Println("Il numero di lanci minimo necessari a vincere la partita è: ", numero_lanci)
	fmt.Println("I lanci necessari a vincere la partita sono: ", lanci)
	fmt.Println("I lanci minimi necessari a vincere la partita senza prendere scale o serpenti sono: ", MosseMinimeSenzascaleoSerpenti(g, 1, 30))
}
