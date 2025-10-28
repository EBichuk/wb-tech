package main

import "fmt"

/*
Реализовать алгоритм бинарного поиска встроенными методами языка.
Функция должна принимать отсортированный слайс и искомый элемент, возвращать индекс элемента или -1, если элемент не найден.

Подсказка: можно реализовать рекурсивно или итеративно, используя цикл for.
*/

func binarySearch(lst []int, item int) int {
	left, right := 0, len(lst)-1
	for left <= right {
		mid := left + (right-left)/2
		guess := lst[mid]

		if guess == item {
			return mid
		}

		if guess > item {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	nums := []int{1, 3, 6, 6, 7, 8, 9, 11, 18}

	fmt.Println(binarySearch(nums, 11))
}
