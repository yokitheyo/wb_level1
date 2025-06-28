package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	duration := flag.Int("t", 10, "duration in seconds")
	flag.Parse()

	if *duration <= 0 {
		fmt.Println("duration must be greater than 0")
		return
	}
	stop := time.After(time.Duration(*duration) * time.Second)

	fmt.Printf("countdown for %d seconds\n\n", *duration)

	dataCh := make(chan int)
	wg := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-stop:
				fmt.Println("time is up!")
				close(dataCh)
				return
			case dataCh <- i:
				fmt.Printf("sended: %d\n", i)
				i++
				time.Sleep(700 * time.Millisecond)
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range dataCh {
			fmt.Printf("received: %d\n", data)
		}
	}()

	wg.Wait()
	fmt.Println("program finished")
}
