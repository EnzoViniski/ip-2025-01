package main

import "fmt"

func main() {

	vetor := make([]int, 30)

	for i := 0; i < 30; i++ {
		fmt.Scan(&vetor[i])
	}

	for i := 0; i < 30; i++ {
		teste := vetor[i] % 2
		if teste == 0 {
			novo := vetor[i] * 2
			fmt.Print(novo, " ")
		} else {
			novo := vetor[i] * 3
			fmt.Print(novo, " ")
		}
	}

}
