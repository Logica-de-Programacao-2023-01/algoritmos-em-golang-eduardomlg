package main

import "fmt"

// Faça um algoritmo que leia o preço de um produto e mostre o seu valor com desconto de 10%
func main() {
	var preco float64
	var resultado float64
	fmt.Print("Qual o preço? ")
	fmt.Scan(&preco)
	resultado = preco * 0.9
	fmt.Print(resultado)
}
