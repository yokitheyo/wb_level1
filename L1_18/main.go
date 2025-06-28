package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	val int64
}

func (c *Counter) Inc() {
	atomic.AddInt64(&c.val, 1)
}

func (c *Counter) Val() int64 {
	return atomic.LoadInt64(&c.val)
}

func main() {
	wg := sync.WaitGroup{}

	counter := &Counter{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Inc()
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter.Val())
}
