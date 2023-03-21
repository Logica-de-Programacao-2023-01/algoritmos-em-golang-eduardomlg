package main

import "fmt"

// Faça um algoritmo que leia um número
// inteiro e mostre se ele é par ou ímpar.
func main() {
	var x int64
	fmt.Print("Digite o número ")
	fmt.Scan(&x)
	if x%2 == 0 {
		fmt.Print("O número é par")
	} else {
		fmt.Print("O número é impar")
	}
}
