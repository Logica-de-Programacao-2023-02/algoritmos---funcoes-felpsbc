package main

import (
	"errors"
	"fmt"
)

func numeroSlice(numero int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New("O slice está vazio")
	}

	for _, valor := range slice {
		if valor == numero {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	numero := 5
	slice := []int{1, 2, 3, 4, 5}

	encontrado, err := numeroSlice(numero, slice)

	if err != nil {
		fmt.Println(err)
	} else if encontrado {
		fmt.Printf("%d está presente no slice.\n", numero)
	} else {
		fmt.Printf("%d não está presente no slice.\n", numero)
	}
}
