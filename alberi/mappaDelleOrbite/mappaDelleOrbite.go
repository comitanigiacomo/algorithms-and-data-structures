package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	mappa := leggi()
	fmt.Println(contaOrbiteTotali(mappa))
	s := pathYou(mappa)
	fmt.Println(youToSan(mappa, s))
}

//implemento un albero mediante la tabella dei padri
// l'albero è rappresentato quindi da un vettore i cui elementi contengono il valore associato al nodo e l'indice della posizione del padre nel vettore
// in go posso rappresentare questo vettore mediante una mappa da string a string, in modo tale da ottenere immediatamente per ogni oggetto il padre associato mappa[elemento]padre
//le orbite dirette sono i collegamenti padre figlio all'interno dell'albero
// le orbite indirette sono i collegamenti che vanno da una radice ad un insieme di figli escluso quello a cui si arriva con un orbita diretta
// il numero di orbite indirette di un oggetto è il numero di nodi che sipossono raggiungere a partire da una radice eslusi i figli diretti di quella radice nell'albero

//per calcolare il numero di orbite dirette e indirette all'interno dell'albero, posso :

//1) costruire un albero implementato mediante la tabella dei padri

//2) per ogni nodo dell'albero guardo il padre e incremento un contatore, itero il procedimento fino a quando arrivo alla radice
// in questo modo ottengo il numero delle orbite indirette o dirette

//per trovare il numero di spostamenti minimi per andare da YOU a SAN :

//1) salvo in un vettore il percorso da YOU a COM
//2) partendo da SAN guardo ogni volta il padre ed implemento un contatore fino a quando non arrivo ad una stringa contenuta nel percorso di YOU
// 3) partedo da quella stringa scansiono al contrario l'array del percorso di YOU fino alla fine dell'array, incrementando ogni volta il contatore

//funzione che legge in input una serie di orbite e ritorna una mappa da string a string map[elemento]padre
func leggi() map[string]string {
	mappa := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), ")")
		mappa[args[1]] = args[0]
	}
	return mappa
}

// funzione che data in input una mappa ritorna il numero di orbite dirette e indirette totali
func contaOrbiteTotali(mappa map[string]string) (count int) {
	for elemento := range mappa {
		count += contaOrbite(mappa, elemento)
	}
	return count
}

// funzione ricorsiva che dato un elemento della mappa comta le sue orbite dirette e indirette
func contaOrbite(mappa map[string]string, s string) int {
	_, ok := mappa[s]
	if !ok {
		return 0
	}
	return 1 + contaOrbite(mappa, mappa[s])
}

// ritorna un array che contiene il percorso oer andare da YOU a COM
func pathYou(mappa map[string]string) (s []string) {
	str := "YOU"
	for str != "COM" {
		tmp := mappa[str]
		s = append(s, tmp)
		str = tmp
	}
	return s
}

// ritorna true se s contiene str, false altrimenti
func contiene(s []string, str string) bool {
	for _, v := range s {
		if str == v {
			return true
		}
	}
	return false
}

// calcola il numero minimo di orbite per andare da YOU a SAN
func youToSan(mappa map[string]string, s []string) int {
	var count int
	var flag bool
	tmp := mappa["SAN"]
	for !contiene(s, tmp) {
		count++
		str := tmp
		tmp = mappa[str]
	}
	for i := len(s) - 1; i >= 0; i-- {
		if flag {
			count++
		}
		if s[i] == tmp {
			flag = true
		}
	}
	return count
}
