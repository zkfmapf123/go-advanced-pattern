package main

import (
	"fmt"
	"time"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// TimeSleep 을 사용해야 boring이 호출됨
	go boring("hello world")

	fmt.Println("start")
	time.Sleep(10 * time.Second)
	fmt.Println("end")

}
