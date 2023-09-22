package main

import (
	"errors"
	"fmt"
	"sort"
)

func ordenarCrescente(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	ordenado := make([]int, len(slice))
	copy(ordenado, slice)
	sort.Ints(ordenado)

	return ordenado, nil
}

func main() {
	numeros := []int{5, 3, 1, 4, 2}
	ordenados, err := ordenarCrescente(numeros)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ordenados)
	}
}
