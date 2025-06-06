package main

import (
	"fmt"
)

func main() {
	vetor := make([]int, 10)
	vetor2 := make([]int, 5)

	for i := 0; i < 10; i++ {
		fmt.Scan(&vetor[i])
	}

	for i := 0; i < 5; i++ {
		fmt.Scan(&vetor2[i])
	}

	for i := 0; i < 10; i++ {
		fmt.Println("Número ", vetor[i], ":")
		for j := 0; j < 5; j++ {
			teste := vetor[i] % vetor2[j]
			if teste == 0 {
				fmt.Printf("Divisível por %d na posição %d\n", vetor2[j], j)
			}
		}
	}

}
