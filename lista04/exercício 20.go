package main

import (
	"fmt"
)

func main() {
	vetor := make([]int, 20)

	for i := 0; i < 20; i++ {
		fmt.Scan(&vetor[i])
	}

	x := make([]int, 20)

	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			if vetor[i] == vetor[j] {
				x[i]++
			}
		}
		fmt.Println("Número ", vetor[i], " sorteado ", x[i], " vezes.")
	}

	for i := 0; i < 20; i++ {
		fmt.Scan(&vetor[i])
	}

}
