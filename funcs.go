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

// Алгоритм Решето Эратосфена для нахождения простых чисел
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
