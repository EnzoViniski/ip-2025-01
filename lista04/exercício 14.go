package main

import "fmt"

func main() {

	vetor1 := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Scan(&vetor1[i])

	}

	vetor2 := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Scan(&vetor2[i])

	}

	for i := 0; i < 10; i++ {
		fmt.Print(vetor1[i], " ", vetor2[i], " ")

	}
}
