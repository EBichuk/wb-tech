package main

import (
	"fmt"
	"time"
)

func OrSelect(channels ...<-chan interface{}) <-chan interface{} {
	done := make(chan interface{})

	go func() {
		for {
			for _, channel := range channels {
				select {
				case <-channel:
					close(done)
					return
				default:
				}
			}
		}
	}()

	return done
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-OrSelect(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v\n", time.Since(start))
}
