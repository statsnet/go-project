package main

import (
	"fmt"
	"math/rand"
)

// https://taskcode.ru/matrix/qty
func main() {
	arr_count := 10
	int_count := 20
	repeated := []int{}

	arr := [][]int{}
	for i := 0; i < arr_count; i++ {
		fmt.Println(i)
		lst := []int{}
		counter := 0
		for j := 0; j < int_count; j++ {
			a := rand.Intn(15)
			if a == 5 {
				counter += 1
			}
			lst = append(lst, a)
		}
		if counter > 3 {
			repeated = append(repeated, i)
		}
		arr = append(arr, lst)
	}
	fmt.Println("Исходная матрица: ", arr)
	fmt.Println("Номера массивов где число 5 повторяется 3 и более раз: ",repeated)
}
