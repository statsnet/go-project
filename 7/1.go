package main

import "fmt"

func counter(arr []int) {
	fmt.Println(len(arr))
}

func main() {
	arr := []int{1,2,3,4,5}
	counter(arr)
}