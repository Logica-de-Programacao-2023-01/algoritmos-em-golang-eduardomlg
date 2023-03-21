package main

import (
	"fmt"
)

// Faça um algoritmo que leia
// três números inteiros e mostre o menor deles.
func main() {
	var num1 int64
	var num2 int64
	var num3 int64
	fmt.Print("Quais são os números ")
	fmt.Scan(&num1, &num2, &num3)
	menor := num1
	if num2 < menor {
		menor = num2
	}
	if num3 < menor {
		menor = num3
	}
	fmt.Println("O menor número é ", menor)
}
