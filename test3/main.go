package main

import (
	"log"
	"sync"
)

var mu sync.Mutex

//1000 goroutines truy cap chung vao m(map)

func errFunc() {
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 1; j < 10000; j++ {
				mu.Lock()
				if _, ok := m[j]; ok {
					delete(m, j)
					continue
				}
				m[j] = j * 10
				mu.Unlock()
			}
		}()
	}

	log.Print("done")
}

func main() {
	errFunc()
}
