package main

import "fmt"

func main() {
	var x int
	fmt.Println("Qual tabuada voce deseja exibir?")
	fmt.Scan(&x)
	for i := 1; i < 11; i++ {
		fmt.Println(i * x)
	}
}
