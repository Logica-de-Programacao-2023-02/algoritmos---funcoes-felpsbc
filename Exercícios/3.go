package main

import (
	"fmt"
	"strings"
)

func main() {
	x := []string{"felipe ", "joao ", "pedro"}
	resultado := concatenarStrings(x)
	fmt.Println(resultado)
}
func concatenarStrings(slice []string) string {
	return strings.Join(slice, "")
}
