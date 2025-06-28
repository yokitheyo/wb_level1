package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	numWorkers := flag.Int("w", 3, "number of workers")
	flag.Parse()

	if *numWorkers <= 0 {
		fmt.Println("number of workers must be greater than 0")
		return
	}

	fmt.Printf("starting with %d workers....\n\n\n", *numWorkers)

	ctx, cancel := context.WithCancel(context.Background())
	dataChan := make(chan int)
	wg := sync.WaitGroup{}

	for i := 1; i <= *numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, dataChan, &wg)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(dataChan)
				return
			default:
				rndNum := rand.Intn(1000)
				dataChan <- rndNum
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)

	<-c

	fmt.Println("\ninterrupted signal")
	cancel()

	fmt.Println("waiting for workers to finish...\n\n")
	wg.Wait()
	fmt.Println("all workers finished")
}

func worker(ctx context.Context, id int, dataChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("worker %d started\n", id)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d stopped\n", id)
			return
		case data, ok := <-dataChan:
			if !ok {
				fmt.Printf("worker %d: channel closed\n", id)
				return
			}
			fmt.Printf("worker %d received data: %d\n", id, data)
		}
	}
}
