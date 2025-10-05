package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

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

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(numWorkers)

	for w := 1; w <= numWorkers; w++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("worker", w, "stoppped job")
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
