package main

import "fmt"

func main() {
	vetor := make([]int, 100)

	x := 1

	for i := 0; i < 100; i++ {
		vetor[i] = x
		x += 2
	}

	for i := 0; i < 100; i++ {
		fmt.Println(vetor[i])
	}
}
