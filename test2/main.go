package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var M map[string]string
var mu sync.Mutex

func Input(a int, b int) {
	M = make(map[string]string)
	for i := a; i < b; i++ {
		time.Sleep(1 * time.Millisecond)
		mu.Lock()
		t := strconv.Itoa(i)
		M[t] = t
		mu.Unlock()
	}
}

func Output() {
	fmt.Print(M)
}

func main() {
	Input(666, 1000)
	go Input(0, 333)
	go Input(333, 666)
	Output()
}
