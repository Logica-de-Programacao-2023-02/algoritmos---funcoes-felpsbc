package main

import (
	"errors"
	"fmt"
	"strings"
)

func ConcatenarStrings(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("O slice está vazio")
	}
	resultado := strings.Join(slice, ", ")
	return resultado, nil
}

func main() {
	x := []string{"Eu ", "Como", " Frango"}
	resultado, err := ConcatenarStrings(x)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}
}
