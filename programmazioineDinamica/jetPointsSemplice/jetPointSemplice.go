//knapsack problem
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	oggetti, peso := leggi()
	fmt.Println(riempiZaino(peso, oggetti))
}

type oggetto struct {
	peso  int
	value int
}

func leggi() ([]*oggetto, int) {
	scanner := bufio.NewScanner(os.Stdin)
	var peso int
	var oggetti []*oggetto
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		if len(args) == 1 {
			peso, _ = strconv.Atoi(args[0])
		} else {
			newoggetto := new(oggetto)
			newoggetto.peso, _ = strconv.Atoi(args[1])
			newoggetto.value, _ = strconv.Atoi(args[0])
			oggetti = append(oggetti, newoggetto)
		}
	}
	return oggetti, peso
}

func riempiZaino(peso int, oggetti []*oggetto) int {
	zaino := make([]int, peso+1)
	for i := 0; i < len(zaino); i++ {
		for j := 0; j < len(oggetti); j++ {
			if i+oggetti[j].peso <= peso && zaino[i]+oggetti[j].value > zaino[i+oggetti[j].peso] {
				zaino[i+oggetti[j].peso] = zaino[i] + oggetti[j].value
			}
		}
	}
	return (zaino[len(zaino)-1])
}
