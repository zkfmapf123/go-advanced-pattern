package main

import (
	"fmt"
	"time"
)

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintln(msg, i)
		time.Sleep(1 * time.Second / 2)
	}
}

func boringDeadLock(msg string, c chan string) {
	for i := 0; i < 1; i++ {
		c <- fmt.Sprintln(msg, i)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go boring("hello world", c1)

	// channel을 사용하면 알아서 동기화됨
	for i := 0; i < 10; i++ {
		fmt.Println("c1", <-c1)
	}

	// channel 양이 모자르면 -> Deadlock
	for i := 0; i < 10; i++ {
		fmt.Println("c2", <-c2)
	}
}
