package main

import "fmt"

/*
Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.

Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.
*/
func uniqueStrings(inputStrings []string) []string {
	uniqueStr := make([]string, 0)

	hashMap := make(map[string]struct{})

	for _, str := range inputStrings {
		if _, ok := hashMap[str]; !ok {
			hashMap[str] = struct{}{}

			uniqueStr = append(uniqueStr, str)
		}
	}

	return uniqueStr
}

func main() {
	inputStrings := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(uniqueStrings(inputStrings))
}
