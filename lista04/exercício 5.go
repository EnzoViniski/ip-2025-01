package main

import "fmt"

func main(){
	v := [10]int{}
	menor := 100000
	posicao := -1
	for i := 0; i < 10; i++{
		fmt.Scan(&v[i])
		if v[i] < menor {
			menor = v[i]
			posicao = i
		}

	}

	fmt.Println("O menor elemento do vetor é", menor)
	fmt.Print("e sua posição dentro do vetor é:", posicao)
}
