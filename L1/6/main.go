package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var maxWaitSeconds = 3

func randomWork() {
	time.Sleep(time.Duration(rand.Intn(maxWaitSeconds)+1) * time.Second)
}

func stopByCondition() {
	wg := &sync.WaitGroup{}

	stopCondition := false

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			if stopCondition {
				fmt.Println("goroutine is stopping")
				break
			}
			randomWork()
			fmt.Println("goroutine do some work")
		}
	}()

	time.Sleep(time.Second * 3)
	stopCondition = true

	wg.Wait()
	fmt.Println("goroutine stopped by condition")
}

func stopbyChannel() {
	wg := &sync.WaitGroup{}
	stopCh := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stopCh:
				fmt.Println("goroutine is stopping")
				return
			default:
				randomWork()
				fmt.Println("goroutine do some work")
			}
		}
	}()

	time.Sleep(time.Second * 3)

	stopCh <- struct{}{}

	wg.Wait()
	fmt.Println("goroutine stopped by channel")
}

func stopByContext() {
	wg := &sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine is stopping")
				return
			default:
				randomWork()
				fmt.Println("goroutine do some work")
			}
		}
	}()

	time.Sleep(time.Second * 3)

	cancel()

	wg.Wait()
	fmt.Println("goroutine stopped by context cancellation")
}

func stopByGoExit() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			randomWork()
			fmt.Println("goroutine do some work")
			runtime.Goexit()
		}
	}()

	wg.Wait()
	fmt.Println("goroutine stopped by runtime.Goexit")
}

func stopByTimeout() {
	wg := &sync.WaitGroup{}

	timeout := time.After(time.Second * 4)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-timeout:
				fmt.Println("goroutine is stopping")
				return
			default:
				randomWork()
				fmt.Println("goroutine do some work")
			}
		}
	}()

	wg.Wait()
	fmt.Println("goroutine stopped by timeout")
}

func main() {
	fmt.Println("-- start: stop by condition --")
	stopByCondition()
	fmt.Println("-- end: stop by condition --")

	fmt.Println("\n-- start: stop by channel ---")
	stopbyChannel()
	fmt.Println("-- end: stop by channel ---")

	fmt.Println("\n-- start: stop by context --")
	stopByContext()
	fmt.Println("-- end: stop by context --")

	fmt.Println("\n-- start: stop by runtime.Goexit --")
	stopByGoExit()
	fmt.Println("-- end: stop by runtime.Goexit --")

	fmt.Println("\n-- start: stop by timeout --")
	stopByTimeout()
	fmt.Println("-- end: stop by timeout --")
}
