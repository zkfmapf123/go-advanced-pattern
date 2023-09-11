package main

import (
	"fmt"
	"time"
)

func someFunc(v int) {
	fmt.Print(v, "\t")
}

func main() {

	// 비동기 형태로 들어갔는데, 프로그램이 일찍 끝남
	go someFunc(1)
	go someFunc(2)
	go someFunc(3)

	fmt.Println("hello world")

	// 시간을 좀더 기다려준다면, 비동기 작업이 끝나고 나머지가 실행됨
	time.Sleep(time.Second * 2)
}
