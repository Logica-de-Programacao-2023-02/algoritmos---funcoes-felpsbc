package main

import (
	"errors"
	"fmt"
)

func substituirVogais(texto string) (string, error) {
	if texto == "" {
		return "", errors.New("A string está vazia")
	}

	vogais := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}

	runes := []rune(texto)

	for i, char := range runes {
		if vogais[char] {
			runes[i] = '*'
		}
	}
	novaString := string(runes)

	return novaString, nil
}

func main() {
	texto := "Aula de Programacao"
	novo, err := substituirVogais(texto)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(novo)
	}
}
