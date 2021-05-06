package main

import (
	"fmt"
)

// Найти максимальные элементы столбцов матрицы
func main() {
	arr := [][]int{{5,3,4,1,2}, {7,6,10,8,9}}
	max_arr := []int{}

	for _, v := range arr {
		max := v[0]
		for _, value := range v {
			if value > max {
				max = value
			}
		}
		max_arr = append(max_arr, max)
	}
	fmt.Println(max_arr)
}
