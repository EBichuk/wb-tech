package main

import (
	"fmt"
	"sync"
)

/*
Реализовать безопасную для конкуренции запись данных в структуру map.

Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).

Проверьте работу кода на гонки (util go run -race).
*/

var numWorkers = 5
var numberOfRecords = 10

type SyncMap struct {
	data map[int]any
	mu   *sync.RWMutex
}

func NewSyncMap() *SyncMap {
	return &SyncMap{
		data: make(map[int]any),
		mu:   &sync.RWMutex{},
	}
}

func (sm *SyncMap) Set(key int, value any) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.data[key] = value
}

func (sm *SyncMap) Get(key int) (any, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	value, found := sm.data[key]
	return value, found
}

func main() {
	wg := &sync.WaitGroup{}

	m := NewSyncMap()

	wg.Add(numWorkers)
	for i := range numWorkers {
		go func() {
			defer wg.Done()
			for j := range numberOfRecords {
				m.Set(numberOfRecords*i+j, i)
			}
		}()
	}

	wg.Wait()

	for i := range numWorkers*numberOfRecords - 1 {
		val, ok := m.Get(i)
		if ok {
			fmt.Println(val)
		}
	}
}
