package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stampaMatrice(creaMatrx(leggi()))
}

// crea una mappa contenente per ogni jet
func leggi() (map[int]int, int, []int) {
	var s []int
	mappa := make(map[int]int)
	scanner := bufio.NewScanner(os.Stdin)
	var peso int
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		if len(args) == 1 {
			peso, _ = strconv.Atoi(args[0])
		} else {
			num1, _ := strconv.Atoi(args[0])
			num2, _ := strconv.Atoi(args[1])
			s = append(s, num1)
			mappa[num1] = num2
		}
	}
	return mappa, peso, s
}

// utilizzando la tecnica della programmazione dinamica costruisce una matrice contenente per ogni riga la soluzione ottima usando un certo quantitativo di oggetti
func creaMatrx(pesi map[int]int, peso int, riga []int) [][]int {
	intervallo := pesi[riga[1]] - pesi[riga[0]]
	var colonna []int
	trovaColonna := make(map[int]int)
	var count int
	for i := pesi[riga[0]]; i <= peso; i += intervallo {
		colonna = append(colonna, i)
		trovaColonna[i] = count
		count++
	}
	matrx := make([][]int, len(riga))
	for i := 0; i < len(riga); i++ {
		matrx[i] = make([]int, len(colonna))
	}
	for i := 0; i < len(riga); i++ {
		for j := 0; j < len(colonna); j++ {
			if colonna[j]-pesi[riga[i]] >= 0 {
				matrx[i][j] += riga[i]
			}
			if colonna[j]-pesi[riga[i]] >= pesi[riga[0]] {
				if i > 0 {
					matrx[i][j] += matrx[i-1][trovaColonna[colonna[j]-pesi[riga[i]]]]
				}
			}
			if i > 0 {
				if matrx[i][j] <= matrx[i-1][j] {
					matrx[i][j] = matrx[i-1][j]
				}
			}
		}
	}
	return matrx
}

// stampa una matrice
func stampaMatrice(matrx [][]int) {
	for i := 0; i < len(matrx); i++ {
		fmt.Println(matrx[i])
	}
}
