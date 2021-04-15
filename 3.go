package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Введите число А")
	fmt.Scan(&a)
	fmt.Println("Введите число Б")
	fmt.Scan(&b)
	fmt.Println("Введите число В")
	fmt.Scan(&c)

	var max, min, avg int
	if a > b && a > c {
		max = a
		if b < c {
			min = b
			avg = c
		} else {
			min = c
			avg = b
		}
	}
	if b > a && b > c {
		max = b
		if a < c {
			min = a
			avg = c
		} else {
			min = c
			avg = c
		}
	}
	if c > a && c > b {
		max = c
		if a < b {
			min = a
			avg = b
		} else {
			min = b
			avg = a
		}
	}
	result := fmt.Sprintf("Минимальным числом является %d, средним %d, и самым крупным %d", min, avg, max)
	fmt.Println(result)
}
