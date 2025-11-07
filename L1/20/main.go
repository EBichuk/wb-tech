package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает порядок слов в строке.

Пример: входная строка:

«snow dog sun», выход: «sun dog snow».

Считайте, что слова разделяются одиночным пробелом.
Постарайтесь не использовать дополнительные срезы, а выполнять операцию «на месте».
*/

func changeOrderWordsInPlace(str string) string {
	runes := []rune(str)

	reverseRunes(runes, 0, len(runes)-1)

	start := 0
	for i := 0; i < len(runes); i++ {
		if runes[i] == ' ' {
			reverseRunes(runes, start, i-1)
			start = i + 1
		}
	}

	reverseRunes(runes, start, len(runes)-1)

	return string(runes)
}

func reverseRunes(runes []rune, start, end int) {
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
}

func changeOrderWords(str string) string {
	words := strings.Fields(str)

	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}

	resStr := strings.Join(words, " ")

	return resStr
}

func main() {
	inputStr := "snow dog sun"

	fmt.Println(changeOrderWords(inputStr))
	fmt.Println(changeOrderWordsInPlace(inputStr))
}
