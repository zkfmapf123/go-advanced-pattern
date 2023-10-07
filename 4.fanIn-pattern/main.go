package main

import (
	"fmt"
)

func boring(msg string, num int) chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < num; i++ {
			c <- fmt.Sprintf("%d >> %s ", i, msg)
		}
	}()

	return c
}

func fanin(bs ...chan string) chan string {
	c := make(chan string)
	for _, ch := range bs {

		go func(cv chan string) {
			// 각각의 channel에 값이 쌇일때까지 계속 Channeling
			for {
				c <- <-cv
			}
		}(ch)
	}

	return c
}

func main() {
	num := 5
	b1, b2 := boring("hello", num), boring("bye", num)
	fan := fanin(b1, b2)

	// channel의 값이 존재하기때문에 다 소진될때까지 출력
	for i := 0; i < num; i++ {
		fmt.Println(<-fan)
	}

}
