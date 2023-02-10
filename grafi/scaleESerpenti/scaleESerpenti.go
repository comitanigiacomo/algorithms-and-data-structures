package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(calcolaMinimo(leggi(), 5))
}
func leggi() []int {
	scanner := bufio.NewScanner(os.Stdin)
	var arr []int
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		num1, _ := strconv.Atoi(args[0])
		num2, _ := strconv.Atoi(args[1])
		if len(arr) == 0 {
			for i := 0; i < num1*num2; i++ {
				arr = append(arr, 0)
			}
		}
		arr[num1-1] = num2 - 1
	}
	return arr
}

func calcolaMinimo(arr []int, partenza int) ([]int, int) {
	mappa := make(map[int]int)
	mappina := make(map[int]int)
	coda := []int{partenza}
	for len(coda) != 0 {
		tmp := coda[0]
		coda = coda[1:]
		if tmp == 29 {
			break
		}
		for i := 1; i < 7; i++ {
			pos := tmp + i
			if pos < 30 {
				if arr[pos] != 0 {
					pos = arr[pos]
				}
				_, ok := mappa[pos]
				if !ok {
					coda = append(coda, pos)
					mappa[pos] = tmp
					mappina[pos] = i
				}
			}
		}
	}
	fmt.Println(mappa)
	fmt.Println(mappina)
	lanci := []int{}
	elem := len(arr) - 1
	for mappa[elem] != 0 {
		lanci = append(lanci, mappina[elem])
		elem = mappa[elem]
	}
	lanci = append(lanci, mappina[elem])
	for i := 0; i < len(lanci)/2; i++ {
		lanci[i], lanci[len(lanci)-1-i] = lanci[len(lanci)-1-i], lanci[i]
	}
	return lanci, len(lanci)
}
