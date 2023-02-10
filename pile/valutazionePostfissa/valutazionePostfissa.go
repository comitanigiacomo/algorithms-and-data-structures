package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println(valuta("5 3 - 2 *"))
}

// eventualmente posso anche implementarla mediante una slice, eseguendo push e pop nella posizione più a destra

// ritorna true se la pila è vuota, false altrimenti
/*func isEmpty(p []string) bool {
	return len(p) == 0
}*/

// inserisce un nuovo elemento nella pila
func push(p *[]string, s string) {
	*p = append(*p, s)
}

// estrae l'elemento in cima alla pila
/*func pop(p []string) (string, []string) {
	tmp := p[len(p)-1]
	p = p[:len(p)-1]
	return tmp, p
}*/

// valuta un'espressione in notazione postfissa
func valuta(espressione string) string {
	p := []string{}
	for _, v := range espressione {
		if unicode.IsNumber(v) {
			push(&p, string(v))
		} else if v == '*' || v == '/' || v == '+' || v == '-' {
			num1, _ := strconv.Atoi(p[len(p)-2])
			num2, _ := strconv.Atoi(p[len(p)-1])
			tmp := calcola(num1, num2, v)
			push(&p, string(tmp))
		}
	}
	return p[len(p)-1]
}

func calcola(num1, num2 int, v rune) string {
	var tmp int
	switch v {
	case '*':
		tmp = num1 * num2
	case '/':
		tmp = num1 / num2
	case '+':
		tmp = num1 + num2

	case '-':
		tmp = num1 - num2
	}
	return strconv.Itoa(tmp)
}

// penso sia piu comodo in go implementarla mediante lista concatenata per non avere casini con la pop
