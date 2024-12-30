package main

import "fmt"

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10, 1, 13}
	fmt.Println("Исходный массив:", arr)
	mrgsrt := mergeSort(arr)
	qusort := quickSort(arr)
	fmt.Println("Отсортированный массив MS:", mrgsrt)
	fmt.Println("Отсортированный массив QS:", qusort)
	
}
