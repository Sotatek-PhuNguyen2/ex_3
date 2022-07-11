package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var M map[string]string
var mu sync.Mutex

func chanRoutine() {
	log.Print("hello 1")
	pipe := make(chan string, 10)
	go func() {
		time.Sleep(1 * time.Second)
		pipe <- "hello 3"
		receiver := <-pipe
		log.Print(receiver)
	}()
	log.Print("hello 2")
}

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

func Test4() {
	var file, err = os.Open("file.txt")
	if err != nil {
		fmt.Print("Error opening file.txt: ", err)
	}
	defer file.Close()

	c := make(chan string, 10)
	index := 0

	scanner := bufio.NewScanner(file)
	mu.Lock()
	for scanner.Scan() {
		go func() {
			//time.Sleep(1 * time.Second)
			c <- scanner.Text()
			index++
		}()
		fmt.Printf("So dong hien tai: %d co gia tri la: %s \n", index, <-c)
	}
	if err := scanner.Err(); err != nil {
		fmt.Print("Error reading file.txt: ", err)
	}
	mu.Unlock()
}

func main() {
	Input(666, 1000)
	go Input(0, 333)
	go Input(333, 666)
	Output()
}
