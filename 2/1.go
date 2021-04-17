package main

import "fmt"

func main()  {
	var year int
	fmt.Print("Введите год:")
	fmt.Scan(&year)
	if year%4 != 0 {
		fmt.Println("Обычный год")
	} else {
		fmt.Println("Високосный год")
		if year%100 == 0 {
			fmt.Println("Обычный год")
		} else {
			fmt.Println("Високосный год")
		}
	}
}
