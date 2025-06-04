package main

import (
	"fmt"
)

func main() {
	vetor := make([]float64, 10)

	soma := 0.0

	for i := 0; i < 10; i++ {
		fmt.Println("Digite a altura do jogador:")
		fmt.Scan(&vetor[i])
		soma += vetor[i]
	}

	media := soma / 10.0

	for i := 0; i < 10; i++ {
		if vetor[i] > media {
			fmt.Println(vetor[i])
		}
	}
}
