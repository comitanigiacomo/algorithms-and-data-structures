package main

import "fmt"

func main() {

	arr := make([]int, 21)

	oggetti := make(map[int]int)
	oggetti[5] = 2
	oggetti[3] = 3
	oggetti[7] = 8

	fmt.Println(zaino(arr, oggetti))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func zaino(arr []int, oggetti map[int]int) int {
	fmt.Println(arr)
	for i := 1; i < len(arr); i++ {
		for peso, valore := range oggetti {
			if peso <= i {
				arr[i] = max(valore+arr[i-peso], arr[i])
			}
		}
	}
	fmt.Println(arr)
	return arr[len(arr)-1]
}
