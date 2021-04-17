package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var start, end string
	fmt.Print("Введите цифру start:")
	fmt.Scan(&start)
	fmt.Print("Введите цифру end:")
	fmt.Scan(&end)

	digit1, err1 := strconv.Atoi(start)
	if err1 != nil {
		log.Fatal(err1)
	}
	digit2, err2 := strconv.Atoi(end)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("Результат сложения: ",digit1 + digit2)

}
