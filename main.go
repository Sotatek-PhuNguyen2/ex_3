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
var wg sync.WaitGroup

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

func Input(wg *sync.WaitGroup) {
	M = make(map[string]string)
	for i := 0; i < 1000; i++ {
		time.Sleep(1 * time.Millisecond)
		mu.Lock()
		t := strconv.Itoa(i)
		M[t] = t
		mu.Unlock()
	}
	wg.Done()
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
	c := make(chan string, 1000)
	var file, err = os.Open("file.txt")
	if err != nil {
		fmt.Print("Error opening file.txt: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c <- scanner.Text()
	}

	wg.Add(3)
	index1 := 0
	for i := 0; i < 3; i++ {
		go func() {
			for range <-c {
				index1++
				fmt.Printf("Dong thu %d co gia tri la: %s \n", index1, <-c)
			}
			wg.Done()
		}()
	}
	close(c)

	wg.Wait()

	if err := scanner.Err(); err != nil {
		fmt.Print("Error reading file.txt: ", err)
	}
}
func main() {
	/**TEST1
	chanRoutine()
	*/

	/** TEST2
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go Input(&wg)
	}
	wg.Wait()
	Output()
	*/

	/**TEST3
	errFunc()
	*/

	Test4()
}
