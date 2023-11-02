package main

import (
	"errors"
	"fmt"
	"strconv"
)

func Soma(num int) (int, error) {
	if num <= 0 {
		return 0, errors.New("O número não pode ser negativo ou igual a 0")
	}

	soma := 0

	numstring := strconv.Itoa(num)

	for _, char := range numstring {
		digito, err := strconv.Atoi(string(char))
		if err != nil {
			return 0, err
		}
		soma += digito
	}
	return soma, nil
}

func main() {
	num := 6666
	resultado, err := Soma(num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("A soma dos dígitos é: %d\n", resultado)
	}
}
