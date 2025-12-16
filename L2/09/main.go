package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
Написать функцию Go, осуществляющую примитивную распаковку строки, содержащей повторяющиеся символы/руны.
*/

var ErrInvalidUnpackStr = errors.New("invalid unpack string")

func unpackingString(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	strRune := []rune(str)

	// string starting with a number and string consisting of numbers are incorrect
	if strRune[0] > '0' && strRune[0] <= '9' {
		return "", ErrInvalidUnpackStr
	}

	var builder strings.Builder
	prev := '0'
	count := 0
	isEscape := false

	for _, s := range strRune {
		if s == '\\' {
			isEscape = true
			continue
		}

		if s >= '0' && s <= '9' && !isEscape {
			count = count*10 + int(s-'0')
		} else {
			if count > 0 {
				builder.WriteString(strings.Repeat(string(prev), count-1))
				count = 0
			}

			// write it down once and remember rune
			builder.WriteRune(s)
			prev = s
			isEscape = false
		}
	}

	if isEscape {
		return "", ErrInvalidUnpackStr
	}

	// unpacking the last symbol
	if count > 1 {
		builder.WriteString(strings.Repeat(string(prev), count-1))
	}

	return builder.String(), nil
}

func main() {
	str, err := unpackingString("a10f")
	fmt.Println(str, err)
}
