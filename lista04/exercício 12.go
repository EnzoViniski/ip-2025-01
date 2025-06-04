package main

import (
	"fmt"
)

func main() {
	vetor := make([]int, 15)
	freq := make([]int, 15)

	freqr := make([]int, 11)

	for i := 0; i < 15; i++ {
		fmt.Printf("Digite a nota do aluno %d: ", i+1)
		_, err := fmt.Scan(&vetor[i])
		if err != nil {
			fmt.Println("Erro na leitura:", err)
			return
		}
	}

	for i := 0; i < 15; i++ {
		freq[i] = 0
		for j := 0; j < 15; j++ {
			if vetor[j] == vetor[i] {
				freq[i]++
			}
		}
	}

	fmt.Println("NOTA -- FREQUENCIA ABSOLUTA -- FREQUENCIA RELATIVA")
	fmt.Println("--------------------------------------------------")

	for i := 0; i < 11; i++ {
		freqr[i] = freq[i] / 11
		fmt.Println(" ", i, "       ", "  ", freq[i], "                 ", freqr[i])
	}
}
