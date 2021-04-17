package main

import "fmt"

func main() {
	// Стороны треугольника
	var a, b, c int
	a, b, c = 15, 20, 30
	if a == b && b == c {
		fmt.Print("Равносторонний треугольник")
	} else if a == b && b != c {
		fmt.Println("Равнобедренный треугольник")
	} else {
		fmt.Println("Разносторонник треугольник")
	}
}