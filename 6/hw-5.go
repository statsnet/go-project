package main

import (
	"fmt"
	"math/rand"
)

func main() {
	m := 10
	n := 5
	arr := [][]int{}
	for i:=0;i<m;i++ {
		lst := []int{}
		for j:=0;j<n;j++{
			lst = append(lst, rand.Intn(200))
		}
		arr = append(arr, lst)
	}
	maxi := -1

	for i := 0; i < m; i ++ {
		mini := 200
		for j := 0; j < n; j++ {
			if arr[i][j] < mini {
				mini = arr[i][j]
			}
		}
		if mini > maxi {
			maxi = mini
		}
	}
	fmt.Println("Максимальный среди минимальных: ", maxi)
}
