package main

import "fmt"

func main() {
	var x int
	fmt.Println("Escreva um numero?")
	fmt.Scan(&x)
	for i := 1; i < x+1; i++ {
		if x%i == 0 {
			fmt.Println(i)
		}
	}
}
