package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов) —
т.е. вывести элементы, присутствующие и в первом, и во втором.

Пример:
A = {1,2,3}
B = {2,3,4}
Пересечение = {2,3}
*/
func intersectionSets(s1, s2 []int) []int {
	hashMap := make(map[int]struct{})
	intersection := make([]int, 0)

	for _, elem := range s1 {
		hashMap[elem] = struct{}{}
	}

	for _, elem := range s2 {
		if _, ok := hashMap[elem]; ok {
			intersection = append(intersection, elem)
		}
	}

	return intersection
}

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{2, 3, 4}

	fmt.Println(intersectionSets(s1, s2))
}
