package main

import "fmt"

func main() {
	var even_number int
	var odd_number int
	for i := 3; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			even_number += i
		} else {
			odd_number += i
		}
	}
	fmt.Println(even_number)
	fmt.Println(odd_number)
}