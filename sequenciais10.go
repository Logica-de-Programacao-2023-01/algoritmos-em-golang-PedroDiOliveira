package main

import "fmt"

func main() {
	var peso int
	var conta = 2.20462
	fmt.Print("Quanto voce pesa?")
	fmt.Scan(&peso)
	libras := int(float64(peso) * conta)
	fmt.Print("Seu peso em libras é:", libras)
}
