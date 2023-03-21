package main

import "fmt"

// Faça um algoritmo que leia dois números inteiros e mostre o maior deles.
func main() {
	var num1 int64
	var num2 int64
	fmt.Print("Qual o primeiro número? ")
	fmt.Scan(&num1)
	fmt.Print("Qual o segundo número? ")
	fmt.Scan(&num2)
	if num1 > num2 {
		fmt.Println(num1, " é maior que ", num2)
	} else if num1 < num2 {
		fmt.Println(num2, " é maior que ", num1)
	} else {
		fmt.Println(num1, " e ", num2, " são números iguais")
	}
}
