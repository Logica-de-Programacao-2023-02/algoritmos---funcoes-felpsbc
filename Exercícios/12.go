package main

import (
	"errors"
	"fmt"
	"strings"
)

func Maiuscula(nomes []string) (string, error) {
	if len(nomes) == 0 {
		return "", errors.New("O slice está vazio!")
	}
	var resultado []string

	for _, nome := range nomes {
		if len(nomes) > 0 && strings.ToUpper(nome[0:1]) == nome[0:1] {
			resultado = append(resultado, nome)
		}
	}
	rs := strings.Join(resultado, " ")
	return rs, nil
}

func main() {
	var nomes = []string{"Felipe", "Maria", "joao", "enzo"}
	resultado, err := Maiuscula(nomes)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}
}
