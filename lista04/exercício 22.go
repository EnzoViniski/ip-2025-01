package main

import (
	"fmt"
)

func main() {
	vetor := make([]int, 10)
	vetor2 := make([]int, 10)

	for i := 0; i < 10; i++ {
		fmt.Println("Digite o número da conta e o saldo inicial:")
		fmt.Scan(&vetor[i], &vetor2[i])
	}

	for {
		for i := 0; i < 10; i++ {
			fmt.Printf("Conta: %d Saldo: %d\n", vetor[i], vetor2[i])
		}
		fmt.Println("\n1. Efetuar Depósito")
		fmt.Println("2. Efetuar Saque")
		fmt.Println("3. Consultar Ativo Bancário")
		fmt.Println("4. Finalizar Programa")

		var opcao int
		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			var conta, valor int
			fmt.Print("Digite o número da conta e o valor do depósito: ")
			fmt.Scan(&conta, &valor)

			encontrou := false
			for j := 0; j < 10; j++ {
				if conta == vetor[j] {
					vetor2[j] += valor
					encontrou = true
					fmt.Println("Depósito realizado com sucesso!")
					break
				}
			}
			if !encontrou {
				fmt.Println("Conta não encontrada!")
			}

		case 2:
			var conta, valor int
			fmt.Print("Digite o número da conta e o valor do saque: ")
			fmt.Scan(&conta, &valor)

			encontrou := false
			for j := 0; j < 10; j++ {
				if conta == vetor[j] {
					encontrou = true
					if valor > vetor2[j] {
						fmt.Println("Saldo insuficiente!")
					} else {
						vetor2[j] -= valor
						fmt.Println("Saque realizado com sucesso!")
					}
					break
				}
			}
			if !encontrou {
				fmt.Println("Conta não encontrada!")
			}

		case 3:
			soma := 0
			for _, saldo := range vetor2 {
				soma += saldo
			}
			fmt.Printf("Ativo bancário total: %d\n", soma)

		case 4:
			fmt.Println("Programa finalizado.")
			return

		default:
			fmt.Println("Opção inválida!")
		}
	}
}
