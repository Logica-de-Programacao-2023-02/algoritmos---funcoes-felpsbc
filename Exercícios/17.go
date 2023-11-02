package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func ordenarStrings(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("O slice está vazio")
	}

	sort.Strings(slice)

	novaString := strings.Join(slice, " ")

	return novaString, nil
}

func main() {
	slice := []string{"felipe", "guilherme", "otavio", "enzo"}
	novaString, err := ordenarStrings(slice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(novaString)
	}
}
