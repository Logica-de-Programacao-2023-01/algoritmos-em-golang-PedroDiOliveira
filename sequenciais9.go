package main

import "fmt"

func main() {
	//Faça um algoritmo que leia o preço de um produto e mostre o seu valor com desconto de 10%.
	var x float64
	fmt.Print("Qual é o preço do produto? ")
	fmt.Scan(&x)
	desconto := x * 0.1
	preçoFinal := x - desconto
	fmt.Print("O preço final do produto fica na casa dos:", preçoFinal)
}
