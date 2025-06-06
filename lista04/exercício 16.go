package main

import "fmt"

func main() {

	vetor := make([]int, 50)

	for i := 0; i < 50; i++ {
		fmt.Scan(&vetor[i])
	}

	x := make([]int, 50)
	moda := -1
	maior := 0
	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			if vetor[i] == vetor[j] {
				x[i]++
			}
		}
		if x[i] > maior {
			moda = vetor[i]
			maior = x[i]
		}
	}

	fmt.Print("A moda é: ", moda)

}
