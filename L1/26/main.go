package main

import (
	"flag"
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке встречаются один раз (т.е. строка состоит из уникальных символов).

Вывод: true, если все символы уникальны, false, если есть повторения. Проверка должна быть регистронезависимой, т.е. символы в разных регистрах считать одинаковыми.

Например: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false.

Подумайте, какой структурой данных удобно воспользоваться для проверки условия.
*/

func isUniqueStr(str string) bool {
	mapUnique := make(map[rune]struct{})

	runes := []rune(strings.ToLower(str))

	for _, s := range runes {
		_, ok := mapUnique[s]
		if ok {
			return false
		}
		mapUnique[s] = struct{}{}
	}

	return true
}

func main() {
	str := flag.String("s", "VeryGoodDay", "checking string")

	flag.Parse()

	fmt.Println(isUniqueStr(*str))
}
