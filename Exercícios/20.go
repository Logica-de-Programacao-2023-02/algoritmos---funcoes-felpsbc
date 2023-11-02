package main

import (
	"errors"
	"fmt"
)

func Tamanho(slice []string) ([]string, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	resultado := make([]string, 0)

	for _, str := range slice {
		if len(str) > 5 {
			resultado = append(resultado, str)
		}
	}

	return resultado, nil
}

func main() {
	slice := []string{"Felipe", "jogao", "matheus", "aula", "copo", "faculdade"}
	resultado, err := Tamanho(slice)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Strings com mais de 5 caracteres:", resultado)
	}
}
