package main

import (
	"fmt"
	"sort"
)

func main() {

	vetor := make([]int, 10)

	for i := 0; i < 10; i++ {
		fmt.Scan(&vetor[i])
	}

	sort.Slice(vetor, func(i, j int) bool {
		return vetor[i] < vetor[j]
	})

	fmt.Print(vetor)

}
