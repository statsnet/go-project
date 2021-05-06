package main

import (
	"fmt"
)

// Вывести на экран индексы всех минимальных элементов матрицы.
func main() {
	arr := [][]int{{5,3,4,1,2}, {7,6,10,8,9}}
	minimal_indexes := []int{}

	for _, v := range arr {
		min := v[0]
		var mini int
		for i, value := range v {
			if value < min {
				min = value
				mini = i
				fmt.Println(value, i)
			}
		}
		minimal_indexes = append(minimal_indexes, mini)
	}
	fmt.Println(minimal_indexes)
}
