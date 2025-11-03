package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счётчик, которая будет инкрементироваться в конкурентной среде (т.е. из нескольких горутин).
По завершению программы структура должна выводить итоговое значение счётчика.

Подсказка: вам понадобится механизм синхронизации, например, sync.Mutex или sync/Atomic для безопасного инкремента.
*/
type Counter struct {
	value int
	mu    *sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{
		value: 0,
		mu:    &sync.Mutex{},
	}
}

func (c *Counter) Add() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func main() {
	c := NewCounter()
	wg := &sync.WaitGroup{}

	wg.Add(1000)
	for range 1000 {
		go func() {
			defer wg.Done()
			c.Add()
		}()
	}

	wg.Wait()
	fmt.Println(c.value)
}
