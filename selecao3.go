package main

import "fmt"

func main() {
	var numero int
	fmt.Println("Qual é o numero?")
	fmt.Scan(&numero)
	if numero%2 == 0 {
		fmt.Println("O numero é par!")
	} else {
		fmt.Println("O numero é impar!")
	}
}
