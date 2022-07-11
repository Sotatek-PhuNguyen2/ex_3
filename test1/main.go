package main

import (
	"log"
	"time"
)

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

func main() {
	chanRoutine()
}
