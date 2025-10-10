package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	jobsCh := make(chan int, 1)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timeout := time.After(time.Second * 5)

	wg := &sync.WaitGroup{}
	wg.Add(1)

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
				fmt.Println(v)
				time.Sleep(time.Second)
			}
		}
	}()

l:
	for {
		select {
		case <-timeout:
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
