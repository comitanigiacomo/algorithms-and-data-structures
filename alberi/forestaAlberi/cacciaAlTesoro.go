//MODELLAZIONE:

/*' possibile modellare la situazione mediante una foresta di alberi binari radicati
i nodi dell'albero  rappresentano le informazioni associate ai foglietti, trovare la soluzione del problema significa trovare il valore del nodo radice dell'albero, ovvero il valore misterioso del primo foglietto
le foglie della foreata rappresentano gli oggetti
in questa implementazione non ci possono essere due foglietti con lo stesso nome, e un foglietto non puo' avere un valore negativo

1) dato il nome di un oggetto calcolarne il numero misterioso:

	data la chiave di un nodo della foresta, trovarne il valore

2) dato il nome di un oggetto, scrivere l'espressione che riassume le operazioni da fare per calcolare il numero misterioso associato all'oggetto

	data la chiave di un nodo, scrivere l'espressione che riassume le operazioni da fare per calcolare il valore del nodo

3) stampare in ordine crescente, tutti i numeri sui biglietti

	stampare in ordine crescente i valori dei nodi

4) dato il nome di un oggetto a , stampare i nomi degli oggetti che servono ad a

	data la chiave di un nodo a, stampare le chiavi dei nodi che servono ad a

e' opportuno implementare i nodi dell'albero mediante una struct oggetto avente i seguenti campi :

	type oggetto struct {
		chiave string
		sx     string
		dx    string
		val    int
		op     string
	}

	sx e dx sono due puntatori a oggetto che rappresentano rispettivamente il figlio sinistro e il figlio destro del nodo
	il campo val contiene il numero se il nodo e' una foglia, -1 altrimenti
	il campo op contiene una stringa vuota se il nodo e' una foglia, l'operazione da effettuare altrimenti

la foresta di alberi puo' essere implementata mediante una struct avente i seguenti campi :

	type foresta struct {
		padri map[string]*oggetto
		nodi map[string]*oggetto
	}

la mappa padri conterra' per ogni nodo il suo padre, solo se presente
la mappa nodi conterra' per ogni chiave di un nodo, il puntatore al nodo che lo rappresenta


1) dato il nome di un oggetto calcolarne il numero misterioso:

	data la chiave di un nodo, lanciare una funzione ricorsiva che calcola il suo valore

2) dato il nome di un oggetto, scrivere l'espressione che riassume le operazioni da fare per calcolare il numero misterioso associato all'oggetto

	data la chiave di un oggetto lanciare una funzione ricorsiva che ritorna una stringa contenente le operazioni necessarie a calcolare il valore del nodo, la ricorsione si ferma
	nel momento in cui si analizza un nodo foglia della foresta

3) stampare in ordine crescente, tutti i numeri sui biglietti

	stampare in ordine crescente mediante una visita in mid-order i valori dei nodi foglia della foresta

4) dato il nome di un oggetto a , stampare i nomi degli oggetti che servono ad a

	data la chiave di un nodo a , lanciare una funzione ricorsiva che stampa le chiave dei nodi che servono ad a e si ferma nel caso in cui il nodo analizzato sia una foglia

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f := costruisciForesta(leggi())
	stampaForesta(f)
	fmt.Println(calcolaPrezzo(f, "letto"))
	fmt.Println(espressione(f, "letto"))
	stampaCrescente(f)
	miServe(f, "letto")
}

type oggetto struct {
	chiave string
	sx     string
	dx     string
	val    int
	op     string
}

type foresta struct {
	padri map[string]*oggetto
	nodi  map[string]*oggetto
}

func leggi() map[string]*oggetto {
	nodi := make(map[string]*oggetto)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		if len(args) == 2 {
			num, _ := strconv.Atoi(args[1])
			nodi[args[0][:len(args[0])-1]] = &oggetto{args[0][:len(args[0])-1], "", "", num, ""}
		} else {
			nodi[args[0][:len(args[0])-1]] = &oggetto{args[0][:len(args[0])-1], args[1], args[3], -1, args[2]}
		}
	}
	fmt.Println(nodi)
	return nodi
}

func costruisciForesta(nodi map[string]*oggetto) foresta {
	var f foresta
	padri := make(map[string]*oggetto)
	f.nodi = nodi
	for _, o := range f.nodi {
		if o.sx != "" {
			padri[o.sx] = o
		}
		if o.dx != "" {
			padri[o.dx] = o
		}
	}
	f.padri = padri
	return f
}

func stampaForesta(f foresta) {
	for s := range f.nodi {
		_, ok := f.padri[s]
		if ok {
			fmt.Println(s, f.padri[s].chiave)
		} else {
			fmt.Println(s)
		}
	}
}

func calcolaPrezzo(f foresta, nome string) int {
	ogg := f.nodi[nome]
	var ris int
	if ogg.sx == "" && ogg.dx == "" {
		return ogg.val
	}
	switch ogg.op {
	case "*":
		ris = calcolaPrezzo(f, ogg.sx) * calcolaPrezzo(f, ogg.dx)
	case "/":
		ris = calcolaPrezzo(f, ogg.sx) / calcolaPrezzo(f, ogg.dx)
	case "+":
		ris = calcolaPrezzo(f, ogg.sx) + calcolaPrezzo(f, ogg.dx)
	case "-":
		ris = calcolaPrezzo(f, ogg.sx) - calcolaPrezzo(f, ogg.dx)
	}
	return ris
}

func espressione(f foresta, nome string) string {
	ogg := f.nodi[nome]
	if ogg.sx == "" && ogg.dx == "" {

		return strconv.Itoa(ogg.val)
	}
	return "( " + espressione(f, ogg.sx) + " " + ogg.op + " " + espressione(f, ogg.dx) + " )"
}

func stampaCrescente(f foresta) []int {
	var foglie []int
	for _, v := range f.nodi {
		if v.sx == "" && v.dx == "" {
			foglie = append(foglie, v.val)

		}
	}
	sort.Ints(foglie)
	fmt.Println(foglie)
	return foglie
}

func miServe(f foresta, nome string) {
	ogg := f.nodi[nome]
	if ogg.sx == "" && ogg.dx == "" {
		return
	}
	fmt.Print(ogg.sx, " ")
	fmt.Print(ogg.dx, " ")
	miServe(f, ogg.sx)
	miServe(f, ogg.dx)
}
