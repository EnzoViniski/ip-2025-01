package main

import (
	"fmt"
)

func main() {
	var T int

	vetor := make([]int, 24)
	vetor2 := make([]int, 24)

	fmt.Println("Janela:")
	for i := 0; i < 24; i++ {
		fmt.Printf("A poltrona %d está ocupada? (1-sim, 0-nao)\n", i+1)
		fmt.Scan(&vetor[i])
	}

	fmt.Println("Corredor:")
	for i := 0; i < 24; i++ {
		fmt.Printf("A poltrona %d está ocupada? (1-sim, 0-nao)\n", i+1)
		fmt.Scan(&vetor2[i])
	}

	z := 0
	for i := 0; i < 24; i++ {
		if vetor[i] == 0 || vetor2[i] == 0 {
			z++
		}
	}
	if z == 0 {
		fmt.Println("Todas as poltronas estão ocupadas")
	} else {
		fmt.Println("Qual deseja? (1-Janela, 2-Corredor)")
		fmt.Scan(&T)

		escolha := 0
		y := 0
		if T == 1 {
			fmt.Println("Escolha uma das poltronas:")
			for i := 0; i < 24; i++ {
				if vetor[i] == 0 {
					fmt.Print(i+1, " ")
					y++
				}
			}
			if y > 0 {
				fmt.Scan(&escolha)
				vetor[escolha-1] = 1
			} else {
				fmt.Println("Nenhuma poltrona esta disponivel")
			}
		}
		x := 0
		if T == 2 {
			fmt.Println("Escolha uma das poltronas:")
			for i := 0; i < 24; i++ {
				if vetor2[i] == 0 {
					fmt.Print(i+1, " ")
					x++
				}
			}
			if y > 0 {
				fmt.Scan(&escolha)
				vetor2[escolha-1] = 1
			} else {
				fmt.Println("Nenhuma poltrona esta disponivel")
			}
		}
	}
}
