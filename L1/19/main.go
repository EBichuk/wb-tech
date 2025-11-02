package main

import "fmt"

/*
Разработать программу, которая переворачивает подаваемую на вход строку.

Например: при вводе строки «главрыба» вывод должен быть «абырвалг».

Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.),
то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).
*/

func turnString(inputStr string) string {
	str := []rune(inputStr)
	n := len(str)

	for i := 0; i < n/2; i++ {
		str[i], str[n-i-1] = str[n-i-1], str[i]
	}

	return string(str)
}

func main() {
	str := "главрыба"
	fmt.Println(turnString(str))
}
