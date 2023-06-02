package main

import "fmt"

type oggetto struct {
	peso, valore int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func zaino(pesoMassimo int, oggetti map[int]oggetto) int {
	n := len(oggetti)
	matrice := make([][]int, n+1)
	for i := range matrice {
		matrice[i] = make([]int, pesoMassimo+1)
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= pesoMassimo; j++ {
			if oggetti[i].peso > j {
				matrice[i][j] = matrice[i-1][j]
			} else {
				matrice[i][j] = max(matrice[i-1][j], oggetti[i].valore+matrice[i-1][j-oggetti[i].peso])
			}
		}
	}

	return matrice[n][pesoMassimo]
}

func main() {
	oggetti := []oggetto{
		{peso: 3, valore: 3},
		{peso: 2, valore: 4},
		{peso: 5, valore: 4},
		{peso: 2, valore: 1},
		{peso: 6, valore: 7},
	}

	pesoMassimo := 10

	oggettiMap := make(map[int]oggetto)
	for i, obj := range oggetti {
		oggettiMap[i+1] = obj
	}

	fmt.Println(zaino(pesoMassimo, oggettiMap))
}
