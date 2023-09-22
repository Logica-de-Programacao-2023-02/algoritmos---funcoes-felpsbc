package main

import (
	"errors"
	"fmt"
	"strings"
)

func extrairPalavras(s string) ([]string, error) {
	if s == "" {
		return nil, errors.New("A string está vazia")
	}

	palavras := strings.Fields(s)
	return palavras, nil
}

func main() {
	texto := "Esta é uma string de exemplo"
	palavras, err := extrairPalavras(texto)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(palavras)
	}
}
