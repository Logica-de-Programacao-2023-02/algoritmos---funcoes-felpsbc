package main

import (
	"fmt"
)

func calcularMedia(slice []int) float64 {
	if len(slice) == 0 {
		return 0
	}

	soma := 0
	for _, valor := range slice {
		soma += valor
	}
	media := float64(soma) / float64(len(slice))
	return media
}

func main() {

	numeros := []int{10, 20, 30, 40, 50}

	media := calcularMedia(numeros)

	fmt.Printf("A média dos valores é: %.2f\n", media)
}
