package main

import (
	"fmt"
	"math"
)

func main() {
	var N int

	fmt.Scan(&N)

	vetor := make([]int, N)

	menor1, menor2, menor3 := math.MaxInt32, math.MaxInt32, math.MaxInt32
	x1, x2, x3 := -1, -1, -1

	for i := 0; i < N; i++ {
		fmt.Scan(&vetor[i])
	}

	for i := 0; i < N; i++ {
		if vetor[i] < menor1 {
			menor1 = vetor[i]
			x1 = i
		}
	}

	for i := 0; i < N; i++ {
		if vetor[i] < menor2 && vetor[i] != menor1 {
			menor2 = vetor[i]
			x2 = i
		}
	}

	for i := 0; i < N; i++ {
		if vetor[i] < menor3 && vetor[i] != menor1 && vetor[i] != menor2 {
			menor3 = vetor[i]
			x3 = i
		}
	}

	fmt.Printf("Funcionario %d : %d meses\n", x1, menor1)
	fmt.Printf("Funcionario %d : %d meses\n", x2, menor2)
	fmt.Printf("Funcionario %d : %d meses\n", x3, menor3)
}
