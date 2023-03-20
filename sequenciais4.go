package main

import "fmt"

func main() {
	var n1 int
	var n2 int
	var n3 int
	p1 := 2
	p2 := 3
	p3 := 5
	somaPesos := p1 + p2 + p3
	fmt.Print("Qual é o primeiro numero?")
	fmt.Scan(&n1)
	fmt.Print("Qual é o segundo numero?")
	fmt.Scan(&n2)
	fmt.Print("Qual é o terceiro numero?")
	fmt.Scan(&n3)
	m1 := n1 * p1
	m2 := n2 * p2
	m3 := n3 * p3
	somaValores := m1 + m2 + m3
	resultado := somaValores / somaPesos
	fmt.Println("A media ponderada é:", resultado)
}
