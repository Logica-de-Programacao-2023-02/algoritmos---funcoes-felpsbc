package main

import (
	"fmt"
	"strings"
)

func main() {
	x := "Olá, Mundo!"
	quantidadeVogais := contarVogais(x)
	fmt.Printf("A quantidade de vogais em %s é: %d\n", x, quantidadeVogais)
}

func contarVogais(s string) int {
	s = strings.ToLower(s)
	vogais := "aeiou"
	count := 0

	for _, char := range s {
		if strings.ContainsAny(vogais, string(char)) {
			count++
		}
	}

	return count
}
