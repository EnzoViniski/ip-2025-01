package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64

	vetor := make([]float64, 15)

	for i := 0; i < 15; i++ {
		fmt.Println("Digite um número: ")
		fmt.Scan(&x)

		if x < 0 {
			vetor[i] = -1
		} else {
			x = math.Sqrt(x)
			vetor[i] = x
		}
	}

	for i := 0; i < 15; i++ {
		fmt.Println(vetor[i])
	}
}
