package main

import "fmt"

func main (){
	vetor := make([]int, 10)

	soma := 0
	for i := 0; i < 10; i++{
		fmt.Scan(&vetor[i])
	}

	fmt.Println("Pares: ")

	for i := 0; i < 10; i++{
		teste := vetor[i] % 2
		if teste == 0 {
			fmt.Println(vetor[i])
			soma += vetor[i]
		}
	}

	fmt.Println("Soma pares: ", soma)

	soma = 0

	fmt.Println("Ímpares:")

	for i := 0; i < 10; i++{
		teste := vetor[i] % 2
		if teste != 0 {
			fmt.Println(vetor[i])
		    soma ++
		}
	}
	
	fmt.Println("Quantidade ímpares: ", soma)

}
