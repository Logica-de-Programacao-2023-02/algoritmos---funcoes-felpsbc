package main

import (
	"errors"
	"fmt"
)

type FuncaoInteiro func(int) int

func aplicar(slice []int, f FuncaoInteiro) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("O slice está vazio")
	}

	resultado := 0
	for _, valor := range slice {
		resultado += f(valor)
	}

	return resultado, nil
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	duplicar := func(x int) int {
		return x * 2
	}
	resultado, err := aplicar(slice, duplicar)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Resultado da soma após aplicar a função:", resultado)
	}
}
