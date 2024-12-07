package main

// Бинарный поиск
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2 // Избегаем переполнения при больших значениях
		if arr[mid] == target {
			return mid // Элемент найден
		} else if arr[mid] < target {
			left = mid + 1 // Сужаем поиск к правой половине
		} else {
			right = mid - 1 // Сужаем поиск к левой половине
		}
	}

	return -1 // Элемент не найден
}
