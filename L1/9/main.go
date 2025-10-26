package main

import (
	"fmt"
	"time"
)

/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа x из массива, во второй – результат операции x*2.
После этого данные из второго канала должны выводиться в stdout.
То есть, организуйте конвейер из двух этапов с горутинами: генерация чисел и их обработка.
Убедитесь, что чтение из второго канала корректно завершается.
*/

func writer(inputLst []int) <-chan int {
	ch := make(chan int)

	go func() {
		for _, v := range inputLst {
			ch <- v
			time.Sleep(time.Second)
		}

		close(ch)
	}()

	return ch
}

func doubler(inputCh <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for i := range inputCh {
			out <- i * 2
		}
		close(out)
	}()

	return out
}

func reader(inputCh <-chan int) {
	for i := range inputCh {
		fmt.Println(i)
	}
}

func main() {
	lst := make([]int, 0)
	for i := 10; i < 20; i++ {
		lst = append(lst, i)
	}

	reader(doubler(writer(lst)))
}
