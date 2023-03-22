package main

import "fmt"

func main() {
	var peso int
	var altura int
	fmt.Print("Quanto voce pesa?")
	fmt.Scan(&peso)
	fmt.Print("Qual a sua altura?")
	fmt.Scan(&altura)
	alt := altura * altura
	fmt.Println("A seu Indice de massa corporea é:", peso/alt)
	//IMCOMPLETO
}
