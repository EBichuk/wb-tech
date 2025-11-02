package main

import (
	"strings"
)

/*
Рассмотреть следующий код и ответить на вопросы: к каким негативным последствиям он может привести и как это исправить?

var justString string

func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

1. Утечка памяти.
	Переменная justString это срез первых 100 байт большой строки v. justString ссылается на базовый массив байт v (1 Кб),
	поэтому сборщик мусора не сможет освободить память, хотя используется всего 100 байт

Вопрос: что происходит с переменной justString?
	Изначально: justString это строка длиной 100 байт, которая создана на основе слайса, ссылающегося на строку v длиной 1024 байтов
	После: 		justString это строка длиной 100 байт, которая создана на основе слайса длиной 100 байт куда скопированы первые 100 байт
				изначальной скроки
*/

var justString string

func createHugeString(count int) string {
	return strings.Repeat("w", count)
}

func someFunc() {
	v := createHugeString(1 << 10)

	temp := make([]byte, 100)
	copy(temp, v[:100])
	justString = string(temp)
}

func main() {
	someFunc()
}
