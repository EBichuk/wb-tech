package main

import (
	"flag"
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a, b,
значения которых > 2^20 (больше 1 миллион).

Комментарий: в Go тип int справится с такими числами, но обратите внимание на возможное переполнение для ещё больших значений.
Для очень больших чисел можно использовать math/big.
*/

// максимальное int = 9223372036854775807
func printBigIntOperations(aString, bString string) {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString(aString, 10)
	b.SetString(bString, 10)

	sum := new(big.Int).Add(a, b)
	sub := new(big.Int).Sub(a, b)
	mul := new(big.Int).Mul(a, b)
	div := new(big.Int).Div(a, b)

	fmt.Printf("Произведение \t%v\n", mul)
	fmt.Printf("Частное \t%v\n", div)
	fmt.Printf("Сумма   \t%v\n", sum)
	fmt.Printf("Разность \t%v\n", sub)
}

func main() {
	a := flag.String("a", "100000000", "1 number")
	b := flag.String("b", "123456789", "2 number")

	flag.Parse()

	printBigIntOperations(*a, *b)
}
