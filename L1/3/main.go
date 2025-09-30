package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

/*
Реализовать постоянную запись данных в канал (в главной горутине).

Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.

Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.
*/
func main() {
	if len(os.Args) < 2 {
		fmt.Println("enter the number of workers")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	if numWorkers <= 0 {
		fmt.Println("the number of workers must be positive")
		return
	}

	jobsCh := make(chan int, numWorkers)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(numWorkers)

	for w := 1; w <= numWorkers; w++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-jobsCh:
					if !ok {
						return
					}
					fmt.Println("worker", w, "started job", v)
					time.Sleep(time.Second)
					fmt.Println("worker", w, "finished job", v)
				}
			}
		}()
	}

l:
	for {
		select {
		case <-ctx.Done():
			break l
		default:
		}
		jobsCh <- getData()
	}

	close(jobsCh)
	wg.Wait()
}

func getData() int {
	return rand.Intn(100)
}
