package main

import (
	"fmt"
)

func main() {
	vetor := make([]int, 50)

	a1 := 0
	a2 := 1
	for i := 0; i < 50; i++ {
		if i == 0 {
			vetor[i] = 1
		} else {
			vetor[i] = a1 + a2
			a1 = a2
			a2 = vetor[i]
		}
	}

	for i := 0; i < 50; i++ {
		fmt.Println(vetor[i])
	}

}
