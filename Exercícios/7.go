package main

import (
	"errors"
	"fmt"
)

func exemploDeFuncao(valor int) int {
	return valor * 2
}

func aplicarFuncao(slice []int, funcao func(int) int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}
	resultado := make([]int, len(slice))
	for i, valor := range slice {
		resultado[i] = funcao(valor)
	}
	return resultado, nil
}
func main() {
	numeros := []int{1, 2, 3, 4, 5}
	resultado, err := aplicarFuncao(numeros, exemploDeFuncao)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}
}
