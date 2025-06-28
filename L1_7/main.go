package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.RWMutex
	m  map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]int),
	}
}

func (s *SafeMap) Write(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func (s *SafeMap) Read(key string) (int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.m[key]
	return v, ok
}

func main() {
	safeMap := NewSafeMap()
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", i%10)
			safeMap.Write(key, i)
		}()
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key-%d", i)
		if val, ok := safeMap.Read(key); ok {
			fmt.Printf("%s: %d\n", key, val)
		}
	}
}
