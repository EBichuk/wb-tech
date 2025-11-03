package main

import "fmt"

/*
Реализовать алгоритм быстрой сортировки массива встроенными средствами языка.
Можно использовать рекурсию.

Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел.
Для выбора опорного элемента можно взять середину или первый элемент.
*/

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[0]
	left := 0

	nums[left], nums[len(nums)-1] = nums[len(nums)-1], nums[left]

	for i, v := range nums {
		if v < pivot {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}

	nums[left], nums[len(nums)-1] = nums[len(nums)-1], nums[left]

	quickSort(nums[:left])
	quickSort(nums[left+1:])

	return nums
}

func main() {
	fmt.Println(quickSort([]int{10, 10, 10, 66, 7, 3, 2, 20, 12, 33, 10}))
}
