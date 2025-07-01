package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	done := make(chan bool)

	go func() {
		<-time.After(duration)

		done <- true
	}()
	<-done
}

func main() {
	start := time.Now()
	fmt.Printf("sleep for 3 seconds... (time: %v)\n", start.Format(time.RFC1123))
	sleep(3 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("sleep finished (time: %v, elapsed: %v)\n", time.Now().Format(time.RFC1123), elapsed)
}
