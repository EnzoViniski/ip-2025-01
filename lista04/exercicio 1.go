package main

import "fmt"

func main() {
	vetorNum := []int{20, 55, 60, 22, 44, 23, 52, 51, 35, 90}

	valor := 0

	for i := 0; i < 10; i++ {
		if vetorNum[i] > 50 {
			fmt.Println(vetorNum[i], " na posição ", i)
			valor = 1
		}
	}

	if valor == 0 {
		fmt.Println("Não há valores maiores que 50")
	}
}
