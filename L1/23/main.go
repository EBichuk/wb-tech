package main

import "fmt"

/*
Удалить i-ый элемент из слайса. Продемонстрируйте корректное удаление без утечки памяти.

Подсказка: можно сдвинуть хвост слайса на место удаляемого элемента (copy(slice[i:], slice[i+1:])) и уменьшить длину слайса на 1.
*/
func delItemFromSlice(sl []int, i int) []int {
	copy(sl[i:], sl[i+1:])
	sl = sl[:len(sl)-1]

	return sl
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sl, len(sl), cap(sl))

	nums := delItemFromSlice(sl, 4)
	fmt.Println(nums, len(nums), cap(nums))
}
