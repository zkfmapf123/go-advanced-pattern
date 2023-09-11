package main

import (
	"fmt"
)

func simpleReceiveChannelMsg(m string) {
	myChannel := make(chan string)
	go func() {
		myChannel <- m
	}()

	msg := <-myChannel
	fmt.Println(msg)
}

func simpleReceiveChannelUseSelect(m string) {
	myChannel, anotherChannel := make(chan string), make(chan string)

	go func() {
		myChannel <- "data"
	}()

	go func() {
		anotherChannel <- "cow"
	}()

	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)

	case anotherMsgFromChannel := <-anotherChannel:
		fmt.Println(anotherMsgFromChannel)
	}
}

func main() {
	simpleReceiveChannelMsg("Hello")
	simpleReceiveChannelUseSelect("hello")
	fmt.Println("HEllo world")
}
