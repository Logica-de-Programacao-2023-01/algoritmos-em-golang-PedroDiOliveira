package main

import "fmt"

func main() {
	var salario float64
	fmt.Print("Qual o valor do salario?")
	fmt.Scan(&salario)
	aumento := salario * 0.15
	novoSalario := aumento + salario
	fmt.Print("Com o aumento, o salario passou a ser:", novoSalario)
}
