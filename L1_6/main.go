package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func stopByCh() {
	fmt.Println("\n === Stop with channel === \n")

	stopCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-stopCh:
				fmt.Println("stopped with ch")
				return
			default:
				fmt.Println("working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	close(stopCh)
	time.Sleep(100 * time.Millisecond)
}

func stopByContext() {
	fmt.Println("\n === Stop with context === \n")

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stopped with context")
				return
			default:
				fmt.Println("working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(100 * time.Millisecond)
}

func stopByContextWithTimeout() {
	fmt.Println("\n === Stop with context with timeout === \n")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stopped with context with timeout")
				return
			default:
				fmt.Println("working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(3 * time.Second)
}

func stopByGoexit() {
	fmt.Println("\n === Stop with goexit === \n")

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("stopped with goexit")

		for i := 0; i < 5; i++ {
			if i == 3 {
				fmt.Println("stopping with goexit")
				runtime.Goexit()
			}
			fmt.Println("working...")
		}
	}()
	wg.Wait()
}
func stopByPanic() {
	fmt.Println("\n=== stop by panic ===")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("goroutine stopped by panic: %v\n", r)
			}
		}()

		for i := 0; i < 5; i++ {
			if i == 3 {
				panic("panic!")
			}
			fmt.Println("working... ")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	wg.Wait()
}

func main() {
	stopByCh()
	stopByContext()
	stopByContextWithTimeout()
	stopByGoexit()
	stopByPanic()
	fmt.Println("\nmain finished")
}
