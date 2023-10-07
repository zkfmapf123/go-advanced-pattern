package main

import "fmt"

func boring(msg string, quit chan string) chan string {

	c := make(chan string)

	go func() {

		for {
			select {
			case c <- msg:
				//  do nothing

			case <-quit:
				return
			}
		}
	}()

	return c
}

func main() {
	quit := make(chan string)
	b := boring("Hello world", quit)

	for i := 0; i < 10; i++ {
		fmt.Println(<-b)
	}

	// 종료조건
	quit <- "End"
}
