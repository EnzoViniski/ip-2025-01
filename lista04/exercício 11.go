package main

import (
	"fmt"
)

func main() {
	vetor := make([]float64, 100)

	for i := 0; i < 100; i++ {
		fmt.Println("Digite um valor: ")
		fmt.Scan(&vetor[i])
	}

	S := 0.0

	for i := 0; i < 50; i++ {
		S += (vetor[i] - vetor[99-i]) * (vetor[i] - vetor[99-i]) * (vetor[i] - vetor[99-i])
	}

	fmt.Println("O resultado é:", S)
}
