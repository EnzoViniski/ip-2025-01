package main

import (
	"fmt"
)

func main() {
	var vetor1 [10]int
	var vetor2 [5]int
	var pares []int
	var impares []int

	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&vetor1[i])
		if vetor1[i]%2 == 0 {
			pares = append(pares, vetor1[i])
		} else {
			impares = append(impares, vetor1[i])
		}
	}

	fmt.Println("Digite 5 números inteiros para o segundo vetor:")
	soma := 0
	for i := 0; i < 5; i++ {
		fmt.Scan(&vetor2[i])
		soma += vetor2[i]
	}

	var resultadoPares []int
	var resultadoImpares []int

	for _, valor := range pares {
		resultadoPares = append(resultadoPares, valor+soma)
	}

	for _, valor := range impares {
		resultadoImpares = append(resultadoImpares, valor+soma)
	}

	fmt.Println("Resultado dos pares:", resultadoPares)
	fmt.Println("Resultado dos ímpares:", resultadoImpares)
}
