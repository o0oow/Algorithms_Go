package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 7

	index := binarySearch(arr, target)
	if index != -1 {
		fmt.Printf("Элемент %d найден на индексе %d\n", target, index)
	} else {
		fmt.Println("Элемент не найден")
	}
}
