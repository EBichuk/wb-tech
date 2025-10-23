package main

import (
	"flag"
	"fmt"
)

/*
Дана переменная типа int64. Разработать программу, которая устанавливает i-й бит этого числа в 1 или 0.

Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).

Подсказка: используйте битовые операции (|, &^).
*/

func ChangeBit(num int64, i int, isZeroBit bool) int64 {
	if isZeroBit {
		num &^= 1 << (i - 1)
	} else {
		num |= 1 << (i - 1)
	}

	return num
}

func main() {
	num := flag.Int64("n", 5, "int64 number")
	i := flag.Int("i", 1, "changing bit")
	isZeroBit := flag.Bool("zero", true, "change 1 to 0")

	flag.Parse()

	fmt.Println(ChangeBit(*num, *i, *isZeroBit))
}

/*
-n=6 -zero=false  		| 7  110    -> 111
-n=5              		| 4  101    -> 100
-n=46 -i=5 -zero=false	| 62 101110 -> 111110
*/
