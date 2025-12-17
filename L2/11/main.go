package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

/*
Напишите функцию, которая находит все множества анаграмм по заданному словарю.
*/

func findAllAnagrams(words []string) map[string][]string {

	anagramMap := make(map[string][]string)
	resultMap := make(map[string][]string)

	for _, word := range words {
		word := strings.ToLower(word)

		r := []rune(word)
		slices.Sort(r)
		sortedWord := string(r)

		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	for sortedWord := range anagramMap {
		if sl := anagramMap[sortedWord]; len(sl) > 1 {
			firstWord := sl[0]
			sort.Strings(sl)
			resultMap[firstWord] = sl
		}
	}

	return resultMap
}

func main() {
	input := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	fmt.Println(findAllAnagrams(input))
}
