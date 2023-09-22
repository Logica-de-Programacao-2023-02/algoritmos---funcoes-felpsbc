package main

import (
	"errors"
	"fmt"
)

func numerosPares(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	var pares []int
	for _, valor := range slice {
		if valor%2 == 0 {
			pares = append(pares, valor)
		}
	}

	return pares, nil
}

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6}
	pares, err := numerosPares(numeros)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pares)
	}
}
