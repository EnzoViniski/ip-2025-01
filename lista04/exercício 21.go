package main

import (
	"fmt"
)

func main() {
	var C int

	fmt.Scan(&C)

	vetor := make([]float64, 10)
	for i := 0; i < 10; i++ {
		fmt.Scan(&vetor[i])
	}

	if C == 1 {
		for i := 0; i < 10; i++ {
			fmt.Print(vetor[i], " ")
		}
	}
	if C == 2 {
		for i := 9; i >= 0; i-- {
			fmt.Print(vetor[i], " ")
		}
	}
	if C == 0 {
		return
	}
}
