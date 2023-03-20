package main

import "fmt"

func main() {
	// Faça um algoritmo que leia três números inteiros e mostre a soma entre eles.
	var n1 int
	var n2 int
	var n3 int
	fmt.Print("Primeiro numero:")
	fmt.Scan(&n1)
	fmt.Print("Segundo numero:")
	fmt.Scan(&n2)
	fmt.Print("Terceiro numero:")
	fmt.Scan(&n3)
	fmt.Println("A soma dos numeros inteiros é:", n1+n2+n3)
}
