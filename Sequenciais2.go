package main

import "fmt"

func main() {
	//Faça um algoritmo que leia um número inteiro e mostre o seu dobro, triplo e quadruplo.
	var n1 int
	fmt.Print("Numero:")
	fmt.Scan(&n1)
	fmt.Print("    Dobro:", n1*2, "    Triplo:", n1*3, "    Quadruplo:", n1*4)
}
