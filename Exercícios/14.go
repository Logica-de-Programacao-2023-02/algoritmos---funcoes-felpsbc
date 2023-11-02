package main

import (
	"errors"
	"fmt"
)

func PontosComum(slice1, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New("Um dos slices está vazio")
	}

	slice1Map := make(map[int]bool)
	for _, num := range slice1 {
		slice1Map[num] = true
	}
	comum := make([]int, 0)

	for _, num := range slice2 {
		if slice1Map[num] {
			comum = append(comum, num)
		}
	}

	return comum, nil
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 2, 6, 5}

	resultado, err := PontosComum(slice1, slice2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}
}
