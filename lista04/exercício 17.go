package main

import "fmt"

func main() {

	vetor := make([]int, 10)

	for i := 0; i < 10; i++ {
		fmt.Scan(&vetor[i])
	}

	for i := 0; i < 10; i++ {
		x := 0
		if vetor[i] == 1 {
			fmt.Println("o numero ", vetor[i], " na posição ", i)
		} else {
			for j := 1; j < vetor[i]; j++ {
				teste := vetor[i] % vetor[j]
				if teste == 0 {
					x++
				}
			}
			if x == 1 {
				fmt.Println("o numero ", vetor[i], " na posição ", i)
			}
		}
	}

}
