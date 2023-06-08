/*
Sono nella situazione in cui ho uno "zaino" dove posso inserire degli oggetti di diversi tipi, e ho a disposizione un quantità illimitata di  oggetti.

posso quindi risolvere il problema utilizzando l'approccio del knapsack problem semplificato

Rappresento lo zaino come un array dove in ogni posizione è presente il numero di swindles massimo ottenibile in base al numero massimo di jet points supportato
*/

package main

import (
	"fmt"
)

type oggetto struct {
	jetPoints, valore int
}

func calcolaValoreMassimo(jetPoints int, tabella map[int]int) int {
	arr := make([]int, jetPoints+1)

	for i := 1; i <= jetPoints; i++ {
		for valore, jetPointsReq := range tabella {
			if jetPointsReq <= i {
				arr[i] = max(arr[i], valore+arr[i-jetPointsReq])
			}
		}
	}

	return arr[jetPoints]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	tabella := map[int]int{
		100: 20000,
		107: 22000,
		124: 24000,
		133: 26000,
		139: 28000,
		155: 30000,
		172: 32000,
		178: 34000,
		184: 36000,
		190: 38000,
		195: 40000,
	}

	jetPoints := 140000
	valoreMassimo := calcolaValoreMassimo(jetPoints, tabella)
	fmt.Println("Valore massimo in swindle:", valoreMassimo)
}
