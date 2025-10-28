package main

import "fmt"

/*
Разработать программу, которая в runtime способна определить тип переменной, переданной в неё
(на вход подаётся interface{}). Типы, которые нужно распознавать: int, string, bool, chan (канал).

Подсказка: оператор типа switch v.(type) поможет в решении.
*/

func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println(i, "is integer")
	case string:
		fmt.Println(i, "is string")
	case bool:
		fmt.Println(i, "is bool")
	case chan any:
		fmt.Println(i, "is chan any")
	default:
		fmt.Println(i, "is unknown type")
	}
}

func main() {
	printType(5)
	printType("str")
	printType(make(chan any))
	printType(true)
	printType(1.23)
}
