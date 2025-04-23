package main

import (
	"fmt"
)

func main() {
	var vetor1 [10]int
	var pares []int
	var impares []int

	somap := 0
	quanti := 0

	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&vetor1[i])
		if vetor1[i]%2 == 0 {
			pares = append(pares, vetor1[i])
			somap += vetor1[i]
		} else {
			impares = append(impares, vetor1[i])
			quanti += 1
		}
	}

	fmt.Println("a)", pares)
	fmt.Println("b)", somap)
	fmt.Println("c)", impares)
	fmt.Println("d)", quanti)

}
