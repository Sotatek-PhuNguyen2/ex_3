package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

var mu sync.Mutex

func Input() {
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
	Input()
}
