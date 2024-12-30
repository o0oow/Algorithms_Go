package main

import "math"

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

// Функция для слияния двух отсортированных массивов
func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	// Сливаем элементы, сравнивая их
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Добавляем оставшиеся элементы (если есть)
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// Рекурсивная функция для сортировки
func mergeSort(arr []int) []int {
	// Базовый случай: массив из одного элемента
	if len(arr) <= 1 {
		return arr
	}

	// Делим массив на две части
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// Сливаем отсортированные половины
	return merge(left, right)
}

// Быстрая сортировка
func quickSort(arr []int) []int {
	// Базовый случай: массив длиной 0 или 1 уже отсортирован
	if len(arr) <= 1 {
		return arr
	}

	// Опорный элемент
	pivot := arr[len(arr)/2]

	// Три части: меньше, равны, больше
	less := []int{}
	equal := []int{}
	greater := []int{}

	for _, num := range arr {
		if num < pivot {
			less = append(less, num)
		} else if num == pivot {
			equal = append(equal, num)
		} else {
			greater = append(greater, num)
		}
	}

	// Рекурсивно сортируем части и объединяем
	return append(append(quickSort(less), equal...), quickSort(greater)...)
}

// Алгоритм Решето Эратосфена для нахождения простых чисел ДО ЧИСЛА N
func sieveOfEratosthenes(n int) []int {
	is_Prime := make([]bool, n+1) // Булевый массив для проверки
	for i := 2; i <= n; i++ {
		is_Prime[i] = true // Все числа изначально считаются простыми
	}

	for i := 2; i*i <= n; i++ {
		if is_Prime[i] {
			for j := i * i; j <= n; j += i { // Убираем кратные i
				is_Prime[j] = false
			}
		}
	}

	// Собираем все оставшиеся простые числа
	primes := []int{}
	for i := 2; i <= n; i++ {
		if is_Prime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

// Проверка, является ли число простым
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false // Найден делитель, значит число не простое
		}
	}
	return true
}
