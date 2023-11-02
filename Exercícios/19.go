package main

import (
	"errors"
	"fmt"
)

func Primo(numero int) bool {
	if numero <= 1 {
		return false
	}
	if numero <= 3 {
		return true
	}
	if numero%2 == 0 || numero%3 == 0 {
		return false
	}

	for i := 5; i*i <= numero; i += 6 {
		if numero%i == 0 || numero%(i+2) == 0 {
			return false
		}
	}

	return true
}

func numerosPrimos(numero int) ([]int, error) {
	if numero < 0 {
		return nil, errors.New("O número é negativo")
	}

	primos := make([]int, 0)
	for i := 2; i <= numero; i++ {
		if Primo(i) {
			primos = append(primos, i)
		}
	}

	return primos, nil
}

func main() {
	numero := 20
	resultado, err := numerosPrimos(numero)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Números primos até", numero, ":", resultado)
	}
}
