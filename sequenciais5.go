package main

import "fmt"

func main() {
	var idade int
	fmt.Print("Qual a sua idade em anos?")
	fmt.Scan(&idade)
	fmt.Print("A sua idade em dias é:", idade*365)
}
