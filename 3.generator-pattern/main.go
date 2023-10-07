package main

import (
	"fmt"
	"time"
)

const NUM = 10

// channel 자체를 만들어서 그냥 줘버림
func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < NUM; i++ {
			c <- fmt.Sprintln(msg, i)
			time.Sleep(1 * time.Second)
		}
	}()

	return c
}

func main() {
	lee, ahn := boring("lee"), boring("ahn")

	for i := 0; i < NUM; i++ {
		fmt.Println(<-lee)
		fmt.Println(<-ahn)
	}

	fmt.Println("end")
}
