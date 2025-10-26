package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Поменять местами два числа без использования временной переменной.

Подсказка: примените сложение/вычитание или XOR-обмен.
*/
func main() {
	if len(os.Args) < 2 {
		fmt.Println("enter two numbers")
		return
	}

	n1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("enter number: ", err)
		return
	}
	n2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("enter number: ", err)
		return
	}

	n1 = n2 + n1 // a = a + b
	n2 = n1 - n2 // b = (a + b) - b = a
	n1 = n1 - n2 // a = (a + b) - a = b

	fmt.Println("first number =", n1, "second number =", n2)
}

/*
using XOR

func main() {
	if len(os.Args) < 2 {
		fmt.Println("enter two numbers")
		return
	}

	num1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	num2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return
	}

	n1 ^= n2
	n2 ^= n1
	n1 ^= n2

	fmt.Println("first number =", n1, "second number =", n2)
}

*/
