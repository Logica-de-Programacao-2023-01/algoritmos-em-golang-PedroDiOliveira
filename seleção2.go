package main

import "fmt"

func main() {
	var n1 int
	var n2 int
	var n3 int
	fmt.Print("Qual é o primeiro numero?")
	fmt.Scan(&n1)
	fmt.Print("Qual é o segundo numero?")
	fmt.Scan(&n2)
	fmt.Print("Qual é o terceiro numero?")
	fmt.Scan(&n3)
	if n1 < n2 && n1 < n3 {
		fmt.Println("O menor numero é:", n1)
	} else if n2 < n1 && n2 < n3 {
		fmt.Println("O menor numero é:", n2)
	} else {
		fmt.Println("O menor numero é:", n3)
	}
}
