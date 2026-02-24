package main

import "fmt"

func main() {
	fmt.Println("O resultado da soma é:", Soma(5, 10))
}

func Soma(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mult(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}