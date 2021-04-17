package main

import (
	"fmt"
	"strconv"
)

func main() {
	var start, end string
	fmt.Scan(&start)
	fmt.Scan(&end)
	start_digit, _ := strconv.Atoi(start)
	end_digit, _ := strconv.Atoi(end)

	var evenNumber, oddNumber, total int
	for i := start_digit; i < end_digit; i++ {
		total += i
		if i%2 == 0 {
			fmt.Println(i)
			evenNumber += i
		} else {
			oddNumber += i
		}
	}
	fmt.Println("Общая сумма всех цифр", total)
	fmt.Println("Сумма четных чисел", evenNumber)
	fmt.Println("Сумма нечетных чисел", oddNumber)
}