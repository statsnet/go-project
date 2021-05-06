package main

import "fmt"
// СОРТИРОВКА ПО ВОЗРАСТАНИЮ
// https://taskcode.ru/matrix/sorting
func main() {
	arr := [][]int{
		{3, -2, 6, 4},
		{8, 1, 12, 2},
		{5, 4, -8, 0},
	}
	for _, lst := range arr {
		for i, _ := range lst {
			lowest_value_index := i
			for a := 0; a < len(lst); a++ {
				if lst[a] > lst[lowest_value_index] {
					lowest_value_index = a
				}
				lst[i], lst[lowest_value_index] = lst[lowest_value_index], lst[i]
			}
		}
		fmt.Println(lst)
	}
}
