package main

import "fmt"

func main() {
	//Faça um algoritmo que leia o número de dias trabalhados por um funcionário
	//e o valor da sua diária e calcule o seu salário.
	var dias int
	var diaria int
	fmt.Print("Quantos dias voce trabalha?")
	fmt.Scan(&dias)
	fmt.Println("Qual o valor da sua diaria?")
	fmt.Scan(&diaria)
	resultado := dias * diaria
	fmt.Println("O seu salario e igual a:", resultado)
}
