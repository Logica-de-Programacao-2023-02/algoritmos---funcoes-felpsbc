package main

import "fmt"

func main() {

	var x int = 0
	fmt.Println("Escreva um número para saber se ele é primo")
	fmt.Scan(&x)
	if x <= 0 {
		fmt.Println("O número não pode ser negativo(ERROR)")
	}
	resultado := isPrime(x)
	fmt.Println(resultado)

}
func isPrime(x int) bool {
	if x == 2 {
		return true
	} else if x%2 == 0 {
		return false
	} else {
		return true
	}
}
