package main

import (
	"fmt"
)

func encontrarPosicao(slice []int, valor int) int {
	for i, v := range slice {
		if v == valor {
			return i
		}
	}
	return -1
}

func main() {
	numeros := []int{10, 5, 8, 20, 15}
	valor := 8

	posicao := encontrarPosicao(numeros, valor)
	if posicao != -1 {
		fmt.Printf("O valor %d foi encontrado na posição %d.\n", valor, posicao)
	} else {
		fmt.Printf("O valor %d não foi encontrado no slice.\n", valor)
	}
}
