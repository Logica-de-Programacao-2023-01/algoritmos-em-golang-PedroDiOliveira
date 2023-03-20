package main

import (
	"fmt"
)

func main() {
	var n1 int
	var n2 int
	fmt.Print("Primeiro numero:")
	fmt.Scan(&n1)
	fmt.Print("Segundo numero:")
	fmt.Scan(&n2)
	if n1 > n2 {
		fmt.Println("O seu numero maior é:", n1)
	} else {
		fmt.Println("O seu numero maior é:", n2)
	}
}
