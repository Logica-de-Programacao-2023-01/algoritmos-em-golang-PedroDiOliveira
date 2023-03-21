package main

import "fmt"

func main() {
	var x int
	fmt.Print("Qual o numero escolhido?")
	fmt.Scan(&x)
	r1 := x - 1
	r2 := x + 1
	fmt.Println("O antecessor de", x, "é", r1, ", ja o sucessor é ", r2)
}
